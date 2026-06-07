# Cox's Contributor Guide

Thank you for your interest in helping to make Cox better! Cox is and forever will be a driven by the community. We depend on volunteers lending their time and expertise to make Cox better for everyone.

## Requirements

- OSX, Linux or Windows (via [WSL2](https://github.com/cox722/go-fullstack-cox/issues/7))
- Node 16+
- Go 1.17+
- C++ compiler in your $PATH (for cgo to compile V8)

## Setting up Cox for Development

Run the following commands to download and run Cox locally:

```sh
git clone https://github.com/cox722/go-fullstack-cox
cd Cox
make # fresh installs take a few minutes because of V8
go run main.go
```

After running `go run main.go`, you should see the following:

```
  Usage:
    Cox [flags] [command]

  Flags:
    -C, --chdir  Change the working directory

  Commands:
    build    build the production server
    create   create a new project
    run      run the development server
    tool     extra tools
    version  Show package versions
```

If you run into any problems, please [open an issue](https://github.com/cox722/go-fullstack-cox/issues/new).

### Developing with Docker

We've included a [Dockerfile](./Dockerfile) that you can use to run Cox within a Docker container. This can be helpful for testing Cox within a Linux environment and debugging CI issues.

To build and start a Docker container with Cox, run the following commands:

```sh
docker build -t Cox:latest contributing
docker run -it --rm -v $(pwd):/Cox Cox /bin/bash
```

To build a Docker image based on another Node.js and/or Go version, run one of the following commands:

```shell
docker build --build-arg NODE_VERSION=18.4.0 -t Cox:latest contributing
docker build --build-arg GO_VERSION=1.18.2 -t Cox:latest contributing
docker build --build-arg NODE_VERSION=18.4.0 --build-arg GO_VERSION=1.18.2 -t Cox:latest contributing
```

## Developing Cox with a Project

Now that you have Cox running locally, you can use the `-C, --chdir` functionality to test Cox against different projects.

You can use one of the projects in `example/`

```sh
# Run the development server on the example/hn application
go run main.go -C example/hn run

# Build a binary for the example/hn application
go run main.go -C example/hn build
```

Or you can create your own project:

```sh
# Scaffold a new hello application
go run main.go create hello

# Run the development server for the hello application
go run main.go -C hello run

# Build a binary for the hello application
go run main.go -C hello build
```

## Disabling Code Generation

Sometimes you want to hack on the generated code in `Cox/`, but if you run `Cox run` any changes you make will be overridden. To avoid this, you can call your app's `main.go` file directly:

```sh
go run Cox/cmd/app/main.go
```

You may encounter an error like this:

```sh
Coxhttp: discard client does not support render
```

This occurs because apps built with `Cox run` depends on a Cox server to provide hot reloads and client-side file bundling. To start a Cox server, run the following:

```sh
Cox tool bs
```

This will start a Cox server. Next, restart your app server and pass the Cox server address in as an environment variable:

```
Cox_LISTEN=<Cox server address> go run Cox/cmd/app/main.go
```

Finally, reload the page and you should be good to go. Happy hacking!

## Issues to Work On

Issues with the [good first issue](https://github.com/cox722/go-fullstack-cox/issues?q=is%3Aissue+is%3Aopen+sort%3Aupdated-desc+label%3A%22good+first+issue%22) or [help wanted](https://github.com/cox722/go-fullstack-cox/issues?q=is%3Aissue+is%3Aopen+label%3A%22help+wanted%22) labels would be good candidates to work on.

## Running Tests

You can use the following commands to run all the tests:

```sh
make test
```

## Publishing a new Version

To publish a new version of Cox, you will need your a clean local main branch that matches the remote main branch. You'll also need to have commit access on both Github and publish rights on NPM.

You can run then run:

1. Write the `Changelog.md`. You can use `git log v0.1.0..HEAD` to see the commits since the last release.
2. Bump the version in `version.txt`.
3. Run `make publish`.

## Note for Developing with Cox on Windows

Currently, Cox must be developed using WSL/WSL2 if developing on Windows. While there is a desire to get native Windows development supported, the effort to get there will be a HEAVY lift, as we will need to work closely with the [v8go](https://github.com/rogchap/v8go) team to get their Windows support added back in to v8go, as well as a handful of other large effort-intensive tasks, such as finding a consistent way to identify file descriptors in Windows. There are some efforts working towards accomplishing these goals, but as they are such large tasks, it has been decided that they will be placed on the back burner. If anyone wants to contribute to these efforts, you can find a beginning list of tasks [here](https://github.com/cox722/go-fullstack-cox/discussions/81).
