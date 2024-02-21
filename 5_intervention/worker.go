package main

import (
	"log"
	"fmt"
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

	var greeting string

	for i := 0; i < 30; i++ {
		workflow.Sleep(ctx, 1*time.Second)
		greeting = fmt.Sprintf("--- %s world %v ---", hello, i)
		if i == 10 {
			var resume bool
			chn := workflow.GetSignalChannel(ctx, "resume")
			chn.Receive(ctx, &resume)
		}
		logger.Info(greeting)
	}

	return greeting, nil
}
