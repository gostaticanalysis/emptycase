package main

import (
	"github.com/gostaticanalysis/emptycase"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(emptycase.Analyzer) }
