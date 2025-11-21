# Go lang

- https://gophercises.com/

- Go was created by Google to solve specific problems
  - Huge systems
  - Too many dependencies
  - Slow compile times
  - Hard concurrency (threads, locks)
- **Go Features**
  - simple language
  - fast compilation
  - built in concurrency
  - small binaries
  - automatic memory management
  - Go runs without a runtime(JVM, node)
  - packages & modules system

---

- Create a new module

```go
go mod init <module_name>
```

### File Convetions

- `go.mod` similar to package.json | module name, dependency list etc.
- `go.sum` this is a checksum file or to verify the packages

### Commands

- run a single file
- only compiles a single file

```go
go run main.go
```

- run a entire module, compile entire module
- requires module initialization `go mod init <name>`

```go
go run .
```

- `go mod tidy` adds missing import to `go.sum` file
- Run in root folder
- cleans and fixes dependencies
- adds missing imports, removes unused imports

```go
go mod tidy
```
