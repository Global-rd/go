package writer

type Writer interface {
	Write(string) error
	Flush() error
	Close() error
}
