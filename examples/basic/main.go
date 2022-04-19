package main

import (
	"github.com/floostack/transcoder/ffmpeg"
	ffstatic_darwin_amd64 "github.com/go-ffstatic/darwin-amd64"
)

func main() {
	ff := ffmpeg.New(&ffmpeg.Config{
		FfmpegBinPath:  ffstatic_darwin_amd64.FFmpegPath(),
		FfprobeBinPath: ffstatic_darwin_amd64.FFprobePath(),
	})
	ch, err := ff.Input("./input.avi").Output("./output.mp4").Start(ffmpeg.Options{})
	if err != nil {
		panic(err)
	}
	for range ch {
	}
}
