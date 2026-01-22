# Go Exercise Repository Instructions

## Project Overview
This is a **Go learning repository** containing weekly programming exercises, organized as standalone executable programs. Each exercise demonstrates progressively more complex Go concepts, from basic I/O to HTTP servers with form handling.

## Project Structure
- **`week_1/`**: Four exercises covering:
  - `exercise_1.go`: Basic console output
  - `exercise_2.go`: HTTP server basics
  - `exercise_3.go`: HTTP with query parameters and form handling
  - `exercise_4.go`: Stateful HTTP with POST request parsing and value persistence

Each file is a complete, standalone program that can be executed independently.

## Key Patterns & Conventions

### Package Structure
- All exercises use `package main` with a `main()` entry point
- Single-file programs (no imports of code from other exercise files)
- No external dependencies beyond Go standard library

### HTTP Exercise Pattern (exercises 2-4)
All HTTP exercises follow this structure:
```go
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    // Handle request
})
http.ListenAndServe(":8088", nil)
```
- **Fixed port**: All HTTP exercises use `:8088`
- **Single route**: All handle only the root path `/`
- **Inline handlers**: Request handling via anonymous functions
- **HTML strings**: Responses embed HTML directly using backtick string literals

### Progressive Complexity
- **ex 1**: Simple `fmt.Println()`
- **ex 2**: Basic HTTP with hardcoded response
- **ex 3**: Query parameter reading (`r.FormValue("q")`)
- **ex 4**: POST method handling and stateful value passing via hidden form inputs

## Developer Workflow

### Running Exercises
```bash
cd week_1
go run exercise_1.go
```

For HTTP exercises (2-4), the server starts on `http://localhost:8088`:
```bash
go run exercise_3.go  # Then visit http://localhost:8088/?q=YourName
go run exercise_4.go  # Then visit http://localhost:8088
```

### Code Quality
- No external linters or build scripts configured
- Exercises follow basic Go conventions (no gofmt enforced)
- Error handling is minimal (see `exercise_4.go` line 9: `count, _ := strconv.Atoi(...)`)

## AI Agent Guidance
When adding new exercises:
- Maintain single-file structure per exercise
- Use `:8088` port for all HTTP services
- Embed HTML directly in string literals (no template files)
- Follow the progressive complexity pattern
- Include `package main` and a `main()` function
- Exercises are self-contained and independently executable
