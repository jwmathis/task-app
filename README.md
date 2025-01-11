# Task-App in Go

This is a simple command-line application for managing tasks, written in Go. This app allows you to add, 
view, update, and delte tasks. It uses basic file I/O for data persistence and offeres an easy-to-use 
interface for managing your task list.

## Features
* Add a new task: Add a task with a description using simple commands
* View tasks: Quickly display all tasks with their details
* Update tasks: Simple to modify existing tasks
* Delete tasks: Remove tasks from your list
* Data persistence: Tasks are saved and persist across app restarts

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
go build -o task-app
```
2. Afterwards, use the following command to run the app:

```
./task-app <command> [options]
```
### Available Commands
1. *Add Task* Add a new task to the list
```
task-app add "Task description"
```
2. *View Tasks* View all tasks
```
task-app view
```
3. *Mark task complete* Mark task done
```
task-app done <task-id>
```
4. *Delete task* Delete task by ID
```
task-app del <task-id>
```