# Tiny URL

Tiny URL is a URL shortener service built with Go and Redis. It allows you to generate short, easy-to-share URLs from long URLs.

## Getting Started

To get started with Tiny URL, you need to have Docker and Docker Compose installed on your machine.

## Installation
1. Clone this repository to your local machine:

```bash
git clone https://github.com/mikekbnv/tiny_url.git
```
 2. Navigate to the project directory:
```bash
cd tiny_url
```
 3. Start the application using Docker Compose:
```bash
docker-compose up
```
 4. Start the application without Docker Compose:

    Change ```Addr: "redis:6379"``` to ```Addr: ":6379"``` in main.go and run
    ```bash 
    go mod download
    ``` 
    then 
    ```bash 
    go run main.go
    ```
    Also, make sure that your redis-serve is up.
 
The Tiny URL application will be at http://localhost:2831.

## Usage
### Shortening a URL
To shorten a URL, send a POST request to the application's root endpoint with the "url" parameter set to the long URL you want to shorten. You can use tool like curl.

#### Option 1 using curl:
```bash
curl -X POST -d "url=https://example.com" http://localhost:2831
```
The response will contain the shortened URL.

#### Option 2 using the web:
 1. Open your web browser and visit http://localhost:2831.
 2. On the web interface, you will find a form where you can enter a long URL.
 3. Submit the form by clicking the "Shorten" button.
 4. Then will receive a shortened URL.

Both options will provide you with a shortened URL that you can use to access the original long URL.