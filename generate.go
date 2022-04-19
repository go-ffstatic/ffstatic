package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
)

var kinds = []string{
	"ffmpeg",
	"ffprobe",
}
var systems = []string{
	"darwin",
	"freebsd",
	"win32",
	"linux",
}
var arches = []string{
	"arm64",
	"ia32",
	"arm",
	"x64",
}

const pkg = "ffstatic"
const tag = "b4.4.0-rc.11"

func download(kind, system, arch string) error {
	fileName := fmt.Sprintf("%s_%s_%s", kind, system, arch)
	constraint := []string{}
	switch arch {
	case "ia32":
		constraint = []string{"386"}
	case "x64":
		constraint = []string{"amd64"}
	default:
		constraint = append(constraint, arch)
	}
	switch system {
	case "win32":
		constraint = append(constraint, "windows")
	default:
		constraint = append(constraint, system)
	}
	{
		downloadURL := fmt.Sprintf("https://github.com/descriptinc/ffmpeg-ffprobe-static/releases/download/%s/%s-%s-%s", tag, kind, system, arch)

		resp, err := http.Get(downloadURL)
		if err != nil {
			return fmt.Errorf("failed to download %s: %v", downloadURL, err)
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusNotFound {
			return nil
		}
		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
		}

		file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, os.ModePerm)
		if err != nil {
			return fmt.Errorf("failed to open file: %w", err)
		}
		defer file.Close()

		if _, err := io.Copy(file, resp.Body); err != nil {
			return fmt.Errorf("failed to write file: %w", err)
		}
	}
	{
		file, err := os.OpenFile(fileName+".go", os.O_CREATE|os.O_WRONLY, os.ModePerm)
		if err != nil {
			return fmt.Errorf("failed to open file: %w", err)
		}
		defer file.Close()

		if _, err := fmt.Fprintf(file, `//go:generate go run ./generate
//go:build %s
// +build %s

package %s

import (
	_ "embed"
)

//go:embed %s
var %sBinary []byte
`,
			strings.Join(constraint, " && "),
			strings.Join(constraint, ","),
			pkg,
			fileName,
			kind,
		); err != nil {
			return fmt.Errorf("failed to write file: %v", err)
		}
	}
	return nil
}

func main() {
	var g sync.WaitGroup

	for _, kind := range kinds {
		for _, system := range systems {
			for _, arch := range arches {
				kind, system, arch := kind, system, arch
				g.Add(1)
				go func() {
					defer g.Done()
					err := download(kind, system, arch)
					if err != nil {
						fmt.Printf("Fail to download %s-%s-%s: %s\n", kind, system, arch, err.Error())
					}
				}()
			}
		}
	}
	g.Wait()
}
