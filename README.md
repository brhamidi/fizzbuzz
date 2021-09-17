
Coverage Testing: 87.5% :innocent:


# Fizzbuzz

This is a server that computes the result of a fizzbuzz .

Read the file subject.txt for more reliable information about how the server works.

The server can be run in persistent mode or inmemory mode for storing statistics.

Just change the **STORE_MODE** environment variable to choose mode: set value (inmemory|persistent) in docker-compose file.

Simply run `make up` and go to http://localhost:3000/swagger/index.html#/

## Requirements

- GNU Make
- Docker-Compose
- Docker
- Golang >= 1.17 (required only for development purpose)

### Environment variables

To run the project, you need to set environment variables. The default variables are set in the `docker-compose.yml` file.

### Run

They are two mode to run the service:

- default mode
- development mode

`make up` to start the stack with all necessary services to run service.

`make dev` should be equivalent to the default mode with a hot reload system in addition, useful for development purposes.

### Local services

- **Swagger API Definition:**  http://localhost:3000/swagger/index.html

## Quality code

You need to run `make tools` to install the different tools needed for testing, linting ...

### Testing

`make test` to execute unit tests.

You can check the code coverage of the project by running this commands:

- `make cover`
- `make cover-html`

### Linting

We use [staticcheck](https://staticcheck.io/) as linter.

`make lint` to execute linting.
