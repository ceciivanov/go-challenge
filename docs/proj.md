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
|   |__ data/
|       |__ users_data.go
│
└── go.mod


- main.go: Main entry point of application.

- pkg/: Contains all your application packages.
    - handlers/: HTTP request handlers. 
    - models/: Structs and interfaces representing models.
    - utils/: Utility functions.
    - data/: Application in-memory data.
    - tests/: Test files for the application.
