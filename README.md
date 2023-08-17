# Golang: Test-Driven Development(TDD) with Gin and MySQL

### Get wet your hands with the TDD approach

<img src="https://github.com/cp-dharti-r/tdd-in-golang-with-gin-and-mysql/blob/main/img/cover-img.png">
<br />

In this repository, we've covered the implementation of crud operation with Test-Driven Development (TDD).

TDD is a software development approach in which tests are written before the implementation of the code. The goal of TDD is to ensure that the code meets the requirements and behaves as expected by constantly testing the code during the development process.

This can help to catch bugs early on, improve code quality, and increase confidence in the software.

TDD is an iterative process that begins with writing a test, running it to confirm it fails, then writing the minimum amount of code to make the test pass and repeating this process until the feature is complete.

## Getting Started

Following instructions will get you api platform running on your local machine for development and testing.

## Prerequisites

You need following prerequisites before setting up this project on your machine.

- Go with version go1.20 linux/amd64
- Mysql

If not install already, follow [how to install Golang](https://go.dev/doc/install)

## Setup

Following are steps to setup project on local.

- Clone this project.

- Modify database configuration at `main.go` -> `init()` and at `db.go` for testing with your local database credentials.

  ```
  root:[your-database-password]@/[your-database-name]
  ```

- Start Go Server

  ```
  go run main.go
  ```

  Install package by, run:

  ```
  go mod tidy
  ```

- Server will start on `http://localhost:8000`

## Run Tests

- To run the tests of user CRUD operation you need to run the below command on your root directory:

  ```
  go test -v user_test.go
  ```
