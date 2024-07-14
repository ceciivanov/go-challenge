# GlobalWebIndex Engineering Challenge

The GlobalWebIndex Engineering Challenge is an application designed to manage user data and their favorite assets, including charts, insights, and audience profiles. It showcases practical backend development principles and aims to demonstrate a modular approach to application architecture.

## Table of Contents

- [Overview](#overview)
- [Project Structure](#project-structure)
- [Dependencies](#dependencies)
- [Setup and Installation](#setup-and-installation)
  - [Running Locally](#running-locally)
  - [Using Docker](#using-docker)
- [Usage](#usage)
  - [Endpoints](#endpoints)
  - [Examples](#examples)
- [Testing](#testing)
- [Contributing](#contributing)
- [License](#license)

---

## Overview

The project implements a backend service (API) that allows users to:

- Retrieve a list of their favorite assets categorized into charts, insights, and audience profiles.
- Add new assets to their favorites.
- Remove existing assets from their favorites.
- Update details of existing favorite assets.

## Project Structure

The project is organized into the following directories:

- `cmd`: Contains the main application entry point.
- `internal/models`: Contains the data models used in the application. 
- `internal/repository`: Handles in-memory data storage and retrieval.
- `internal/handlers`: Implements HTTP request handlers for the API endpoints.
- `internal/service`: Implements business logic and interacts with repositories.
- `internal/utils`: Contains utility functions, like decoding JSON data.


## Dependencies

The project leverages Go language features and standard libraries. Additionally, it uses some third-party libraries for routing and testing.

- Go version - go1.22.4
- Third-party libraries
    - github.com/gorilla/mux
    - github.com/stretchr/testify


## Setup and Installation

Explain how to set up and install your project.

### Running Locally

To run the application locally, follow these steps:

1. Clone the repository:

```bash
git clone 
``` 

2. Change into the project directory:

```bash
cd globalwebindex-engineering-challenge
```

3. Build the application:

```bash
go mod tidy
go build -o app cmd/app/main.go
```

4. Run the application:

```bash
./app
```

The application will start on port `8080` by default. You can access the API at `http://localhost:8080`.


## Usage

The application exposes RESTful endpoints for interacting with user data and favorite assets. Examples of API requests are provided to demonstrate functionality and usage scenarios.

### Endpoints

The API provides the following endpoints:

- `GET /users/{userID}/favorites`: Retrieve a list of favorite assets for a user.
- `POST /users/{userID}/favorites`: Add a new asset to a user's favorites.
- `PUT /users/{userID}/favorites/{assetID}`: Update details of an existing favorite asset.
- `DELETE /users/{userID}/favorites/{assetID}`: Remove an existing asset from a user's favorites.

### Examples

The following examples demonstrate how to interact with the API using `curl` commands.


## Testing

The project includes unit tests for the service and repository layers. The tests are implemented using the Go standard library and the `testify` library for assertions.

To run the tests, execute the following command:

```bash
go test ./... -cover
```


