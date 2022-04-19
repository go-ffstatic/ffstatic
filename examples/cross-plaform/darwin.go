//go:build darwin

package main

import (
	ff "github.com/go-ffstatic/darwin-amd64"
)

var ffmpegPath = ff.FFmpegPath()
var ffprobePath = ff.FFprobePath()
