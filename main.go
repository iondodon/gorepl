package main

import (
	"fmt"

	"github.com/c-bata/go-prompt"
	"github.com/traefik/yaegi/interp"
)

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "import", Description: "Imports a package"},
		{Text: "func", Description: "Define a new function"},
		{Text: "return", Description: "Return a value in a function"},
		// add more here
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func main() {
	i := interp.New(interp.Options{})

	for {
		in := prompt.Input("> ", completer, prompt.OptionTitle("Go REPL with auto-suggestion and yaegi interpreter"))
		_, err := i.Eval(in)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
}
