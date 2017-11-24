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

## Guidelines

- The `cmd` module is where you will create your commands. Create a file for each command following the example of `cmd/hello.go`

- The `flag_variables.go` is used to declare variables that will receive the flags values, I did this to reuse the variables instead of creating new ones for each command.
  - Example: If have two commands with the variable `var source string`. You can declare the `source` in `command A` file and reuse it in `command B` file but it will get messy to know where the variables are declared. You can create a new variable but probably it will have a name almost equal to `source` and it's redundant because they're for the same thing.

- The `core` module is where you put things that's common and/or can be used to all your application like the application configuration
  - The `config.go` offers basic methods to work with the application configuration, you can create your own methods here
  - The `init.go` is where you put things that you will expose to other packages like the `AppConfiguration`

## Contributing

Bug reports and pull requests are welcome on GitHub at https://github.com/WendellAdriel/go-simple-cli-boilerplate/.
This project is intended to be a safe, welcoming space for collaboration.

## License

This project is available as open source under the terms of the [MIT License](http://opensource.org/licenses/MIT).
