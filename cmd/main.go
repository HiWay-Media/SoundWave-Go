package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/HiWay-Media/SoundWave-Go/soundwave"
	"github.com/HiWay-Media/SoundWave-Go/types"
	"github.com/Paxx-RnD/go-ffmpeg/configuration"
	"github.com/Paxx-RnD/go-ffmpeg/ffprobe"
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
	var conf = configuration.Configuration{
		FfmpegPath:  "ffmpeg",
		FfprobePath: "ffprobe",
	}
	ffprobe := ffprobe.NewFfProbe(&conf, nil)
	log.Println("ffprobe ", ffprobe)
	//
	s := soundwave.NewService(flags, ffprobe)
	fmt.Println("Generating soundwave ", s)
}
