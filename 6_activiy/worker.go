package main

import (
	"context"
	"log"
	"strconv"
	"time"

	"math/rand"

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
	w.RegisterActivity(Greet)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatal(err)
	}
}

func Greet(ctx context.Context, input, index string) (string, error) {
	if rand.Intn(2) == 1 {
		time.Sleep(5 * time.Second)
	}

	return input + " world " + index, nil
}

func MyGreeter(ctx workflow.Context, hello string) (string, error) {
	var greeting string

	for i := 0; i < 10; i++ {
		workflow.Sleep(ctx, 1*time.Second)

		ao := workflow.ActivityOptions{StartToCloseTimeout: 2 * time.Second}
		ctx = workflow.WithActivityOptions(ctx, ao)
		err := workflow.ExecuteActivity(ctx, Greet, hello, strconv.Itoa(i+1)).Get(ctx, &greeting)
		if err != nil {
			return "", err
		}
	}

	return greeting, nil
}
