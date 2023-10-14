package ipcfs

import (
	"fmt"
	"os"
	"syscall"
)

const srvFmt = "/srv/%s"

func RegisterServer(name string, fd *os.File) error {
	var err error
	name = fmt.Sprintf(srvFmt, name)
	if err = syscall.Mkfifo(name, 0666); err != nil {
		return err
	}
	if fd, err = os.Open(name); err != nil {
		return err
	}
	return nil
}

func DeregisterServer(name string, fd *os.File) error {
	var err error
	name = fmt.Sprintf(srvFmt, name)
	fd.Close()
	if err = os.Remove(name); err != nil {
		return err
	}
	return nil
}
