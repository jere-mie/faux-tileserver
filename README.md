# faux-tileserver

A map tile server that procedurally generates map tiles, useful for testing mapping applications

## Air

You can use [Air](https://github.com/air-verse/air) for live reloading during development. Simply install Air with the following command:

```sh
go install github.com/air-verse/air@latest
```

and then you can type `air` in your terminal to run the application.

## Downloading

You can find pre-built fauxts binaries for Windows, Linux, and MacOS on the faux-tileserver repo's [releases page](https://github.com/jere-mie/faux-tileserver/releases/latest) From there, you can download the binaries and add them to your system's PATH variable.

If you prefer downloading via the cli, you can use the following command to download the latest fauxts binary on **Windows** (amd64):

```sh
irm -Uri https://github.com/jere-mie/faux-tileserver/releases/latest/download/fauxts_windows_amd64.exe -O fauxts.exe
```

the following command on **Linux** (amd64):

```sh
curl -L https://github.com/jere-mie/faux-tileserver/releases/latest/download/fauxts_linux_amd64 -o fauxts && chmod +x fauxts
```

the following on **MacOS** (arm64, Apple Silicon):

```sh
curl -L https://github.com/jere-mie/faux-tileserver/releases/latest/download/fauxts_darwin_arm64 -o fauxts && chmod +x fauxts
```

and the following on **MacOS** (amd64, Intel):

```sh
curl -L https://github.com/jere-mie/faux-tileserver/releases/latest/download/fauxts_darwin_amd64 -o fauxts && chmod +x fauxts
```
