# Go lang

- Go lang specification - https://go.dev/ref/spec
- Effective writing - https://go.dev/doc/effective_go

- https://gophercises.com/
- https://quii.gitbook.io/learn-go-with-tests
- https://100go.co/
- https://github.com/apache/answer

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
- Go is a multi paradigm language and has
  - functions
  - structs
  - methods
  - interfaces
  - closures (functional feature)
  - goroutines & channels (concurrency model)
- Go does NOT have:
  - classes
  - inheritance
  - generics in OOP style
  - method overloading
  - operator overloading

---

- [x] helloworld
- [x] external packages
- [x] variables
- [x] constants
- [x] types
- [x] functions
  - [x] multiple return functions
  - [x] variadic functions
- [x] loops & flow-statements
  - [x] for,range
  - [x] if/else
  - [x] switch
- [x] arrays
- [x] slices
- [x] maps
- [x] pointers
- [ ] clousers
- [ ] recursion
- [x] structs, struct methods
- [x] interfaces
- [ ] enums
- [x] error handling
- [ ] go modules

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
