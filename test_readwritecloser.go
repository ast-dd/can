package can

import (
	"syscall"
)

type echoReadWriteCloser struct {
	closed  bool
	writeFd int
}

// NewEchoReadWriteCloser returns a ReadWriteCloser which echoes received bytes
// via an Unix domain socket pair.
func NewEchoReadWriteCloser() ReadWriteCloser {
	pair, err := syscall.Socketpair(syscall.AF_UNIX, syscall.SOCK_DGRAM, 0)
	if err != nil {
		panic(err)
	}
	return NewReadWriteCloser(&echoReadWriteCloser{writeFd: pair[0]}, pair[1])
}

func (rw *echoReadWriteCloser) Read(b []byte) (n int, err error) {
	panic("Read() shouldn't be called anymore")
}

func (rw *echoReadWriteCloser) Write(b []byte) (n int, err error) {
	err = syscall.Sendmsg(rw.writeFd, b, nil, nil, 0)
	n = len(b)

	return
}

func (rw *echoReadWriteCloser) Close() error {
	rw.closed = true
	return nil
}
