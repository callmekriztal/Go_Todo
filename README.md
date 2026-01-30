<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
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

<hr>

<h3>Prerequisites</h3>
<ul>
    <li>Go 1.20 or higher</li>
    <li>Need a linux system</li>
</ul>

<h3>Clone the repository</h3>
<pre><code>git clone https://github.com/callmekriztal/Go_Todo.git
cd Go_Todo</code></pre>

<hr>

<h2>Build Instructions</h2>

<h3>Linux</h3>

<h4>Build for current Linux system</h4>
<pre><code>go build -o todocli</code></pre>

<h4>Build static Linux binary (recommended for servers / Docker)</h4>
<pre><code>CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o todocli</code></pre>

<h4>(Optional) Move binary to PATH</h4>
<pre><code>sudo mv todocli /usr/local/bin/</code></pre>

<hr>

<h3>Windows</h3>

<h4>Build for Windows (native)</h4>
<pre><code>go build -o todocli.exe</code></pre>

<h4>Run the application</h4>
<pre><code>todocli.exe -list</code></pre>

<h4>(Optional) Add to PATH</h4>
<p>
Move <code>todocli.exe</code> to a directory included in your Windows PATH
(e.g. <code>C:\Windows\System32</code> or a custom tools folder).
</p>

<hr>

<h3>Cross-compile Linux binary from Windows</h3>
<p>You can build a Linux binary directly from Windows without Docker or WSL.</p>

<h4>PowerShell</h4>
<pre><code>$env:CGO_ENABLED=0
$env:GOOS="linux"
$env:GOARCH="amd64"
go build -o todocli</code></pre>

<h4>CMD</h4>
<pre><code>set CGO_ENABLED=0
set GOOS=linux
set GOARCH=amd64
go build -o todocli</code></pre>

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

<p><strong>Example:</strong></p>
<pre><code>-edit 1:"Study Go"</code></pre>

<h3>Toggle task completion</h3>
<pre><code>todocli -toggle 0</code></pre>

<h3>Delete a task</h3>
<pre><code>todocli -del 0</code></pre>

<<<<<<< HEAD


<hr>

=======
<div>
    <h3>This is version one and next step will be to:-</h3> 
    <ul>
    <li>1.include storage or daily todos that will be saved automatically each day </li>
    <li>2.Add priority coloum for each task</li>
    <li>3.clear completed task</li>
    </ul>
</div>
<hr>
>>>>>>> 8477b1c (updated the readme)
<h2>License</h2>
<p>MIT License</p>

</body>
</html>
