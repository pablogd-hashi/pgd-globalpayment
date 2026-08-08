# Repository conventions

Durable context for engineers and agents working in this repository. Keep
this file short; add a new section only when a task discovers a convention
future work should not have to rediscover.

## Where business-rule guards belong

**Convention:** a validation rule that (a) is deterministic, (b) requires no
external call (no bank/network/database lookup), and (c) must prevent a
money-movement side effect from starting, belongs in the **Workflow
function**, before the first `workflow.ExecuteActivity` call it guards — not
inside the Activity, and not in the client that starts the Workflow.

Why:

- **Not in the Activity** (`banking-client.go` / `activity.go`): by the time
  an Activity runs, the side effect it performs has already been scheduled
  and is visible in Workflow history. A guard placed there can only stop the
  *next* Activity, not the one it's inside of, and duplicating the guard
  across every money-movement Activity invites drift. `Withdraw` and
  `Deposit` already validate what they alone know how to validate (account
  existence, balance) via `InvalidAccountError` / `InsufficientFundsError`.
- **Not in the starter** (`start/main.go` or any other future Workflow
  starter): client-side validation can be bypassed by any other caller of
  the Workflow, and a rejection that never reaches Temporal leaves no
  Workflow history — no audit trail, nothing to show in Temporal Web,
  nothing a test against the Workflow can assert on.
- **In the Workflow**, before scheduling the Activity: the guard is
  covered by a Workflow test (see `workflow_test.go`), runs before any
  Activity is scheduled (so Temporal Web shows zero Activity events for a
  rejected request), and is deterministic replay-safe since it only reads
  the Workflow input.

## Testing this convention

A Workflow-level guard should have a Workflow test (using
`testsuite.WorkflowTestSuite`) that asserts two things together:

1. `env.GetWorkflowError()` reports the expected rejection.
2. No `env.OnActivity(...)` mock is registered for the Activities the guard
   should prevent. If the Workflow scheduled one anyway, the test
   environment fails on the unexpected call instead of passing — this is
   what proves the side effect never started, not just that the final
   error message looks right.

See `Test_InvalidAmountRejectedBeforeAnyActivity` in `workflow_test.go`.
