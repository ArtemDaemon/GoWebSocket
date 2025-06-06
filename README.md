# Go WebSocket Application

This project is a simple Go application that demonstrates how to use WebSockets to update content in real-time. It includes both synchronous and asynchronous data loading methods to showcase performance differences.

Этот проект — простое приложение Go, демонстрирующее, как использовать WebSockets для обновления контента в реальном времени. Он включает в себя как синхронные, так и асинхронные методы загрузки данных для демонстрации различий в производительности.

## Project Structure

```
go-web-socket
├── main.go
├── handlers
│   └── websocket.go
├── utils
│   └── data_loader.go
├── go.mod
└── README.md
```

## Getting Started

### Prerequisites

- Go >= 1.16
- A web browser that supports WebSockets

### Installation

1. Clone the repository:

   ```
   git clone https://github.com/yourusername/go-websocket-app.git
   cd go-websocket-app
   ```

2. Install the necessary dependencies:

   ```
   go mod tidy
   ```

### Running the Application

To run the application, execute the following command:

```
go run main.go
```

The server will start on `localhost:8080`. You can access it via your web browser.

### Usage

- Open your web browser and navigate to `http://localhost:8080`.
- The application will establish a WebSocket connection and allow you to send and receive messages in real-time.

### Data Loading

The application demonstrates two methods of loading heavy data:

1. **Synchronous Loading**: Use the `LoadDataSync` function from `utils/data_loader.go` to load data in a blocking manner.
2. **Asynchronous Loading**: Use the `LoadDataAsync` function to load data without blocking the main thread, allowing for better performance and responsiveness.

### Contributing

Feel free to submit issues or pull requests for improvements or bug fixes.

### License

This project is licensed under the MIT License. See the LICENSE file for details.