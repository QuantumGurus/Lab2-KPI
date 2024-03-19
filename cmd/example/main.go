package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	lab2 "github.com/roman-mazur/architecture-lab-2"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	inputFile       = flag.String("f", "", "File containing the expression")
	outputFile      = flag.String("o", "", "File to write the output to")
)

func main() {
	flag.Parse()

	var input io.Reader
	var output io.Writer

	// Handle input source
	if *inputExpression != "" {
		input = strings.NewReader(*inputExpression)
	} else if *inputFile != "" {
		file, err := os.Open(*inputFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error opening input file:", err)
			return
		}
		defer file.Close()
		input = file
	} else {
		fmt.Fprintln(os.Stderr, "No input source specified")
		return
	}

	// Handle output destination
	if *outputFile != "" {
		file, err := os.Create(*outputFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error creating output file:", err)
			return
		}
		defer file.Close()
		output = file
	} else {
		output = os.Stdout
	}

	handler := &lab2.ComputeHandler{
		Input:  input,
		Output: output,
	}

	if err := handler.Compute(); err != nil {
		fmt.Fprintln(os.Stderr, "Error computing expression:", err)
	}
}
