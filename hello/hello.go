package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const chinese = "Chinese"
const helloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const chineseHelloPrefix = "你好, "

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case chinese:
		prefix = chineseHelloPrefix
	default:
		prefix = helloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Chris", ""))
}
