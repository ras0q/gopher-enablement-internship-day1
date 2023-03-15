package main

import (
	"github.com/ras0q/multierrfinder"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(multierrfinder.Analyzer) }
