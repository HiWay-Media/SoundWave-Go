package soundwave

import (
	"github.com/Paxx-RnD/go-ffmpeg/ffprobe"
	"github.com/HiWay-Media/SoundWave-Go/types"
)

type IService interface {

}

type service struct {
	flags   types.Flags
	ffprobe ffprobe.IFfprobe
}


func NewService(flags types.Flags, ffprobe ffprobe.IFfprobe) IService {
	return &service{flags: flags, ffprobe: ffprobe}
}