package hls_test

import (
	"os"
	"path"
	"testing"

	"github.com/erzed/go-hls-transcoder"
)

func TestCmdExecuteFfmpeg(t *testing.T) {
	base, _ := os.Getwd()

	targetPath := path.Join(base, "static")
	srcPath := path.Join(base, "static", "sample.mov")
	ffmpegPath := "/usr/local/bin/ffmpeg"

	err := hls.GenerateHLS(ffmpegPath, srcPath, targetPath, "480p", true)
	if err != nil {
		panic(err)
	}
}
