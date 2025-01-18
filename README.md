# Task-App in Go

This is a simple command-line application for managing tasks, written in Go. This app allows you to add, 
view, update, and delete tasks. It uses basic file I/O for data persistence and offers an easy-to-use 
interface for managing your task list.

## Features
* Add a new task: Add a task with a description using simple commands
* View tasks: Quickly display all tasks with their details.
* Update tasks: Simple to modify existing tasks
* Delete tasks: Remove tasks from your list
* Data persistence: Tasks are saved and persist across app restarts
* When a task is marked done, it turns green. When it is marked todo, it turns red

## Installation

### Prerequisites

Make sure you have the latest version of Go installed. You can download and install it from the official [Go website](https://go.dev/).

1. Clone the repository

Download the project files from GitHub:
```
git clone https://github.com/jwmathis/task-app.git
```
2. Navigate to repository location
```
cd task-app
```

## Usage

### Running the app
1. To start the task app, run:

```
go build -o task.exe
```
2. Afterwards, use the following command to run the app:

```
./task <command> [options]
```
### Available Commands
1. *Add Task* Add a new task to the list
```
./task add "Task description"
```
2. *View Tasks* View all tasks
```
./task view
```
3. *Mark task complete* Mark task done
```
./task mark done <task-id>
```
4. *Mark task todo* Mark task as todo
```
./task mark todo <task-id>
```
5. *Delete task* Delete task by ID
```
./task del <task-id>
```