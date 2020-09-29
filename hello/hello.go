package main

import "fmt"

const englishGreetingPrefix = "Hello, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	if language == "Spanish"{
		return "Hola, " + name
	}
	return englishGreetingPrefix + name
}

func main(){
	fmt.Println(Hello("Chris", ""))
}