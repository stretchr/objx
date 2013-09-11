package main

import (
	"fmt"
	"github.com/stretchr/objx"
)

func p(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}

func main() {

	fmt.Println()

	o := objx.MustJson(`{"name":"Mat"}`)

	p("Obj: %s", o)
	p(`  o.Get("name") = %s`, o.Get("name"))
	p(`  o.Set("name", "Tyler")`)
	o.Set("name", "Tyler")
	p(`  o.Get("name") = %s`, o.Get("name"))
	p(`  o.Set("age", 29)`)
	o.Set("age", 29)
	p("  Obj: %s", o)

	fmt.Println()
	fmt.Println()
	fmt.Println()

}
