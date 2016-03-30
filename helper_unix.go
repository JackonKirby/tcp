// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build darwin dragonfly freebsd linux netbsd openbsd solaris

package tcp

import (
	"net"
	"reflect"
)

func (c *Conn) sysfd() (int, error) {
	switch c := c.Conn.(type) {
	case *net.TCPConn:
		cv := reflect.ValueOf(c)
		switch ce := cv.Elem(); ce.Kind() {
		case reflect.Struct:
			nfd := ce.FieldByName("conn").FieldByName("fd")
			switch fe := nfd.Elem(); fe.Kind() {
			case reflect.Struct:
				fd := fe.FieldByName("sysfd")
				return int(fd.Int()), nil
			}
		}
	}
	return 0, errInvalidConnType
}
