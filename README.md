<!-- PROJECT LOGO -->
<br />
<p align="center">
  <a href="https://github.com/hackstream/zettel">
    <img src="./docs/cover.png" alt="Logo">
  </a>

  <p align="center">
    <br />
    <a href="https://zettel.hackstream.dev/"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/hackstream/zettel/issues">Report Bug</a>
    ·
    <a href="https://github.com/hackstream/zettel/issues">Request Feature</a>
  </p>
</p>


`zettel` is a simple FOSS tool to jot down notes using Zettelkasten methodology. The core aim of `zettel` is to store notes _for life_. The notes are in simple `markdown` format which can be published as a static website.

## Features

- Generate static website using `zettel build`.
- No external dependency of any app or SASS vendor lock in for your notes. `zettel` operates on plain markdown files.
- Generate **Connections** between notes with a simple syntax.
- Visualise all the connections with a Graph UI.
- No `$EDITOR` dependency. The markdown notes work same everywhere!

## Why would I use this?

### Host your own data

If you don't want to lock in your **precious** data with a SAAS provider. You don't have to trust this tool to last forever - for your data to outlast you.
`zettel` uses the notes from local disk and generates connections out of it which can be published as a static webpage.


### Opinionated Zettelkasten Workflow

If you want an opinionated yet productive workflow.This tool is quite blunt about doing just one thing - taking notes using Zettelkasten philosophy. This is not a generic tool to take notes although you can use as one.

The real power of this tool lies in the UX while creating and **connecting** the notes - it helps you organize your notes in a structured format and visualise the _graph_ structure of your notes.


## Philosophy

Our thinking process isn't linear but most of the note taking tools _coerce_ you into structuring your notes in a hierarchical form. Zettelkasten takes a completely different approach: There's no strict hierarchy and any note can be "linked" to another note. Think of it like your **second brain**.


A connection can be created in a note using `[[slug]]` syntax. The slug of the markdown file is replaced with the link to the note and a "connection" is established between these 2 notes. 

In the graph UI terminology, the 2 notes act as the **vertex** of the graph connected by an **edge**.

## Installation

### Grab the latest binary

```shell
$ cd "$(mktemp -d)"
$ curl -sL "https://github.com/hackstream/zettel/releases/download/0.1.0/zettel_0.1.0_$(uname)_amd64.tar.gz" | tar xz
$ mv zettel /usr/local/bin
# zettel should be available now in your $PATH
$ zettel --version
```

## Usage

```shell
NAME:
   zettel - Zettel builds a digital Zettelkasten website for your notes in Markdown.

USAGE:
   zettel.bin [global options] command [command options] [arguments...]

VERSION:
   a5fd18b (2020-09-13 11:51:10 +0530)

AUTHOR:
   Hackstream Devs

COMMANDS:
   init, i   Initializes a new zettel site with default config.
   new, n    Create a new post.
   build, b  Builds a static dist of all notes ready to be published on web.
   help, h   Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --verbose      Enable verbose logging (default: false)
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)
```

### Initialise a new project

`zettel init $SITENAME`: Creates a new `$SITENAME` folder which holds `zettel` config files and a default `index.md`.

### Create a new post

`zettel new $TITLE`: Creates a new `$TITLE.md` inside `content/` directory with the current date in metadata.

### Build website

`zettel build`: Runs a pipeline to iterate over all markdown files in `content/*.md`, create connections across posts and output a `dist` folder with the static assets. This folder contains `index.html` which becomes the root of your website.

## Configuration

A default `config.toml` is created for you when you initialise the site. The following options can be edited:

- `site_name`: Site name for your website.
- `description`: A short one liner description for your website.

## ⭐️ Show your support

Give a ⭐️ if this project helped you!

## Contribution

PRs on Feature Requests, Bug fixes are welcome. Feel free to open an issue and have a discussion first. Read [CONTRIBUTING.md](CONTRIBUTING.md) for more details.

For a list of things to improve visit [TODO.md](TODO.md)

## License

[GPL v3](license)
