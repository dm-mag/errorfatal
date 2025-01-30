# errorfatal
A utility package for Go that simplifies error handling

## Description

This package eliminates the repetitive error checking pattern:
```go
if err != nil {
    log.Fatal(err)
}
```

And replaces it with a more concise form:
```go
import e "github.com/dm-mag/errorfatal"

// Single error check
e.F(os.Chdir("123"))

// Chained operations
e.F(os.WriteFile("123.txt", e.F2(os.ReadFile("456.txt")), 0644))
```
## Installation

go get github.com/dm-mag/errorfatal

## Features

- Simplified error handling
- Reduces code verbosity
- Chain-friendly API
- Zero dependencies

### Examples

``` go
// Basic error check
e.F(os.MkdirAll("path/to/dir", 0755))

// Return value with error check
data := e.F2(os.ReadFile("config.json"))
```
