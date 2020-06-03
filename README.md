# WebGoCat
A command line WebSocket interface with no Node dependency and SSL built in.

Written in Go.

## Installation
- Download the binary that works for your machine from the [releases](https://github.com/elliotaplant/webgocat/releases) page
  - Don't see your OS there? Make an issue requesting support for that OS
- Make sure the binary is executable (`chmod +x /path/to/webgocat` if it isn't)
- Put the binary somewhere in your `$PATH`

OR
- Clone the repo and build it from source (`go build -o webgocat main.go`)

## What does it do?
`webgocat` is a WebSocket client for the command line. It lets you connect to a WebSocket to send and receive text messages.

![](https://media.giphy.com/media/f8ah2wX7WzMZOhiCTq/giphy.gif)

## Why?
I love the tool [`wscat`](https://github.com/websockets/wscat) and have used it for a long time, but I work on machines that don't have Node/NPM installed.
Because `wscat` is installed via NPM and run on Node, it isn't an option for those projects.

I tried using [`websocat`](https://github.com/vi/websocat), but the binaries don't come with SSL built in so I couldn't connect to `wss` endpoints. One could probably add SSL when compiling from the source, but that seemed onerous.

Go lets you compile a binary for almost any target, so I thought that use Go would be a great tool to make a WebSocket client without requiring any other dependencies on the host.

## Like the tool but want more features?
Please create an issue!

## Similar projects
- [wscat](https://github.com/websockets/wscat)
- [websocat](https://github.com/vi/websocat)

## License
MIT
