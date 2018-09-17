package main

import "context"

//here is a line coment

// FileGetter is usd to retriv de code.
type FileGetter interface {
	// GetFiles returs a FilesScanner
	GetFiles(context.Context, *FilesRequest) (FileScanner, error)
}

func main() {
	/*
	   Multi lyne
	   coment here
	*/
}
