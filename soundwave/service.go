package soundwave

import (
	"fmt"
	"log"
	"os/exec"
	"path"
	"strings"

	"github.com/HiWay-Media/SoundWave-Go/types"
	"github.com/Paxx-RnD/go-ffmpeg/ffprobe"
)

type IService interface {
	GenerateWaveForm() string
}

type service struct {
	flags   types.Flags
	ffprobe ffprobe.IFfprobe
}

func NewService(flags types.Flags, ffprobe ffprobe.IFfprobe) IService {
	return &service{flags: flags, ffprobe: ffprobe}
}

func (s *service) GenerateWaveForm() string {
	probe, err := s.ffprobe.GetProbe(s.flags.Input)
	//
	if err != nil {
		panic(fmt.Sprintf("could not probe: %v", err))
	}
	log.Println("probe ", probe)
	/*duration, err := strconv.ParseFloat(probe.GetFirstVideoStream().Duration, 64)
	if err != nil {
		panic(fmt.Sprintf("could not get video duration: %v", err))
	}*/
	outputPath := path.Join(path.Dir(s.flags.Output), "output.png") 
	log.Println(outputPath)
	//ffmpeg -i input -filter_complex "compand,showwavespic=s=640x120" -frames:v 1 output.png
	cmd := exec.Command(
		"ffmpeg",
		"-i",
		s.flags.Input,
		"-filter_complex",
		fmt.Sprintf("compand,showwavespic=%dx%d", s.flags.Width, s.flags.Height),
		"-frames:v",
		"1",
		outputPath,
	)
	_, err = cmd.CombinedOutput()
	fmt.Println(strings.Join(cmd.Args, " "))
	if err != nil {
		panic(fmt.Sprintf("failed to extract frame: %v", err))
	}
	return outputPath
}
