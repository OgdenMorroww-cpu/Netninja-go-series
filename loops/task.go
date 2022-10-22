package loops

import "fmt"

func GetTasks() {
	todo_tasks := []string{
		"Wash the car",
		"Take out trash",
		"Go to gym",
		"Record a video tutorials",
		"Walk my dog",
		"Book my flight",
		"Check for black adam movie tickets",
		"Go to the pub",
		"Visit ho chi ming",
		"Vote peter obi",
		"Code for 4 hours",
		"Check my trading earnings",
		"Take a course on docker",
		"Deploy my api to GCP",
	}

	for todo_index, todo_items := range todo_tasks {
		fmt.Printf("My Task No %v %v\n", todo_index, todo_items)
	}
}
