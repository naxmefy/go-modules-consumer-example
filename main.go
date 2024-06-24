package main

import (
	"fmt"
	"github.com/naxmefy/go-modules/lib/a/b/c"
	"github.com/naxmefy/go-modules/module1"
	"github.com/naxmefy/go-modules/module1and2"
	"github.com/naxmefy/go-modules/module2"
)

func main() {
	module1.Hello()
	module2.Hello()

	fmt.Println("And now.......")

	module1and2.Hello()

	c.Hello()
}
