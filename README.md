# GoTodo

A simple command-line Todo List application written in Go.

GoTodo lets you manage tasks directly from your terminal and persists them locally using JSON.

## Features

* Add tasks
* List tasks in a formatted table
* Edit tasks
* Delete tasks
* Toggle task completion
* Clear the entire task list
* Persistent JSON storage
* Task creation and completion timestamps
* Input validation and error handling
* Unit tests

## Storage

Tasks are stored locally at:

```text
~/.todocli/todos.json
```

No database or external service is required.

## Installation

### Linux — AppImage

Download `GoTodo-x86_64.AppImage` from the GitHub Releases page.

Make it executable:

```bash
chmod +x GoTodo-x86_64.AppImage
```

Run it:

```bash
./GoTodo-x86_64.AppImage --list
```

The AppImage does not require installation or Go to be installed.

### Linux — Standalone Binary

Download `todocli-linux-amd64` from the GitHub Releases page.

Make it executable:

```bash
chmod +x todocli-linux-amd64
```

For system-wide installation, move it to `/usr/local/bin`:

```bash
sudo install -m 755 todocli-linux-amd64 /usr/local/bin/todocli
```

You can now run it from anywhere:

```bash
todocli --list
```

Check that it is available:

```bash
which todocli
```

Expected:

```text
/usr/local/bin/todocli
```

### Build from Source

Requirements:

* Go 1.20 or newer
* Linux, macOS, or Windows

Clone the repository and build:

```bash
git clone https://github.com/callmekriztal/Go_Todo.git
cd Go_Todo
go build -trimpath -ldflags="-s -w" -o todocli .
```

Install system-wide on Linux:

```bash
sudo install -m 755 todocli /usr/local/bin/todocli
```

Verify:

```bash
todocli --help
```

## Usage

### Add a task

```bash
todocli --add "Learn Go"
```

### List tasks

```bash
todocli --list
```

### Edit a task

Task indexes start at `1`.

```bash
todocli --edit 1:"Learn advanced Go"
```

Format:

```text
--edit <task_index>:<new_title>
```

### Toggle task completion

```bash
todocli --toggle 1
```

Running the command again marks the task as incomplete.

### Delete a task

```bash
todocli --del 1
```

### Clear all tasks

```bash
todocli --clear
```

This permanently removes every task from the current todo list.

### Show help

```bash
todocli --help
```

## Example

```bash
$ todocli --add "Learn Go"
$ todocli --add "Build a backend"
$ todocli --list
```

Example output:

```text
┌─────────┬─────────────────┬───────────┬───────────────────────────────┬──────────────┐
│ Task No │      Title      │ Completed │          Created At           │ Completed At │
├─────────┼─────────────────┼───────────┼───────────────────────────────┼──────────────┤
│ 1       │ Learn Go        │ ✖️         │ Mon, 24 Aug 2026 19:00:00 IST │              │
│ 2       │ Build a backend │ ✖️         │ Mon, 24 Aug 2026 19:01:00 IST │              │
└─────────┴─────────────────┴───────────┴───────────────────────────────┴──────────────┘
```

Complete a task:

```bash
todocli --toggle 1
```

## Testing

Run the complete test suite:

```bash
go test ./...
```

Run tests with verbose output:

```bash
go test -v
```

Run static analysis:

```bash
go vet ./...
```

## Project Structure

```text
Go_Todo/
├── command.go       # CLI flags and command execution
├── helper.go        # Todo operations and display
├── storage.go       # JSON persistence
├── main.go          # Application entry point
├── helper_test.go   # Unit tests
├── go.mod
├── go.sum
└── README.md
```

## Release

Prebuilt Linux binaries are available through GitHub Releases:

* `GoTodo-x86_64.AppImage` — portable Linux AppImage
* `todocli-linux-amd64` — standalone Linux binary

The AppImage can be run without installation. The standalone binary can be installed system-wide using `/usr/local/bin`.

## License

MIT License

