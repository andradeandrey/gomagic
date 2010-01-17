package main

import (
	"magic"
	"fmt"
)

func main() {
	cookie := magic.Open(magic.MAGIC_DEBUG|magic.MAGIC_CHECK|magic.MAGIC_ERROR)
	magic.Load(cookie, "/usr/share/file/magic")
	r := magic.File(cookie, "./sample.go")
	fmt.Printf("sample.go: %s\n", r)
	magic.Close(cookie)
}
