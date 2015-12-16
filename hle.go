package hle

/*
#cgo CFLAGS: -I ${SRCDIR}/include
#cgo LDFLAGS: -L${SRCDIR}/build -lgohle
#include "hle.h"
*/
import "C"

type Locker struct {
	l C.uint
}

func New() *Locker {
	return &Locker{
		l: C.uint(0),
	}
}

func (l *Locker) Lock() {
	C.acquire_hle(&l.l)
}

func (l *Locker) Unlock() {
	C.release_hle(&l.l)
}
