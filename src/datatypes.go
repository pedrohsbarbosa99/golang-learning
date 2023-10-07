package main

import "fmt"

func main() {
	//  basic types
	var a bool = true
	var b int = 5
	var c float32 = 3.14
	var d string = "Hi!"

	fmt.Println("Boolean: ", a)
	fmt.Println("Integer: ", b)
	fmt.Println("Float: ", c)
	fmt.Println("String: ", d)

	//  strings
	var txt1 string = "Hello!"
	var txt2 string
	txt3 := "World 1"
	fmt.Println(txt1, txt2, txt3)

	//  Array
	var text_array = [4]string{"Texto 1", "texto 2", "Texto 3", "Texto 4"} // here length is defined
	fmt.Println(text_array)

	var text_inferred_array = [...]string{"Texto 1", "texto 2", "Texto 3", "Texto 4"}
	fmt.Println(text_inferred_array)

	
}