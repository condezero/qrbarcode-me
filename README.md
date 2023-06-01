# Barcode and QR Code Generation API in Golang

This is an API written in the Go programming language (Golang) that allows generating barcodes and QR codes based on provided content. It uses the `gozxing` and `gin-gonic` libraries to perform the code generation and provide an HTTP interface to access the functionality.

## Prerequisites

- Go (Golang) installed on the system.
- The following Go libraries installed:
  - `github.com/gin-gonic/gin`
  - `github.com/joho/godotenv`
  - `github.com/makiuchi-d/gozxing`
- Docker installed on the system (if you want to run the API using Docker).

## Configuration

The API uses an `.env` file for configuration. Make sure to create this file in the root directory of the project and set the following environment variables:

- `GIN_MODE`: The mode in which the Gin framework should run (e.g., `"debug"`, `"release"`).
- `PORT`: The port number on which the API server should listen for incoming requests.

If the `.env` file is not found, the API will attempt to load the configuration from the system environment variables.

## Usage (without Docker)

To use the API without Docker, follow these steps:

1. Install the required dependencies by running the following command:
   ```
   go get github.com/gin-gonic/gin github.com/joho/godotenv github.com/makiuchi-d/gozxing
   ```

2. Create an `.env` file in the root directory of the project and set the required environment variables as mentioned in the "Configuration" section.

3. Build and run the API server using the following command:
   ```
   go run main.go
   ```

4. The API server will start running on the specified port. You can send HTTP POST requests to the provided endpoints to generate barcodes and QR codes.

## Usage (with Docker)

To use the API with Docker, follow these steps:

1. Install Docker on your system by following the official Docker installation guide for your operating system.

2. Create an `.env` file in the root directory of the project and set the required environment variables as mentioned in the "Configuration" section.

3. Build the Docker image using the following command:
   ```
   docker build -t barcode-api .
   ```

4. Run the Docker container using the following command:
   ```
   docker run -p <host-port>:<container-port> -d barcode-api
   ```

   Replace `<host-port>` with the port number on your host machine where you want to access the API, and `<container-port>` with the port number specified in the `.env` file.

5. The API server will start running inside the Docker container. You can send HTTP POST requests to the provided endpoints to generate barcodes and QR codes.

## API Endpoints

- `POST /generate/barcode`: Generate a barcode. The request body should contain a JSON object with the following fields:
  - `content` (string): The content for the barcode.
  - `width` (integer): The width of the barcode image.
  - `height` (integer): The height of the barcode image.

- `POST /generate/qrcode`: Generate a QR code. The request body should contain a JSON object with the same fields as for barcode generation.

The API will respond with the generated code image as an octet-stream.

## Example

Here's an example of how to generate a barcode using the API:

Endpoint: `POST /generate/barcode`

Request body:
```json
{
  "content": "123456789",
  "width": 200,
  "height": 100
}
```

Response:
```
HTTP/1.1 200 OK
Content-Type: application/octet-stream

<barcode image data>
```

## Error Handling

- If the request payload is invalid or missing required fields, the API will respond with a status code of 400 (Bad Request).

- If an error occurs during code generation, the API will respond with a status code of 500 (Internal Server Error) and provide an error message.

## License

This API is released under the MIT License. See the [LICENSE](LICENSE) file for more details.