package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "os"
    "os/exec"
    "regexp"
    "strconv"
    "strings"
    "time"
)

// Структура ответа
type Response struct {
    Host   string  `json:"host"`
    Type   string  `json:"type"`
    Result float64 `json:"result"` // Используем float для всего (http код тоже число)
    Error  string  `json:"error,omitempty"`
}

var apiKey string

func main() {
    // Получаем ключ при старте
    apiKey = os.Getenv("API_KEY")
    if apiKey == "" {
        fmt.Println("WARNING: API_KEY not set!")
    }

    http.HandleFunc("/", handleRequest)
    
    // Запускаем сервер
    fmt.Println("Server started on :80")
    if err := http.ListenAndServe(":80", nil); err != nil {
        panic(err)
    }
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    // 1. Проверка API-ключа
    query := r.URL.Query()
    userKey := query.Get("key")

    if apiKey != "" && userKey != apiKey {
        w.WriteHeader(http.StatusForbidden)
        json.NewEncoder(w).Encode(map[string]string{"error": "Auth failed"})
        return
    }

    // 2. Валидация параметров
    host := query.Get("host")
    if host == "" {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]string{"error": "host required"})
        return
    }

    // 3. Выбор метода
    method := query.Get("method")
    if method == "" {
        method = "ping"
    }

    var result float64

    // 4. Выполнение
    switch method {
    case "http":
        result = checkHTTP(host, "http")
    case "https":
        result = checkHTTP(host, "https")
    default: // ping
        method = "ping"
        result = checkPing(host)
    }

    // 5. Ответ
    resp := Response{
        Host:   host,
        Type:   method,
        Result: result,
    }
    json.NewEncoder(w).Encode(resp)
}

func checkPing(host string) float64 {
    // В Go exec.Command безопаснее shell_exec, так как аргументы экранируются
    // Но для надежности запускаем именно ping
    // -c 3 (пакета), -W 2 (таймаут сек), -q (тихий режим)
    cmd := exec.Command("ping", "-c", "3", "-W", "2", "-q", host)
    output, err := cmd.CombinedOutput()

    if err != nil {
        return 0
    }

    // Парсим вывод Linux ping (ищем rtt min/avg/max/mdev)
    // Пример: rtt min/avg/max/mdev = 10.100/15.200/20.300/1.200 ms
    re := regexp.MustCompile(`(?m)/(\d+\.\d+)/(\d+\.\d+)/`)
    matches := re.FindStringSubmatch(string(output))

    if len(matches) >= 3 {
        val, _ := strconv.ParseFloat(matches[2], 64) // Берем avg (2-я группа)
        return val
    }

    return 0
}

func checkHTTP(host, scheme string) float64 {
    // Убираем протокол если юзер его ввел сам
    host = strings.TrimPrefix(host, "http://")
    host = strings.TrimPrefix(host, "https://")
    
    url := fmt.Sprintf("%s://%s", scheme, host)

    client := &http.Client{
        Timeout: 5 * time.Second,
        // Для HTTPS не проверяем валидность сертификата (как в PHP примере)
        Transport: &http.Transport{
            TLSClientConfig: nil, // Можно добавить InsecureSkipVerify: true если надо
        },
    }

    // Делаем HEAD запрос (только заголовки)
    resp, err := client.Head(url)
    if err != nil {
        return 0
    }
    defer resp.Body.Close()

    return float64(resp.StatusCode)
}