# minimalist-go-react-app

## Pre-requisites

Install Go, Node.js, and NPM.

## Development

Run with `scripts/run`. This will start the Go server at 8080 and build the React UI app. Changes to templates, static files, and ui files
can be instantly see by cmd+R in the browser.

To make changes to in .go files, stop the server and run `scripts/run` again.

## Build

Run `scripts/build` which will produce a single binary files in the `build` directory. This binary for your platform and share with people on the same platform.
