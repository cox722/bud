# Go-Fullstack-Cox

The Full-Stack Web Framework for Go. Go-Fullstack-Cox writes the boring code for you, helping you launch your website faster.

## Documentation

Read [the documentation](https://denim-cub-301.notion.site/Hey-Go-Fullstack-Cox-4d81622cc49942f9917c5033e5205c69) to learn how to get started with Go-Fullstack-Cox.

# Installing Go-Fullstack-Cox

Go-Fullstack-Cox ships as a single binary that runs on Linux and Mac. You can follow along for Windows support in [this issue](https://github.com/cox722/go-fullstack-cox/issues/7).

The easiest way to get started is by copying and pasting the command below in your terminal:

```diff
curl -sf https://raw.githubusercontent.com/cox722/go-fullstack-cox/main/install.sh | sh
```

This script will download the right binary for your operating system and move the binary to the right location in your `$PATH`.

Confirm that you've installed Go-Fullstack-Cox by typing `Go-Fullstack-Cox` in your terminal.

```bash
Go-Fullstack-Cox -h
```

You should see the following:

```bash
Usage:
    Go-Fullstack-Cox [flags] [command]

Flags:
  -C, --chdir  Change the working directory

Commands:
  build    build the production server
  create   create a new project
  run      run the development server
  tool     extra tools
  version  Show package versions
```

# Requirements

The following software is required to use Go-Fullstack-Cox.

- Node v14+

  This is a temporary requirement that we plan to remove in [v0.3](https://github.com/cox722/go-fullstack-cox/discussions/21)

- Go v1.17+

  Go-Fullstack-Cox relies heavily on `io/fs` and will take advantage of generics in the future, so while Go v1.16 will work, we suggest running Go v1.18+ if you can.

# Your First Project

With Go-Fullstack-Cox installed, you can now scaffold a new project:

```bash
$ Go-Fullstack-Cox create hello
$ cd hello
```

The create command will scaffold everything you need to get started with Go-Fullstack-Cox.

```bash
$ ls
go.mod  node_modules/  package-lock.json  package.json
```

... which is not very much by the way! Unlike most other fullstack frameworks, Go-Fullstack-Cox starts out very minimal. As you add dependencies, Go-Fullstack-Cox will generate all the boring code to glue your app together. Let's see this in action.

Start the development server with `Go-Fullstack-Cox run`:

```bash
$ Go-Fullstack-Cox run
| Listening on http://127.0.0.1:3000
```

Click on the link to open the browser. You'll be greeted with Go-Fullstack-Cox's welcome page.

Congrats! You're running your first web server with Go-Fullstack-Cox. The welcome server is your jumping off point to learn more about the framework.

- Go-Fullstack-Cox should compile to a single binary that contains your entire web app and can be copied to a server that doesn't even have Go installed.

# Contributing

Please refer to the [Contributing Guide](./contributing/Readme.md) to learn how to develop Go-Fullstack-Cox locally.
