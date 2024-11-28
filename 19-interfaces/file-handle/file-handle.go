package filehandle

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type FileHandle interface {
	Reader
	Writer
	Close() error
}
