package main

import (
	"flag"
	"fmt"

	"github.com/HiWay-Media/SoundWave-Go/types"
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
