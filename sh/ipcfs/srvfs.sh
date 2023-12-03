#!/bin/sh

# Library for use with srvfs
# Tested with ksh on NetBSD

# srvfs.register(name)
srvfs.register() {
        mkfifo -m 666 /srv/$1
        return
}

# srvfs.deregister(name)
srvfs.deregister() {
        rm /srv/$1
        return
}

# srvfs.read(name) -> string
srvfs.read() {
        cat /srv/$1
        return
}
