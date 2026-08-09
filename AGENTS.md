# Agent orientation

This is a Temporal Go money-transfer codebase: a Workflow withdraws from one
account and deposits into another. Orchestration lives in `workflow.go`;
side-effecting Activities live in `activity.go`. Run `task verify` for
`go vet` and `go test ./...`. Temporal Web: http://localhost:8233
