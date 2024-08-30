# JSON Marshalling Project
=====================================

This project demonstrates how to marshal and unmarshal JSON data in Go. It consists of three packages: `main`, `unmarshalling`, and `marshalling`.

## Overview

The `main` package serves as the entry point of the program, demonstrating both marshaling and unmarshaling capabilities.

*   The `unmarshalling` package, located under the `unmarshalling` directory, contains a function called `TryUnmarshal`. This function fetches JSON data from an API endpoint, unmarshals it into a Go struct (`UserResponse`), and then prints out each user's details.
*   The `marshalling` package, located under the `marshalling` directory, demonstrates how to create a Go struct (`Products`) representing a list of products with attributes and dimensions. It showcases how to marshal this struct into a JSON string.

## Installation

To run this project, you will need:

1.  A Go development environment set up on your machine.
2.  The `encoding/json` package is used for marshaling and unmarshaling JSON data. You can install it by running the command: `go get -u github.com/golang/protobuf/protoc-gen-go`
3.  An API endpoint to fetch user data from (as demonstrated in `TryUnmarshal` function).

## Usage

To run the program, navigate into the project directory and execute:

```go run main.go
```

The output will display both marshaled and unmarshaled JSON data for demonstration purposes.
