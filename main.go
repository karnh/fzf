package main

import (
	"github.com/karnh/fzf/src"
	"github.com/karnh/fzf/src/protector"
)

var revision string

func main() {
	protector.Protect()
	fzf.Run(fzf.ParseOptions(), revision)
}
