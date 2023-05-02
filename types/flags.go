package types

import (
	"flag"
	"fmt"
)

type Flags struct {
	Input  string
	Output string

	Help bool
}

func (f *Flags) Set() error {
	//
	input := flag.String("i", "", "path of the input file")
	output := flag.String("o", "", "path of the output soundwave file")

	help := flag.Bool("help", false, "shows help")
	//
	flag.Parse()
	//
	f.Input = *input
	f.Output = *output
	f.Help = *help
	//
	if f.Input == "" {
		return fmt.Errorf("need to specify input")
	}
	if f.Output == "" {
		return fmt.Errorf("need to specify output")
	}

	return nil
}
