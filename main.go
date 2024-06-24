package main

import (
	"fmt"
	"github.com/naxmefy/go-modules/lib/a/b"
	"github.com/naxmefy/go-modules/lib/a/b/c"
	"github.com/naxmefy/go-modules/module1"
	"github.com/naxmefy/go-modules/module1and2"
	"github.com/naxmefy/go-modules/module2"
)

func main() {
	module1.Hello()
	module2.Hello()
	fmt.Println("")
	fmt.Println("And now.......")
	fmt.Println("")
	module1and2.Hello()
	fmt.Println("")
	b.Hello()
	fmt.Println("")
	c.Hello()
}
