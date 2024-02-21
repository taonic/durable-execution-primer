package main

import (
	"context"
	"log"

	"go.temporal.io/sdk/client"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatal(err)
	}

	// Start workflow
	log.Print("Running workflow")
	run, err := c.ExecuteWorkflow(
		context.Background(),
		client.StartWorkflowOptions{
			ID:        "my-workflow-1",
			TaskQueue: "my-task-queue",
		},
		"MyGreeter",
		"hello",
	)
	if err != nil {
		log.Fatal(err)
	}

	// Get result
	var result string
	err = run.Get(context.Background(), &result)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Exec result: %v\n", result)
}
