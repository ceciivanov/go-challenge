go-challenge/
│   └── main.go
│
├── pkg/
│   ├── handlers/
│   │   └── user_handlers.go
│   │
│   ├── models/
│   │   ├── asset.go
│   │   ├── chart.go
│   │   ├── insight.go
│   │   └── audience.go
│   |   └── user.go
│   │
│   └── utils/
│       └── decoder.go
│   |
|   |__ repository/
|       |__ mock_data/
|       |   |__ mock_data.go
|       |   |__ random.go
|       |
|       |__ users_repository.go
|   |__
├──     tests/
│       ├── user_handlers_test.go
│
└── go.mod


- main.go: Main entry point of application.

- pkg/: Contains all your application packages.
    - handlers/: HTTP request handlers. 
    - models/: Structs and interfaces representing models.
    - utils/: Utility functions.
    - repository/: Application in-memory data.
    - tests/: Test files for the application.
