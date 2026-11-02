# Task-CLI

A simple command-line task manager built in Go.

This is a personal learning project I'm using to familiarise myself with **Go's fundamentals**, coming from a Java development background.

The goal isn't to build a production-ready application, but to explore Go's syntax, conventions, standard library and approach to software development.

## What I'm Learning

- Go modules and package organisation
- Structs, methods and pointer receivers
- Slices and collection management
- Error handling
- CLI input and control flow
- JSON serialisation and file I/O
- Unit testing with Go's standard library

## Features

- Add tasks
- List existing tasks
- Mark tasks as complete
- Delete tasks
- Save and load tasks using JSON

## Getting Started

**Build the application:**

```bash
go build -o task-cli ./cmd
```

**Run it:**

```bash
./task-cli
```

## Commands

| Command | Description |
|---|---|
| `add` | Add a new task |
| `list` | List all tasks |
| `complete` | Mark a task as complete |
| `delete` | Delete a task |
| `save` | Save tasks to a JSON file |
| `exit` | Exit the application |

## Disclaimer

This project is purely for educational purposes. I'm learning Go as I go, so the implementation may change as I discover better and more idiomatic ways of doing things.