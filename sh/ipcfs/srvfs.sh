#!/bin/sh

# Library for use with srvfs
# Tested with ksh on NetBSD

srvfs.register() {
        mkfifo -m 666 /srv/$1
        return
}

srvfs.deregister() {
        rm /srv/$1
        return
}

srvfs.read() {
        cat /srv/$1
        return
}
