package main

import (
	"flag"
	"fmt"
	"path"

	"github.com/Paxx-RnD/go-ffmpeg/configuration"
	"github.com/Paxx-RnD/go-ffmpeg/ffprobe"
	"github.com/HiWay-Media/SoundWave-Go/types"
	"github.com/HiWay-Media/SoundWave-Go/soundwave"
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
	log.Println("ffprobe ", ffprobe )
	//
	s := soundwave.NewService(flags, ffprobe)
	fmt.Println("Generating soundwave ", s)
}
