# Agent orientation

This repository is a small Temporal Go demo: a money-transfer Workflow that
withdraws from one account and deposits into another. Temporal orchestrates
retries and compensation; it is the runtime, not the subject of changes here.

## Layout

| Path | Role |
|------|------|
| `workflow.go` | Workflow definition — orchestration and activity scheduling |
| `activity.go` | Activity implementations (`Withdraw`, `Deposit`, `Refund`) |
| `banking-client.go` | Mock banking client used by activities |
| `shared.go` | Shared types and the task queue name |
| `start/main.go` | Starts a Workflow run (CLI flags for accounts and amount) |
| `worker/main.go` | Worker that executes Workflows and Activities |
| `workflow_test.go` | Workflow unit tests (Temporal test suite) |

## Running locally

Use the Taskfile from the repo root:

- `task start` — Temporal dev server, Web UI, and worker
- `task verify` — `go vet` and `go test ./...`
- `task demo-invalid` / `task demo-valid` — prepared demo runs

Temporal Web: http://localhost:8233

Read the files above before proposing changes. Keep diffs minimal and run
`task verify` before treating work as done.
