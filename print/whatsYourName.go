package main

import "fmt"

func main() {
	var name string
	hello()
	get_name(name)
}

func hello() {
	fmt.Println("Hello!! I'm your computer..")
}

func get_name(name string) {
	fmt.Print("Let me know your name ? ")
	// fmt.Scanf("%s", &name)
	fmt.Scanln(&name)
	fmt.Printf("Oh! Your name is %s..! HI!!", name)
}