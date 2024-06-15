# faux-tileserver

This project demonstrates a simple tile server implemented using Go Fiber and the gg library. The server generates map tiles dynamically, displaying the X, Y, and Z coordinates.

This is especially useful if you're building an offline mapping application and need a self-hosted tile server. Since all of the tiles are generated dynamically, this application is incredibly lightweight and uses very little resources.

## Getting Set Up

### Building/Running From Source 

- Ensure you have Go v1.22.3+ installed
- Install dependencies with the command `go mod tidy`
- Run the application with `go run .` (don't forget the `.`)
  - You can specify a custom port by appending `-port=XXXX`, where `XXXX` is the port number
  - For example, to run on port 3333, you would type `go run . -port=3333`
- Visit the app at [localhost:3000](http://localhost:3000) (or whatever port you specified while running)

#### Air

You can use [Air](https://github.com/air-verse/air) for live reloading during development. Simply install Air with the following command:

```sh
go install github.com/air-verse/air@latest
```

and then you can type `air` in your terminal to run the application.

## Downloading Pre-Built Binaries

You can find pre-built fauxts binaries for Windows, Linux, and MacOS on the faux-tileserver repo's [releases page](https://github.com/jere-mie/faux-tileserver/releases/latest) From there, you can download the appropriate binary and either add it to your system's PATH variable, or run it directly from whatever directory you place it in.

If you prefer downloading via the cli, use one of the following commands below:

```sh
# Windows amd64
irm -Uri https://github.com/jere-mie/faux-tileserver/releases/latest/download/fauxts_windows_amd64.exe -O fauxts.exe

# Linux amd64
curl -L https://github.com/jere-mie/faux-tileserver/releases/latest/download/fauxts_linux_amd64 -o fauxts && chmod +x fauxts

# MacOS arm64 (Apple Silicon)
curl -L https://github.com/jere-mie/faux-tileserver/releases/latest/download/fauxts_darwin_arm64 -o fauxts && chmod +x fauxts

# MacOS amd64 (Intel)
curl -L https://github.com/jere-mie/faux-tileserver/releases/latest/download/fauxts_darwin_amd64 -o fauxts && chmod +x fauxts
```

### Specifying Custom Port

When running, you can specify a custom port by appending `-port=XXXX` in the CLI, where `XXXX` is your desired port. Thus, the command to run the application on port 3333 would look something like:

```sh
./fauxts -port=3333
```
