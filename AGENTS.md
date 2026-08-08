# AGENTS.md

## Cursor Cloud specific instructions

This is a Temporal "money transfer" demo written in Go (module `money-transfer-project-template-go/app`). It has three runtime pieces:

- Temporal dev server + Web UI (`temporal server start-dev`, Web UI on port `8233`, gRPC on `7233`)
- A Workflow/Activity worker (`go run ./worker`) that polls the `TRANSFER_MONEY_TASK_QUEUE`
- A starter (`go run ./start`) that kicks off a `MoneyTransfer` workflow

The bank in `banking-client.go` is fully mocked (no real network calls), so no external services or secrets are needed.

### Tooling / PATH (non-obvious)

- The update script installs the `temporal` CLI to `~/.temporalio/bin` and `go-task` to `~/go/bin`. Neither directory is on the default `PATH`. A `PATH` export for both was added to `~/.bashrc` during setup, so interactive shells pick them up. If a command reports `temporal: command not found` or `task: command not found` in a non-interactive shell, run: `export PATH="$PATH:/home/ubuntu/.temporalio/bin:/home/ubuntu/go/bin"`.
- Go itself auto-downloads the toolchain pinned in `go.mod` (`toolchain go1.24.7`) on first invocation.

### Running the app

The `Taskfile.yml` is the intended dev entrypoint (requires `go-task`):

- `task start` — starts the dev server + Web UI (if not already up) and the worker (both via `nohup`, logs under `.demo/`).
- `task reload-worker` — restart the worker after editing `workflow.go` / `activity.go` (the worker does NOT hot-reload; you must restart it to pick up source changes).
- `task demo-valid` / `task demo-invalid` — reload the worker and start a sample transfer.
- `task stop` — stop the worker and dev server.
- `task verify` — runs `go vet ./...` and `go test ./... -v`.

Equivalent raw commands (no `go-task` needed):

- Server: `temporal server start-dev --db-filename .demo/temporal.db --ui-port 8233`
- Worker: `go run ./worker`
- Start a workflow: `go run ./start -id=<id> -source=85-150 -target=43-812 -amount=250 -ref=<ref>`

### Notes

- The dev server DB and logs live in `.demo/` (gitignored). The Web UI is at http://localhost:8233.
- On the `lab/base` branch the workflow does NOT reject non-positive amounts; a negative `-amount` still completes. That is expected on the base branch (the rejection logic is the learner exercise), not a bug.
