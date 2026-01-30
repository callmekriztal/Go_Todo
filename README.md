# Go_Todo


A simple command-line Todo List application written in Go.
This tool lets you add, edit, delete, toggle, and list tasks directly from your terminal.
All todos are persisted locally in a JSON file inside your home directory.

<h3>Features<h3>

<ul>
<li>Add new tasks</li>
<li>Edit existing tasks</li>
<li>Delete tasks</li>
<li>Mark tasks as completed</li>
<li>List all tasks in a table format</li>
<li>Persistent storage using JSON</li>
</ul>

Installation
Prerequisites

Go 1.20+ installed

Clone the repository
git clone https://github.com/yourusername/todocli.git
cd todocli

Build the binary
go build -o todocli


(Optional) Move it to your PATH:

sudo mv todocli /usr/local/bin/

Usage

Run the application using command-line flags.

Add a new task
todocli -add "Buy groceries"

List all tasks
todocli -list

Edit a task
todocli -edit 0:Buy milk and bread


Format:

-edit <task_index>:<new_title>

Toggle task completion
todocli -toggle 0

Delete a task
todocli -del 0