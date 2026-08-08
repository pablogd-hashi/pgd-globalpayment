package app

import (
	"fmt"
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

// InvalidAmountError is returned when a transfer is requested with an amount
// that is not strictly greater than zero. It is a Workflow-level input guard,
// not a banking error, so it lives here rather than in banking-client.go.
type InvalidAmountError struct{}

func (m *InvalidAmountError) Error() string {
	return "Transfer amount must be greater than zero"
}

// @@@SNIPSTART money-transfer-project-template-go-workflow
func MoneyTransfer(ctx workflow.Context, input PaymentDetails) (string, error) {

	// Reject invalid transfer amounts before any money-movement Activity is
	// scheduled. This check is deterministic and needs no external call, so
	// it belongs in the Workflow rather than in an Activity: it stops the
	// transfer before Withdraw/Deposit ever run, instead of after. See
	// .cursor/rules/guard-placement.mdc for where this class of guard belongs.
	if input.Amount <= 0 {
		return "", &InvalidAmountError{}
	}

	// RetryPolicy specifies how to automatically handle retries if an Activity fails.
	retrypolicy := &temporal.RetryPolicy{
		InitialInterval:        time.Second,
		BackoffCoefficient:     2.0,
		MaximumInterval:        100 * time.Second,
		MaximumAttempts:        500, // 0 is unlimited retries
		NonRetryableErrorTypes: []string{"InvalidAccountError", "InsufficientFundsError"},
	}

	options := workflow.ActivityOptions{
		// Timeout options specify when to automatically timeout Activity functions.
		StartToCloseTimeout: time.Minute,
		// Optionally provide a customized RetryPolicy.
		// Temporal retries failed Activities by default.
		RetryPolicy: retrypolicy,
	}

	// Apply the options.
	ctx = workflow.WithActivityOptions(ctx, options)

	// Withdraw money.
	var withdrawOutput string

	withdrawErr := workflow.ExecuteActivity(ctx, Withdraw, input).Get(ctx, &withdrawOutput)

	if withdrawErr != nil {
		return "", withdrawErr
	}

	// Deposit money.
	var depositOutput string

	depositErr := workflow.ExecuteActivity(ctx, Deposit, input).Get(ctx, &depositOutput)

	if depositErr != nil {
		// The deposit failed; put money back in original account.

		var result string

		refundErr := workflow.ExecuteActivity(ctx, Refund, input).Get(ctx, &result)

		if refundErr != nil {
			return "",
				fmt.Errorf("Deposit: failed to deposit money into %v: %v. Money could not be returned to %v: %w",
					input.TargetAccount, depositErr, input.SourceAccount, refundErr)
		}

		return "", fmt.Errorf("Deposit: failed to deposit money into %v: Money returned to %v: %w",
			input.TargetAccount, input.SourceAccount, depositErr)
	}

	result := fmt.Sprintf("Transfer complete (transaction IDs: %s, %s)", withdrawOutput, depositOutput)
	return result, nil
}

// @@@SNIPEND
