package greet

import "fmt"

func SayHello() {
	fmt.Println("Hello !")
}

func Greet(name string) string {
	return fmt.Sprintf("Hello %s", name)
}
