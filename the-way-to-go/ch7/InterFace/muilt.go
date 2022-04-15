package InterFace

type Reader interface {
	Read(p []byte) (n int, err error)
}
type Closer interface {
	Close() error
}

type ReadClose interface {
	Reader
	Closer
}
