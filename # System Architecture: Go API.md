# System Architecture: Go API

## Design Pattern: Clean Layered Architecture
We follow a strict dependency flow: **Handler -> Service -> Repository**.

### The Layers
1. **Cmd (`/cmd`)**: The application entry point. Responsible for "Wiring" (Dependency Injection) and starting the server.
2. **Handlers (`/internal/handler`)**: 
   - Responsible for HTTP protocol logic.
   - Parses JSON/Queries.
   - Calls the Service layer.
   - Returns HTTP Status Codes.
3. **Services (`/internal/service`)**: 
   - The **Business Logic**. 
   - Framework-agnostic. 
   - Coordinates data flow and applies domain rules.
4. **Repositories (`/internal/repository`)**: 
   - Data Access Object (DAO). 
   - Handles SQL queries or external API calls.
5. **Models (`/internal/models`)**: 
   - Plain Old Go Objects (POGO). 
   - Shared across all layers.

## Dependency Rule
- Inner layers (Service/Repository) **must not** depend on outer layers (Handler/HTTP).
- Communication between layers should happen via **Interfaces** to allow for easy mocking/testing.

## Implementation Standards

### 1. Error Handling
- Use `fmt.Errorf("...: %w", err)` to wrap errors from lower layers (Repository/Database).
- Handlers should map domain errors to appropriate HTTP status codes.

### 2. Context Management
- Pass `context.Context` through all layers for cancellation and timeouts.

### 3. Request Validation
- Validate inputs in the **Handler** layer.
- Business rule validation happens in the **Service** layer.

## Testing Strategy
- Test business logic in the `Service` layer using **Mock Repositories**.