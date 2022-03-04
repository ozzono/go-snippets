// Sample stdlogging writes log.Logger logs to the Stackdriver Logging.
package main

import (
	"context"
	"log"

	"cloud.google.com/go/logging"
)

func main() {
	ctx := context.Background()

	// Sets your Google Cloud Platform project ID.
	projectID := "ifbra-go"

	// Creates a client.
	client, err := logging.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	// Sets the name of the log to write to.
	logName := "my-log"

	logger := client.Logger(logName).StandardLogger(logging.Info)

	// Logs "hello world", log entry is visible at
	// Stackdriver Logs.
	logger.Println("hello world")
}
