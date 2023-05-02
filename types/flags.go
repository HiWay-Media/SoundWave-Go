package types

import (
	"flag"
	"fmt"
)

type Flags struct {
	Input  string
	Output string
	Width     int
	Height    int
	Help bool
}

func (f *Flags) Set() error {
	//
	input := flag.String("i", "", "path of the input file")
	output := flag.String("o", "output.png", "path of the output soundwave file")
	width := flag.Int("w", 640, "frame width")
	height := flag.Int("h", 120, "frame height")
	//
	help := flag.Bool("help", false, "shows help")
	//
	flag.Parse()
	//
	f.Input = *input
	f.Output = *output
	f.Width = *width
	f.Height = *height
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
