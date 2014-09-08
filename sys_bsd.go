// Copyright 2014 Mikio Hara. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build freebsd dragonfly

package tcp

const (
	sysGETSOCKOPT = 0x76

	sysTCP_KEEPIDLE  = 0x100
	sysTCP_KEEPCNT   = 0x400
	sysTCP_KEEPINTVL = 0x200
	sysTCP_NOPUSH    = 0x4
)
