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

	// Signal workflow
	err = c.SignalWorkflow(context.Background(), "my-workflow-1", "", "resume", nil)
	if err != nil {
		log.Fatalln("Unable to signal workflow", err)
	}

	log.Printf("Signal sent")
}
