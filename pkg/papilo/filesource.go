package papilo

import (
	"bufio"
	"os"
)

// FileSource implements a default file data source
type FileSource struct {
	filepath string
	fdesc    *os.File
}

// NewFileSource returns a new file data source for streaming lines of a file.
// The path parameter is the path of the file to be read.
func NewFileSource(path string) FileSource {
	return FileSource{
		filepath: path,
	}
}

// NewFdSource returns a new file data source for streaming lines of a file.
// The fd parameter is an opened file to be read.
func NewFdSource(fd *os.File) FileSource {
	return FileSource{
		fdesc: fd,
	}
}

// Source is the implementation for the Sourcer interface.
// Defined output for this source is a slice of bytes.
func (f FileSource) Source(p *Pipe) {
	var fd *os.File = f.fdesc
	if fd == nil {
		var err error
		fd, err = os.Open(f.filepath)
		if err != nil {
			panic(err)
		}
	}

	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		p.Write(scanner.Bytes())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
