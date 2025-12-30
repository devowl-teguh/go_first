---
description: How to add a new feature following the Clean Layered Architecture
---

# Add Feature Workflow

Follow these steps to add a new feature to the Go API project. This ensures consistency across the Handler, Service, and Repository layers.

## 1. Define the Data Model
Create a new file in `internal/models/` or add to an existing one. Use GORM tags for database mapping.

```go
type Feature struct {
    ID        int       `json:"id" gorm:"primaryKey"`
    // ... other fields
    CreatedAt time.Time `json:"created_at"`
}
```

## 2. Create the Repository Layer
- Define an interface in `internal/repository/feature_repo.go`.
- Implement the interface using GORM.

## 3. Create the Service Layer
- Define an interface in `internal/service/feature_service.go`.
- Implement the business logic.

## 4. Create the Handler Layer
- Implement HTTP handlers in `internal/handler/feature_handler.go`.
- Handle request/response JSON and HTTP status codes.

## 5. Register Routes
Update [routes.go](file:///Users/asdf/Documents/demo/go_first/internal/api/routes.go) to register the new handlers.

## 6. Dependency Injection
Update [main.go](file:///Users/asdf/Documents/demo/go_first/cmd/api/main.go) to:
- Initialize the new Repository.
- Initialize the new Service.
- Initialize the new Handler and pass it to `RegisterRoutes`.

## 7. Add Tests
Create a test file for the service layer (e.g., `internal/service/feature_service_test.go`) and use a mock repository.
