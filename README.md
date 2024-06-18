# Metrics Extractor

## Overview
Metrics Extractor is a Go application designed to collect and broadcast system metrics in real-time. It utilizes websockets for live data streaming to connected clients.

## Features
- Real-time system metrics extraction.
- Websocket support for live data broadcasting.

## Getting Started

### Prerequisites
- Go (version 1.15 or newer)

### Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/jcescobarn/metricsExtractor.git
   ```
2. Navigate to the project directory:
   ```bash
   cd metricsExtractor
   ```
3. Build the application:
   ```bash
   go build
   ```

### Running the Application
To start the server, run:
```bash
./metricsExtractor
```
The server will start, and you can connect to the websocket endpoint at `ws://localhost:8080/ws` to receive live system metrics.
