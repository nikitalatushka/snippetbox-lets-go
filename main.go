package main // tells Go compiler that this package (directory) is executable

import "fmt" // import "format" package; docs @ https://pkg.go.dev/fmt 

// entry point into the executable package
// shared libraries don't have package main or function main
func main() {
    fmt.Println("Hello world!")
}

// run with `go run <command>`
// specify the local path, the file name, or the full module path
// `$ go run snippetbox.latushka.dev`
// `$ go run main.go`
// `$ go run .`
