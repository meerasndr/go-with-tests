package main

import "fmt"

const englishGreetingPrefix = "Hello, "
const spanishGreetingPrefix = "Hola, "
const spanish = "Spanish"

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	if language == spanish{
		return spanishGreetingPrefix + name
	}
	return englishGreetingPrefix + name
}

func main(){
	fmt.Println(Hello("Chris", ""))
}