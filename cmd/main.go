package main

import (
	"flag"
	"fmt"
	"path"

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
	logfile := path.Join(path.Dir(flags.Input), "log.txt")
	//
	fmt.Println("logfile ", logfile)
}
