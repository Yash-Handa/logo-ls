// this file contain dir type definition
package main

import "time"

type file struct {
	name, ext, indicator string
	modTime              time.Time
	size                 int64 // in bytes
	mode                 string
	modeBits             uint32
	owner, group         string // use syscall package
}

type dir struct {
	info  *file
	files *[]file // all child fils and dirs
	dirs  *[]file // for recursion contain only child dirs
}

// define methods on *dir type only not on file type

func newDir(name string) *dir {
	return new(dir)
}

func (d *dir) print() []byte {
	return []byte("Yo")
}
