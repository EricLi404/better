package private

import _ "unsafe"

//go:linkname hi linkname/out.Hello
func hi() string {
	return "high"
}

func yes() bool {
	return true
}
