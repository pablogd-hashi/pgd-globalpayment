package main

import (
	"context"
	"flag"
	"log"

	"go.temporal.io/sdk/client"

	"money-transfer-project-template-go/app"
)

// @@@SNIPSTART money-transfer-project-template-go-start-workflow
func main() {
	// Flags let the demo start the same Workflow with different inputs
	// (e.g. `task demo-invalid`, `task demo-valid`) without editing code.
	workflowID := flag.String("id", "pay-invoice-701", "Workflow ID")
	sourceAccount := flag.String("source", "85-150", "Source account number")
	targetAccount := flag.String("target", "43-812", "Target account number")
	amount := flag.Int("amount", 250, "Transfer amount")
	referenceID := flag.String("ref", "12345", "Reference ID")
	flag.Parse()

	// Create the client object just once per process
	c, err := client.Dial(client.Options{})

	if err != nil {
		log.Fatalln("Unable to create Temporal client:", err)
	}

	defer c.Close()

	input := app.PaymentDetails{
		SourceAccount: *sourceAccount,
		TargetAccount: *targetAccount,
		Amount:        *amount,
		ReferenceID:   *referenceID,
	}

	options := client.StartWorkflowOptions{
		ID:        *workflowID,
		TaskQueue: app.MoneyTransferTaskQueueName,
	}

	log.Printf("Starting transfer from account %s to account %s for %d", input.SourceAccount, input.TargetAccount, input.Amount)

	we, err := c.ExecuteWorkflow(context.Background(), options, app.MoneyTransfer, input)
	if err != nil {
		log.Fatalln("Unable to start the Workflow:", err)
	}

	log.Printf("WorkflowID: %s RunID: %s\n", we.GetID(), we.GetRunID())

	var result string

	err = we.Get(context.Background(), &result)

	if err != nil {
		log.Println("Workflow returned an error:", err)
		return
	}

	log.Println(result)
}

// @@@SNIPEND
