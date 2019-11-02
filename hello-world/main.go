package main

import (
	"fmt"
)

func HelloWorld() string {
	return "Hello, World!"
}

func Hello(name string) string {
	return "Hello, " + name + "!"
}

func main() {
	fmt.Println(HelloWorld())
	fmt.Println(Hello("Burt.K"))
}
