# minimalist-go-react-app

You can build a web app with Go server and React UI with just three dependencies. You can distribute the app as a single executable binary, while still having a great development experience. This repository demonstrates how to do it.

![](/static/logo.svg)

If you enjoy the minimalist style, use this repository as a template for your own app. You can replace React with any other
UI framework, or install additional dependencies.

## Prerequisites

Install Go and Node.js.

```
brew install go
brew install node
```

## Development

Run the app with `scripts/run`. Once the server starts, open `http://localhost:8080` in your browser.

You can instantly view changes to files in the `static`, `templates`, and `ui` directories by refreshing the page. To view changes to `.go` files, stop the server and run `scripts/run` again.

## Build

Run `scripts/build` which will produce a single executable binary file `minimalist-go-react-app` in the `build` directory. By default, the binary is built for your current operating system and architecture. You can build for other operating systems and architectures like this:

```
# macOS Apple Silicon
GOOS=darwin GOARCH=arm64 scripts/build

# macOS Intel Chip
GOOS=darwin GOARCH=amd64 scripts/build

# Windows x64
GOOS=windows GOARCH=amd64 scripts/build
```

You can find all available combinations of `GOOS` and `GOARCH` in the [official documentation](https://go.dev/doc/install/source#environment).

## Docker

Although the goal is to demonstrate the single executable binary distribution, Docker is a good choice for cloud deployment. `Dockerfile` is included.

If you have Docker installed, you can run `scripts/docker` to build and run the Docker image.
