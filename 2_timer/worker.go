package main

import (
	"log"
	"time"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatal(err)
	}

	// Start worker
	log.Print("Start worker")
	w := worker.New(c, "my-task-queue", worker.Options{StickyScheduleToStartTimeout: 5 * time.Millisecond})
	w.RegisterWorkflow(MyGreeter)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatal(err)
	}
}

func MyGreeter(ctx workflow.Context, hello string) (string, error) {
	logger := workflow.GetLogger(ctx)

	greeting := hello + " world"

	logger.Info("--- Start sleep ---")
	workflow.Sleep(ctx, 10*time.Second)
	logger.Info("--- Done sleep ---")

	return greeting, nil
}
