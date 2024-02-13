# Todo List CLI

Todo List CLI is a command-line interface (CLI) tool for managing your todo list. It allows you to add, mark tasks as complete, delete tasks, and view your todo list from the terminal.

## Installation

Clone the repository:

```bash
git clone https://github.com/Kunal-deve1oper/todo-list-cli.git
```

Navigate to the project directory:

```bash
cd todo-list-cli
```

Build the project:

```bash
go build
```

## Usage

The CLI provides the following commands:

### Adding a Task

To add a new task to your todo list, use the `-a` flag followed by the task description:

```bash
./todo-list-cli -a "Task description"
```

### Completing a Task

To mark a task as complete, use the `-c` flag followed by the index of the task you want to mark as complete. The index refers to the position of the task in the todo list. For example, to mark the first task as complete:

```bash
./todo-list-cli -c 1
```

This command will mark the task at index 1 as complete.

### Deleting a Task

To delete a task from your todo list, use the `-d` flag followed by the index of the task you want to delete. Similar to completing a task, the index refers to the position of the task in the todo list. For example, to delete the second task:

```bash
./todo-list-cli -d 2
```

This command will delete the task at index 2 from your todo list.

### Viewing Tasks

To view all tasks in your todo list, use the `-v` flag:

```bash
./todo-list-cli -v
```

This will display all tasks currently in your todo list.

### Flags

- `-a`: Add a new task to the todo list.
- `-c`: Mark a task as complete by providing its index.
- `-d`: Delete a task from the todo list by providing its index.
- `-v`: View tasks in the todo list.

## Examples

- Adding a task:

```bash
./todo-list-cli -a "Buy groceries"
```

- Marking a task as complete:

```bash
./todo-list-cli -c 1
```

- Deleting a task:

```bash
./todo-list-cli -d 2
```

- Viewing tasks:

```bash
./todo-list-cli -v
```

## Contributing

Contributions are welcome! If you'd like to contribute to this project, please fork the repository and submit a pull request.