package main

import (
	"fmt"
	"strconv"
)

// Constants  \\
const totalTasks int = 100

const (
	LowPriority = iota
	MediumPriority
	HighPriority
)

func main() {
	// Variables \\
	var projectStatus string = "IN PROGRESS"
	tasksCreated := 25
	var projectCompleted bool

	// Convert boolean to string \\
	projectCompletedStr := strconv.FormatBool(projectCompleted)

	// Output  \\
	fmt.Println("Welcome to the Task Management System!")
	fmt.Println("Project start date is: 2024-09-18 00:00:00")
	fmt.Println("Project: Task Management System\n")
	fmt.Printf("Current project status: %s\n", projectStatus)
	fmt.Printf("Tasks completed: %d out of %d\n", tasksCreated, totalTasks)
	fmt.Printf("Task priorities: %d-Low, %d-Medium, %d-High\n", LowPriority, MediumPriority, HighPriority)
	fmt.Printf("Is the project completed? %s\n", projectCompletedStr)
}
