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

	p("o: %s", o)
	p("")
	p("  // Obj() gets the actual object")
	p("  o.Obj() = %s", o.Obj())
	p("")
	p(`  o.IsMap() = %v`, o.IsMap())
	p(`  o.IsList() = %v`, o.IsList())
	p(`  o.Get("name") = %s`, o.Get("name"))
	p("")
	p(`  o.Set("name", "Tyler")`)
	o.Set("name", "Tyler")
	p("")
	p(`  o.Get("name") = %s`, o.Get("name"))
	p("")
	p(`  o.Set("age", 29)`)
	o.Set("age", 29)
	p("")
	p("  o: %s", o)
	p(`  o.Get("name").IsString() = %v`, o.Get("name").IsString())
	p(`  o.Get("name").IsInt() = %v`, o.Get("name").IsInt())

	fmt.Println()
	fmt.Println()
	fmt.Println()

}
