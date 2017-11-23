# Go Simple CLI Boilerplate

> Simple boilerplate to create CLI applications with Go

## How to use

### With Docker

- Clone this repo and enter it

- Run `docker build -t "mytag" .`

- Run `docker run -it --rm mytag hello`

### Without Docker

- Clone this repo and enter it

- Run `go get ./...`

- Run `go run app.go hello` or build the application with `go build` and execute the binary like: `./app hello`
