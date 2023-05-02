package soundwave

import (
	"github.com/Paxx-RnD/go-ffmpeg/configuration"
)


var (
	inputPath      = "input.mp3"
)


func Configuration() *configuration.Configuration {
	return &configuration.Configuration{
		FfprobePath: "ffprobe",
	}
}
//
func TestNewService(t *testing.T) {
	flags := getFlags()
	ffprobe := ffprobe.NewFfProbe(Configuration(), nil)
	/*s := sprite.NewService(flags, ffprobe)
	assert.NotNil(t, s)*/
}


func getFlags() types.Flags {
	flags := types.Flags{
		Input:     inputPath,
		Prefix:    outputPath,
	}
	return flags
}