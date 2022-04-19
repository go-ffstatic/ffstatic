package main

import (
	"github.com/floostack/transcoder/ffmpeg"
)

func main() {
	ff := ffmpeg.New(&ffmpeg.Config{
		FfmpegBinPath:  ffmpegPath,
		FfprobeBinPath: ffprobePath,
	})
	ch, err := ff.Input("./input.avi").Output("./output.mp4").Start(ffmpeg.Options{})
	if err != nil {
		panic(err)
	}
	for range ch {
	}
}
