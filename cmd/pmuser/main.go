package main

import (
	"fmt"
	"github.com/lamg/pmproxy"
	"os"
)

func main() {
	u, e := pmproxy.InitUI(false)
	if e == nil {
		e = u.Run()
	}
	if e != nil {
		fmt.Fprintf(os.Stderr, e.Error())
	}
}
