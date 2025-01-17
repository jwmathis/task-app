# Task-App in Go

This is a simple command-line application for managing tasks, written in Go. This app allows you to add, 
view, update, and delete tasks. It uses basic file I/O for data persistence and offers an easy-to-use 
interface for managing your task list from within the command line utility.

## Features
* Add a new task: Add a task with a description using simple commands
* View tasks: Quickly display all tasks with their details.
* Update tasks: Simple to modify existing tasks
* Delete tasks: Remove tasks from your list
* Data persistence: Tasks are saved and persist across app restarts
* When a task is marked done, it turns green. When it is marked todo, it turns red

## Installation
Build Task-App from source:

### Prerequisites

Install the latest version of Go installed. You can download and install it from the official [Go website](https://go.dev/).

1. Clone the repository

Download the project files from GitHub:
```
git clone https://github.com/jwmathis/task-app.git
```
2. Navigate to repository location
```
cd task-app
```

3. Build the binary
```
go build -o task
```
4. Move the binary to your preferred location
```
mv task.exe C:\Tools
```
5. Add to PATH (optional)
    1. Press Win + S, type "Environment Variables", and select "Edit the system environment variables".
    2. In the System Properties window, click the Environment Variables button.
    3. Under system variables, find and select the Path variable, then click Edit.
    4. Click New and enter the directory containing task.exe (e.g. C\:Tools)
    5. Click OK to save and close all windows.
6. Test the installation: Open a new terminal window and type
```
task --help
```
If installed correctly, the help text for your task app will display

## Usage

### Available Commands
1. *Add Task* Add a new task to the list
```
task add "Task description"
```
2. *View Tasks* View all tasks
```
task view
```
3. *Mark task complete* Mark task done
```
task mark done <task-id>
```
4. *Mark task todo* Mark task as todo
```
task mark todo <task-id>
```
5. *Delete task* Delete task by ID
```
task del <task-id>
```