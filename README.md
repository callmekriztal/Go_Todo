<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Go_Todo</title>
</head>
<body>

<h1>Go_Todo</h1>

<p>
A simple command-line Todo List application written in Go.<br>
This tool lets you add, edit, delete, toggle, and list tasks directly from your terminal.<br>
All todos are persisted locally in a JSON file inside your home directory.
</p>

<hr>

<h2>Features</h2>
<ul>
    <li>Add new tasks</li>
    <li>Edit existing tasks</li>
    <li>Delete tasks</li>
    <li>Mark tasks as completed</li>
    <li>List all tasks in a table format</li>
    <li>Persistent storage using JSON</li>
</ul>

<hr>

<h2>Installation</h2>

<h3>Prerequisites</h3>
<ul>
    <li>Go 1.20 or higher</li>
</ul>

<h3>Clone the repository</h3>
<pre><code>git clone https://github.com/yourusername/todocli.git
cd todocli</code></pre>

<h3>Build the binary</h3>
<pre><code>go build -o todocli</code></pre>

<h3>(Optional) Move binary to PATH</h3>
<pre><code>sudo mv todocli /usr/local/bin/</code></pre>

<hr>

<h2>Usage</h2>

<p>Run the application using command-line flags.</p>

<h3>Add a new task</h3>
<pre><code>todocli -add "Buy groceries"</code></pre>

<h3>List all tasks</h3>
<pre><code>todocli -list</code></pre>

<h3>Edit a task</h3>
<pre><code>todocli -edit 0:Buy milk and bread</code></pre>

<p><strong>Format:</strong></p>
<pre><code>-edit &lt;task_index&gt;:&lt;new_title&gt;</code></pre>

<h3>Toggle task completion</h3>
<pre><code>todocli -toggle 0</code></pre>

<h3>Delete a task</h3>
<pre><code>todocli -del 0</code></pre>

<hr>

<h2>Data Storage</h2>
<p>Todos are stored locally in the following file:</p>
<pre><code>~/.todocli/todos.json</code></pre>

<p>The directory is created automatically if it does not exist.</p>

<hr>

<h2>License</h2>
<p>MIT License</p>

</body>
</html>
