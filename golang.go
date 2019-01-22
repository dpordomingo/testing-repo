// package docummentation
package main

import "context"
import "fmt"

// here is a line coment

// FileGetter is used to retriv the code.
type FileGetter interface {
	// GetFiles returs a FilesScanner and an error
	GetFiles(context.Context, *FilesRequest) (FileScanner, error)
}

func main() {
	fmt.Println("hello world")
	fmt.Println("hello world 2")
	
fmt.Println("hello world 3", "very very very very very very very very very very very very very very very long line, that should be reported by lookout")
	/*
	   multi line
	   coment here
	   new line
	   NEW
	   HERE
	*/
}
