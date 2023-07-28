package out

import _ "linkname/private"
import _ "unsafe"

func Hello() string

//go:linkname  Yes linkname/private.yes
func Yes() bool
