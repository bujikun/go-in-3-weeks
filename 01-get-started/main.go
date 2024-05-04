package main

import (
	"fmt"
	"github.com/jboursiquot/go-proverbs"
	"strconv"
)

const location = "Remote"

var name string

func main() {
	name = "Newton"
	from := "Tanzania"
	var n int
	//p := &name
	var proverb string
	if p, err := proverbs.Nth(4); err == nil {
		proverb = p.Saying
	}
	fmt.Printf("Hello, fellow %s\n", location)
	fmt.Printf("My name is %s and I'm from %s\n", name, from)
	fmt.Printf("By the time %d o'clock EST comes around, we'll know how to code in Go\n", n)
	fmt.Printf("Here's a Go proverb: %s\n", proverb)
	fmt.Println("Let's get started!")
	fmt.Println("" + strconv.Itoa(8))
	//fmt.Println()
	//fmt.Println(reflect.TypeOf(name).Size())
	//fmt.Println(reflect.TypeOf(p).Size())
	//fmt.Println(reflect.TypeOf(n).Size())

}
