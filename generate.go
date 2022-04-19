package main

import (
	"fmt"
	"html/template"
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

var builds = []struct {
	System string
	Arch   string
}{
	{System: "darwin", Arch: "amd64"},
	{System: "darwin", Arch: "arm64"},
	{System: "windows", Arch: "amd64"},
	{System: "windows", Arch: "386"},
	{System: "linux", Arch: "amd64"},
	{System: "linux", Arch: "arm64"},
	{System: "linux", Arch: "arm"},
	{System: "linux", Arch: "386"},
	{System: "freebsd", Arch: "amd64"},
}

const tag = "b4.4.0-rc.11"

var tpls = template.Must(template.ParseGlob("templates/*.tpl"))

func genGoFile(system, arch string) error {
	for _, t := range tpls.Templates() {
		targetFile := fmt.Sprintf("%s-%s/%s", system, arch, strings.Trim(t.Name(), ".tpl"))
		os.Remove(targetFile)

		file, err := os.OpenFile(targetFile, os.O_CREATE|os.O_WRONLY, os.ModePerm)
		if err != nil {
			return fmt.Errorf("open file %s: %v", t.Name(), err)
		}
		if err := t.Execute(file, map[string]interface{}{
			"system": system,
			"arch":   arch,
		}); err != nil {
			return fmt.Errorf("execute template %s: %v", t.Name(), err)
		}
	}
	return nil
}

func download(kind, system, arch string) error {
	targetFile := system + "-" + arch + "/" + kind
	os.Remove(targetFile)

	switch arch {
	case "386":
		arch = "ia32"
	case "amd64":
		arch = "x64"
	}

	switch system {
	case "windows":
		system = "win32"
	}

	downloadURL := fmt.Sprintf("https://github.com/descriptinc/ffmpeg-ffprobe-static/releases/download/%s/%s-%s-%s", tag, kind, system, arch)

	resp, err := http.Get(downloadURL)
	if err != nil {
		return fmt.Errorf("failed to download %s: %v", downloadURL, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	file, err := os.OpenFile(targetFile, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	if _, err := io.Copy(file, resp.Body); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}
	return nil
}

func main() {
	var g sync.WaitGroup

	for _, build := range builds {
		system, arch := build.System, build.Arch
		for _, kind := range kinds {
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
		if err := genGoFile(system, arch); err != nil {
			panic(fmt.Errorf("failed to generate go file: %w", err))
		}
	}
	g.Wait()
}
