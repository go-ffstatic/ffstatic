# ffstatic

Embed ffmpeg and ffprobe static binaries into your go program and use them without preinstalled binaries.

This project is inspired by [ffmpeg-static](https://github.com/eugeneware/ffmpeg-static).

## How it works

We use go embed to embed the ffmpeg/ffprobe binaries as bytes in the go program. The real ffmpeg/ffprobe executable files would be created in the temp directory when you run the program.

With this package, you can use ffmpeg/ffprobe without preinstalling. It would be very helpful to run programs on serverless platforms (ie. lambda/Vercel).

## Installation

Select the version that matches your environment:

- Darwin
  - amd64

    ```sh
    go get -u github.com/go-ffstatic/darwin-amd64
    ```

  - arm64

    ```sh
    go get -u github.com/go-ffstatic/darwin-arm64
    ```

- Linux
  - amd64

    ```sh
    go get -u github.com/go-ffstatic/linux-amd64
    ```

  - arm64

    ```sh
    go get -u github.com/go-ffstatic/linux-arm64
    ```

  - arm

    ```sh
    go get -u github.com/go-ffstatic/linux-arm
    ```

  - 386

    ```sh
    go get -u github.com/go-ffstatic/linux-386
    ```

- Windows
  - amd64

    ```sh
    go get -u github.com/go-ffstatic/windows-amd64
    ```

  - 386
  
    ```sh
    go get -u github.com/go-ffstatic/windows-386
    ```

- FreeBSD
  - amd64

    ```sh
    go get -u github.com/go-ffstatic/freebsd-amd64
    ```

## Basic Usage

You can use `os/exec` to execute ffmpeg/ffprobe directly. I also recommend using ffmpeg wrappers, for example:

- [xfrr/goffmpeg](https://github.com/xfrr/goffmpeg)

- [floostack/transcoder](https://github.com/floostack/transcoder)

- [u2takey/ffmpeg-go](https://github.com/u2takey/ffmpeg-go)

- [AlexEidt/Vidio](https://github.com/AlexEidt/Vidio)

Check [examples/basic](./examples/basic) for more details.

## Cross-platform

If you want to run your program on multiple platforms (dev on darwin, deploy on linux, etc.), you need to use `go:build` constrain for cross compile.

Check [examples/cross-plaform](./examples/cross-plaform) for more details.

## License

MIT
