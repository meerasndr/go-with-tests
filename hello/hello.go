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


	return greetingPrefix(language) + name	
}

func greetingPrefix(language string) (string) {
	prefix := englishGreetingPrefix
	switch language {
	case spanish:
		prefix = spanishGreetingPrefix
	case french:
		prefix = frenchGreetingPrefix
	}

	return prefix
}


func main(){
	fmt.Println(Hello("Chris", ""))
}