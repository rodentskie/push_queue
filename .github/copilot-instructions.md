# Push Queue Copilot Instructions

## Domain & Purpose

**Queue Management System** for educational institutions - think digital queue management like banks, but specialized for schools:
- **Finance Department**: Managing queues for tuition payments, fee transactions
- **Registrar**: Document requests, transcript processing
- **Dynamic & Configurable**: Must support various department workflows and queue types

## Architecture Overview

**Microservices** setup using **Go** backends, **Svelte** frontends, orchestrated in an **Nx monorepo**:

### Service Communication Patterns
- **GraphQL API Gateway** (`api` service) - Single entry point for clients
- **gRPC** - Service-to-service communication for data fetching (queries)
- **RabbitMQ** - Event-driven communication for mutations/state changes
- **MongoDB** - Primary data store (schema design should support dynamic queue configurations)

### Repository Structure
- **Go Workspace**: Uses `go.work` to manage multiple Go modules (apps and libraries) in a single workspace
- **Nx Build System**: Nx orchestrates all build, test, lint, and serve tasks via `project.json` files
- **Module Structure**: Each Go project has its own `go.mod` with module paths like `apps/api` or `library/go/env`
- **Directory Organization**: 
  - `apps/` - Go microservices (entry points with `main.go`) and Svelte frontends
  - `library/go/` - Shared Go libraries (reusable packages across services)

## Critical Workflows

### Creating New Projects

**New Go Library:**
```bash
NAME=mylib && nx g @obiente-lab/nx-go:library $NAME --directory library/go/$NAME
```

**New Go Application:**
```bash
NAME=myapp && nx g @obiente-lab/nx-go:application $NAME --directory apps/$NAME
```

**Remove Project:**
```bash
nx g rm <name>
```

### Development Commands

All commands use Nx executors (never use `go` commands directly):

**For Applications:**
- `nx serve <name>` - Hot reload with `gow` (requires `gow` installed)
- `nx build <name>` - Build binary
- `nx test <name>` - Run tests
- `nx lint <name>` - Lint code
- `nx tidy <name>` - Run `go mod tidy`

**For Libraries:**
- `nx test <name>` - Run tests
- `nx lint <name>` - Lint code
- `nx tidy <name>` - Run `go mod tidy`

## Service Communication Patterns

### gRPC for Queries
- Use gRPC for **synchronous data fetching** between services
- Generate `.proto` files in `library/proto/` for shared contracts
- Each microservice exposes gRPC endpoints for data it owns
- The `api` (GraphQL gateway) aggregates data via gRPC calls

### RabbitMQ for Mutations
- Use RabbitMQ for **asynchronous mutations and events**
- Commands that change state publish to RabbitMQ
- Services subscribe to relevant queues/exchanges
- Enables eventual consistency and decoupled writes

### GraphQL API Gateway
- The `api` service is the **sole client-facing endpoint**
- Resolves GraphQL queries by calling backend services via gRPC
- Handles mutations by publishing to RabbitMQ
- Aggregates responses from multiple microservices

## Project-Specific Conventions

### Module Path Convention
Go modules MUST use workspace-relative paths:
- Applications: `module apps/<app-name>`
- Libraries: `module library/go/<lib-name>`

Example from [apps/api/go.mod](apps/api/go.mod):
```go
module apps/api
go 1.25
```

### MongoDB Schema Design
- Support **dynamic queue configurations** (different departments, workflows)
- Consider embedding vs referencing for queue-ticket relationships
- Index on active queue status and timestamp fields for performance
- Schema should accommodate custom fields per department/queue type

### Project Configuration
Every Go project requires a `project.json` with Nx executors. See [apps/api/project.json](apps/api/project.json) for application pattern and [library/go/env/project.json](library/go/env/project.json) for library pattern.

### Hot Reload Configuration
The `serve` target uses `gow` (not standard Go tooling). Ensure `gow` is installed:
```bash
go install github.com/mitranim/gow@latest
```

## Prerequisites & Dependencies

**Required CLI Tools:**
- `protoc` (v26.1+) - Protocol buffer compiler
- `protoc-gen-go` (v1.31.0) - Go protobuf plugin
- `protoc-gen-go-grpc` (v1.3.0) - Go gRPC plugin
- `gow` - File watcher for hot reload

**Install protobuf tools:**
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.31.0
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0
```

## Key Files

- `go.work` - Go workspace configuration (lists all Go modules)
- `nx.json` - Nx configuration with `@obiente-lab/nx-go` plugin
- `project.json` - Per-project Nx target definitions
- `pnpm-workspace.yaml` - pnpm workspace config for Node tooling
