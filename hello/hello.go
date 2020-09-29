package main

import "fmt"

const englishGreetingPrefix = "Hello, "
const spanishGreetingPrefix = "Hola, "
const frenchGreetingPrefix = "Bonjour, "
const spanish = "Spanish"
const french = "French"

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	if language == spanish{
		return spanishGreetingPrefix + name
	}else if language == french{
		return frenchGreetingPrefix + name
	}
	return englishGreetingPrefix + name
}

func main(){
	fmt.Println(Hello("Chris", ""))
}