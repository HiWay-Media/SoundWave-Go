package main

import (
	"flag"
	"fmt"
)

func main() {
	flags := types.Flags{}
	err := flags.Set()
	if err != nil {
		panic(err)
	}
	//
	if flags.Help {
		flag.Usage()
		return
	}
	//
	fmt.Println("reading flags")
}