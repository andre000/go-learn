package main

import "fmt"

var prefix = map[string]string{
	"EN": "Hello, ",
	"PT": "Olá, ",
	"FR": "Bonjour, ",
}

var defaultName = map[string]string{
	"EN": "Stranger",
	"PT": "Estranho",
	"FR": "Étranger",
}

// Hello : returns the phrase "Hello, {name}" in the defined language
func Hello(name string, lang string) string {
	if lang == "" {
		lang = "EN"
	}

	if name == "" {
		name = defaultName[lang]
	}

	return prefix[lang] + name + "!"
}

func main() {
	fmt.Println(Hello("Andre", ""))
}
