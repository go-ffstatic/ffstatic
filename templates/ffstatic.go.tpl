//generated by github.com/go-ffstatic/ffstatic
//go:build {{.system}} && {{.arch}}
// +build {{.system}},{{.arch}}

package ffstatic_{{.system}}_{{.arch}}

import (
	_ "embed"
    "fmt"
    "os"
)

//go:embed ffmpeg
var ffmpeg []byte

//go:embed ffprobe
var ffprobe []byte

func writeTempExec(pattern string, binary []byte) (string, error) {
	f, err := os.CreateTemp("", pattern)
	if err != nil {
		return "", fmt.Errorf("failed to create temp file: %v", err)
	}
	defer f.Close()
	_, err = f.Write(binary)
	if err != nil {
		return "", fmt.Errorf("fail to write executable: %v", err)
	}
	if err := f.Chmod(os.ModePerm); err != nil {
		return "", fmt.Errorf("fail to chmod: %v", err)
	}
	return f.Name(), nil
}

var (
	FfmpegPath  string
	FfprobeEPath string
)

func init() {
	var err error
	FfmpegPath, err = writeTempExec("ffmpeg_{{.system}}_{{.arch}}", ffmpeg)
	if err != nil {
		panic(fmt.Errorf("failed to write ffmpeg_{{.system}}_{{.arch}}: %v", err))
	}
	FfprobePath, err = writeTempExec("ffprobe_{{.system}}_{{.arch}}", ffprobe)
	if err != nil {
		panic(fmt.Errorf("failed to write ffprobe_{{.system}}_{{.arch}}: %v", err))
	}
}
