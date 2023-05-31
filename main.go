package main

import "fmt"

type Task struct {
	Name string
	Flag bool
}

type TodoList struct {
	Tasks []Task
}

func (tl *TodoList) AddTask(name string) {
	task := Task{Name: name, Flag: false}
	tl.Tasks = append(tl.Tasks, task)
	fmt.Println("Task add:", name)
}

func (tl *TodoList) CompletTask(index int) {
	if index >= 0 && index < len(tl.Tasks) {
		tl.Tasks[index].Flag = true
		fmt.Println("Task complete:", tl.Tasks[index].Name)
	} else {
		fmt.Println("Try againe")
	}
}

func (tl *TodoList) ShowTask() {
	fmt.Println("Your Task:")
	for i, task := range tl.Tasks {
		status := " "
		if task.Flag {
			status = "+"
		}
		fmt.Printf("%d. [%s] %s\n", i+1, status, task.Name)
	}
}

func main() {
	todoList := TodoList{}
	fmt.Println("1. Add task")
	fmt.Println("2. Complete task")
	fmt.Println("3. Print tasks")
	fmt.Println("0. Exit")

	var choise int
	fmt.Println("yor choise:")
	fmt.Scan(&choise)

	switch choise {
	case 0:
		fmt.Println("...")
		return
	case 1:
		fmt.Print("Enter task name: ")
		var name string
		fmt.Scanln(&name)
		todoList.AddTask(name)
	case 2:
		fmt.Print("Enter task index: ")
		var index int
		fmt.Scanln(&index)
		todoList.CompletTask(index - 1)
	case 3:
		todoList.ShowTask()
	default:
		fmt.Println("Invalid choice")
	}

	fmt.Println()
}
