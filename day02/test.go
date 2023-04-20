package main

import "fmt"

func main() {
	var a *string
	*a = "军哥NB"
	fmt.Println(*a)

	var b map[string]string
	b["军哥"] = "YYDS"
	fmt.Println(b)
}
