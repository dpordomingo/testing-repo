package main

import "context"

//here is a line coment

// FileGetter is used to retriv the code.
type FileGetter interface {
	// GetFiles returs a FilesScanner
	GetFiles(context.Context, *FilesRequest) (FileScanner, error)
}

func main() {
	/*
	   Multi line
	   coment here
	*/
}
