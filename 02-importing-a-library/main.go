package main

import (
	"fmt"
	stringutil "github.com/bujikun/go-string-utils"
)

const location = "REMOTE"

var name string

func main() {
	name = "Newton"
	name = stringutil.Upper(name)
	from := "Tanzania"
	fmt.Printf("Hello, fellow %s Gophers\n", stringutil.Lower(location))
	fmt.Printf("My name is %s and I'm from %s\n", name, from)
}
