// Copyright 2010 GoNeedle Authors. All rights reserved.
// Use of this source code is governed by a 
// license that can be found in the LICENSE file.

package needle

import (
	"fmt"
	"os"
	"os/signal"
)

func InstallCtrlCPanic() {
	go func() {
		for s := range signal.Incoming {
			if s == signal.Signal(signal.SIGINT) {
				panic("Ctrl-C signalled")
			}
		}
	}()
}

func Logf(f string, rest ...interface{}) {
	fmt.Fprintf(os.Stderr, f, rest)
}
