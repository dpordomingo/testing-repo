// pacage docummentation
package main

import "context"
import "fmt"

// here is a line coment

// FileGetter is used to retriv the code.
type FileGetter interface {
	// GetFiles returs a FilesScanner
	GetFiles(context.Context, *FilesRequest) (FileScanner, error)
}

func main() {
	fmt.Println("hello world")
	/*
	   multi line
	   coment here
	*/
}
