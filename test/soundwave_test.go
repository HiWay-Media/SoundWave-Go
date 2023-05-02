package soundwave

import (
	"log"
	"testing"

	"github.com/HiWay-Media/SoundWave-Go/soundwave"
	"github.com/HiWay-Media/SoundWave-Go/types"
	"github.com/Paxx-RnD/go-ffmpeg/configuration"
	"github.com/Paxx-RnD/go-ffmpeg/ffprobe"
)

var (
	inputPath = "input.mp3"
)

func Configuration() *configuration.Configuration {
	return &configuration.Configuration{
		FfprobePath: "ffprobe",
	}
}

func TestNewService(t *testing.T) {
	flags := getFlags()
	ffprobe := ffprobe.NewFfProbe(Configuration(), nil)
	log.Println("ffprobe ", ffprobe, " flags ", flags)
	s := soundwave.NewService(flags, ffprobe)
	assert.NotNil(t, s)
}

func getFlags() types.Flags {
	flags := types.Flags{
		Input: inputPath,
		//Prefix: outputPath,
	}
	return flags
}
