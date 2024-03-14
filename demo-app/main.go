package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	streamdal "github.com/streamdal/streamdal/sdks/go"
)

func main() {
	// The example payload containing an email address, which needs to be scrubbed before we generate a sales report
	payloadWithPII := []byte(`{
		"customer": {
			"first_name": "John",
			"last_name": "Doe",
			"email": "john.doe@streamdal.com"
		}
	}`)

	// Initialize the StreamDAL SDK
	sc, _ := streamdal.New(&streamdal.Config{
		ServerURL:       "localhost:8082",
		ServerToken:     "1234",
		ServiceName:     "billing-svc",
		PipelineTimeout: time.Second,
	})

	// Give the SDK some time to pull rules since we're not using the SDK in a server app
	time.Sleep(3 * time.Second)

	// Process the payload using the audience configuration we defined in our terraform file
	resp := sc.Process(context.Background(), &streamdal.ProcessRequest{
		OperationType: streamdal.OperationTypeConsumer,
		OperationName: "gen-sales-report",
		ComponentName: "kafka",
		Data:          payloadWithPII,
	})

	// Ensure the pipeline was successful
	if resp.Status != streamdal.ExecStatusTrue {
		fmt.Printf("Error: %s\n", *resp.StatusMessage)
		return
	}

	// Pretty print our modified JSON, showing the email is now masked
	if err := prettyPrint(resp.Data); err != nil {
		fmt.Printf("Error pretty printing JSON: %s\n", err)
	}
}

func prettyPrint(data []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	pretty, err := json.MarshalIndent(tmp, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(pretty))
	return nil
}
