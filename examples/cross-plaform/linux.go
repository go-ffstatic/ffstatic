//go:build linux

package main

import (
	ff "github.com/go-ffstatic/linux-amd64"
)

var ffmpegPath = ff.FFmpegPath()
var ffprobePath = ff.FFprobePath()
