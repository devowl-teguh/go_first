# Project Directory Layout

```text
.
├── cmd/
│   └── api/
│       └── main.go           # Dependency Injection & Server Start
├── internal/
│   ├── handler/              # HTTP Handlers / Controllers
│   │   └── user_handler.go
│   ├── service/              # Business Logic Interfaces & Impl
│   │   └── user_service.go
│   ├── repository/           # DB Queries & Interfaces
│   │   └── user_repo.go
│   ├── models/               # Structs for DB and JSON
│   │   └── user.go
│   ├── middleware/           # Auth, Logging, Recovery
│   └── config/               # Env vars and Config loading
├── pkg/
│   └── database/             # Shared DB connection helpers
├── .env                      # Local Environment Variables
├── go.mod                    # Go Modules
└── Makefile                  # Build and Dev scripts