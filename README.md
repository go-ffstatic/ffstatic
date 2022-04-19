# ffstatic

Embed ffmpeg and ffprobe static binaries into your go
program and use them without preinstalled binaries.

This project is inspired
by [ffmpeg-static](https://github.com/eugeneware/ffmpeg-static)
.

## How it works

We use go embed to embed the ffmpeg/ffprobe binaries as
bytes in go program. The real ffmpeg/ffprobe executable
files would be created in system temp directory when you run
the program.

With this package, you can use ffmpeg/ffprobe without
preinstalling. It would be very helpful to run program on
serverless platform (ie. lambda\Vercel).

## Installation

Choose the version matched with your environment.

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

## How to use

## Cross compiling

## License

MIT