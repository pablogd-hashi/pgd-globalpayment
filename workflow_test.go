package app

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/testsuite"
)

func Test_SuccessfulTransferWorkflow(t *testing.T) {
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()

	testDetails := PaymentDetails{
		SourceAccount: "85-150",
		TargetAccount: "43-812",
		Amount:        250,
		ReferenceID:   "12345",
	}

	// Mock activity implementation
	env.OnActivity(Withdraw, mock.Anything, testDetails).Return("", nil)
	env.OnActivity(Deposit, mock.Anything, testDetails).Return("", nil)

	env.ExecuteWorkflow(MoneyTransfer, testDetails)
	require.True(t, env.IsWorkflowCompleted())
	require.NoError(t, env.GetWorkflowError())
}

func Test_InvalidAmountRejectedBeforeAnyActivity(t *testing.T) {
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()

	testDetails := PaymentDetails{
		SourceAccount: "85-150",
		TargetAccount: "43-812",
		Amount:        -25,
		ReferenceID:   "12345",
	}

	// Deliberately no env.OnActivity(...) mocks are registered for Withdraw,
	// Deposit, or Refund. If the Workflow scheduled any of them, the test
	// environment would fail with an unexpected-call error instead of the
	// assertions below, so a clean pass here is evidence that no
	// money-movement Activity ran.
	env.ExecuteWorkflow(MoneyTransfer, testDetails)

	require.True(t, env.IsWorkflowCompleted())
	require.Error(t, env.GetWorkflowError())

	// Temporal wraps the Workflow's returned error as an ApplicationError
	// keyed by its Go type name before it reaches the client/test harness.
	var appErr *temporal.ApplicationError
	require.ErrorAs(t, env.GetWorkflowError(), &appErr)
	require.Equal(t, "InvalidAmountError", appErr.Type())
}

func Test_ZeroAmountRejected(t *testing.T) {
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()

	testDetails := PaymentDetails{
		SourceAccount: "85-150",
		TargetAccount: "43-812",
		Amount:        0,
		ReferenceID:   "12345",
	}

	env.ExecuteWorkflow(MoneyTransfer, testDetails)

	require.True(t, env.IsWorkflowCompleted())
	require.Error(t, env.GetWorkflowError())
}

func Test_DepositFailedWorkflow(t *testing.T) {
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()

	testDetails := PaymentDetails{
		SourceAccount: "85-150",
		TargetAccount: "43-812",
		Amount:        250,
		ReferenceID:   "12345",
	}

	// Mock activity implementation
	env.OnActivity(Withdraw, mock.Anything, testDetails).Return("", nil)
	env.OnActivity(Deposit, mock.Anything, testDetails).Return("", errors.New("unable to deposit"))
	env.OnActivity(Refund, mock.Anything, testDetails).Return("", nil)

	env.ExecuteWorkflow(MoneyTransfer, testDetails)
	require.True(t, env.IsWorkflowCompleted())
	require.Error(t, env.GetWorkflowError())
}
