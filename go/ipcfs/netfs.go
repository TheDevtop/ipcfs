package ipcfs

import (
	"fmt"
	"net"
	"os"
)

const (
	netFmt   = "/net/%s"
	unixType = "unix"
)

func RegisterNetwork(name string) (*net.UnixListener, error) {
	var (
		err error
		sd  *net.UnixListener
	)
	name = fmt.Sprintf(srvFmt, name)
	if sd, err = net.ListenUnix(unixType, &net.UnixAddr{
		Name: name,
		Net:  unixType,
	}); err != nil {
		return nil, err
	}
	return sd, nil
}

func DeregisterNetwork(name string, sd *net.UnixListener) error {
	var err error
	name = fmt.Sprintf(srvFmt, name)
	sd.Close()
	if err = os.Remove(name); err != nil {
		return err
	}
	return nil
}
