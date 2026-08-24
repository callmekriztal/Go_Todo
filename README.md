<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
</head>
<body>

<h1>Go_Todo</h1>

<p>
A simple command-line Todo List application written in Go.
</p>

<p>
Go_Todo lets you manage tasks directly from your terminal and
persists them locally using JSON.
</p>

<hr>

<h2>Features</h2>

<ul>
    <li>Add tasks</li>
    <li>List all tasks</li>
    <li>Edit tasks</li>
    <li>Delete tasks</li>
    <li>Toggle task completion</li>
    <li>Clear the entire task list</li>
    <li>Persistent JSON storage</li>
    <li>Task creation and completion timestamps</li>
    <li>Input validation and error handling</li>
    <li>Unit tests</li>
</ul>

<hr>

<h2>Storage</h2>

<p>
Tasks are stored locally at:
</p>

<pre><code>~/.todocli/todos.json</code></pre>

<p>
The application does not require a database or external service.
</p>

<hr>

<h2>Installation</h2>

<h3>Prerequisites</h3>

<ul>
    <li>Go 1.20 or newer</li>
    <li>Linux, macOS, or Windows for building from source</li>
</ul>

<h3>Clone the repository</h3>

<pre><code>git clone https://github.com/callmekriztal/Go_Todo.git
cd Go_Todo</code></pre>

<hr>

<h2>Build from Source</h2>

<h3>Linux</h3>

<pre><code>go build -trimpath -ldflags="-s -w" -o todocli .</code></pre>

<p>Run:</p>

<pre><code>./todocli --help</code></pre>

<h3>Windows</h3>

<pre><code>go build -o todocli.exe .</code></pre>

<h3>Install on Linux</h3>

<pre><code>sudo install -m 755 todocli /usr/local/bin/todocli</code></pre>

<p>Then:</p>

<pre><code>todocli --help</code></pre>

<hr>

<h2>Usage</h2>

<h3>Add a task</h3>

<pre><code>todocli --add "Learn Go"</code></pre>

<h3>List tasks</h3>

<pre><code>todocli --list</code></pre>

<h3>Edit a task</h3>

<pre><code>todocli --edit 1:"Learn advanced Go"</code></pre>

<p>
Task indexes start at <strong>1</strong>.
</p>

<h3>Toggle completion</h3>

<pre><code>todocli --toggle 1</code></pre>

<p>
Running the command again marks the task as incomplete.
</p>

<h3>Delete a task</h3>

<pre><code>todocli --del 1</code></pre>

<h3>Clear all tasks</h3>

<pre><code>todocli --clear</code></pre>

<p>
This removes every task from the current todo list.
</p>

<h3>Show help</h3>

<pre><code>todocli --help</code></pre>

<hr>

<h2>Testing</h2>

<p>Run the test suite:</p>

<pre><code>go test ./...</code></pre>

<p>Run tests with verbose output:</p>

<pre><code>go test -v</code></pre>

<p>Run static analysis:</p>

<pre><code>go vet ./...</code></pre>

<hr>

<h2>Project Structure</h2>

<pre><code>Go_Todo/
├── command.go       # CLI flags and command execution
├── helper.go        # Todo operations and display
├── storage.go       # JSON persistence
├── main.go          # Application entry point
├── helper_test.go   # Unit tests
├── go.mod
├── go.sum
└── README.md</code></pre>

<hr>

<h2>License</h2>

<p>MIT License</p>

</body>
</html>
