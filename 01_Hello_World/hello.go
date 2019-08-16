package main

import "fmt"

var PREFIX = map[string]string{
	"EN": "Hello, ",
	"PT": "Olá, ",
	"FR": "Bonjour, ",
}

var DEFAULT = map[string]string{
	"EN": "Stranger",
	"PT": "Estranho",
	"FR": "Étranger",
}

func Hello(name string, lang string) string {
	if lang == "" {
		lang = "EN"
	}

	if name == "" {
		name = DEFAULT[lang]
	}

	return PREFIX[lang] + name + "!"
}

func main() {
	fmt.Println(Hello("Andre", ""))
}
