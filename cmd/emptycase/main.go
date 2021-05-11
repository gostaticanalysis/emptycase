package main

import (
	"emptycase"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(emptycase.Analyzer) }
