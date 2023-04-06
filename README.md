# GoTastrophe

An ~~joke~~ experiment in golang picoservices.

## Getting Started

To get started with GoTastrophe

1. Ensure you have Go 1.19 installed

1. Clone the repository

1. Design the api of your dreams

1. Write your handler functions for this api in `./routes/routes.go` in the form:

    - `{method}{PascalCasedRouteName}Handler`
    - Eg. `GETHelloWorldHandler`

1. When you're ready to pico your services, run the build script:

    - `go run ./build/main.go`

This will generate a directory and `main.go` for each route you've created and build each main in that same directory.

6. You're ready to build your dreams!

## But... Why?

This repository started as a joke taking microservices to the logical extreme, and poking fun at a good friend and mentor of mine.

If you find any practical value of what is contained within these ~~hallowed halls~~ strange scripts, I wish you the best.:w
