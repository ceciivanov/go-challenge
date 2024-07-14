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
git clone https://github.com/ceciivanov/platform-go-challenge.git
``` 

2. Change into the project directory:

```bash
cd platform-go-challenge
git checkout implementation
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

5. To stop the application, press `Ctrl + C` in the terminal where the app is running.


### Using Docker

To run the application using Docker, follow these steps:

1. Build the Docker image:

```bash
docker build -t platform-go-challenge .
```

2. Run the Docker container:

```bash
docker run -p 8080:8080 platform-go-challenge
```

The application will start inside a Docker container and be accessible at `http://localhost:8080`.

3. To stop the container, use the following command:

```bash
docker stop <container_id>
```

Replace `<container_id>` with the ID of the running container. You can find the container ID by running `docker ps`.



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

```bash
# GET existing user favorites
curl -X GET http://localhost:8080/users/1/favorites

# GET user favorites which do not exist
curl -X GET http://localhost:8080/users/999999/favorites

# DELETE existing user favorite
curl -X DELETE http://localhost:8080/users/1/favorites/1

# DELETE user favorite which does not exist
curl -X DELETE http://localhost:8080/users/1/favorites/999999

# ADD valid user favorite
curl -X POST http://localhost:8080/users/1/favorites \
     -H "Content-Type: application/json" \
     -d '{
          "id": 100,
          "type": "Audience",
          "description": "This audience is a 40 year old",
          "age": 40,
          "ageGroup": "25-45",
          "gender": "Male",
          "birthCountry": "USA",
          "hoursSpentOnMedia": 4,
          "numberOfPurchases": 10
         }'

# ADD user favorite which already exists
curl -X POST http://localhost:8080/users/1/favorites \
     -H "Content-Type: application/json" \
     -d '{
          "id": 2,
          "type": "Insight",
          "description": "Sample Insight for testing",
          "text": "Testing Insight"
         }'

# ADD user favorite with invalid type
curl -X POST http://localhost:8080/users/1/favorites \
     -H "Content-Type: application/json" \
     -d '{
          "id": 200,
          "type": "INVALIDTYPE",
          "description": "Sample Insight for testing",
          "text": "Testing Insight"
         }'

# EDIT the previously added user favorite
curl -X PUT http://localhost:8080/users/1/favorites/100 \
     -H "Content-Type: application/json" \
     -d '{
          "id": 100,
          "type": "Audience",
          "description": "Updated Audience",
          "age": 18,
          "ageGroup": "18-25",
          "gender": "Female",
          "birthCountry": "Greece",
          "hoursSpentOnMedia": 15,
          "numberOfPurchases": 25
         }'

# EDIT user favorite with mismatched id (assetID in URL and id in body)
curl -X PUT http://localhost:8080/users/2/favorites/2 \
     -H "Content-Type: application/json" \
     -d '{
          "id": 1,
          "type": "Insight",
          "description": "Sample Insight for testing",
          "text": "Testing Insight"
         }'

# EDIT user favorite with mismatched type
curl -X PUT http://localhost:8080/users/1/favorites/100 \
     -H "Content-Type: application/json" \
     -d '{
          "id": 100,
          "type": "Insight",
          "description": "Sample Insight for testing",
          "text": "Testing Insight"
         }'
```


## Testing

The project includes unit tests for the service and repository layers. The tests are implemented using the Go standard library and the `testify` library for assertions.

To run the tests, execute the following command:

```bash
go test -v ./... -coverprofile=coverage.out
```

To view the test coverage report in html format on your browser, install the `go tool cover` package and run the following command:

```bash
go tool cover -html=coverage.out
```
