#!/usr/bin/env bash
# Idempotent environment setup for the pgd-globalpayment agentic SDLC lab.
# Installs the Go toolchain deps plus the CLIs the lab and presentation need,
# then makes them discoverable on PATH for future shells.
set -euo pipefail

export PATH="$HOME/.local/bin:$HOME/go/bin:$HOME/.temporalio/bin:$PATH"

echo "==> Go modules"
go mod download

echo "==> Temporal CLI (dev server + Web UI)"
if ! command -v temporal >/dev/null 2>&1 && [ ! -x "$HOME/.temporalio/bin/temporal" ]; then
  curl -sSf https://temporal.download/cli.sh | sh
fi

echo "==> go-task (Taskfile runner)"
if ! command -v task >/dev/null 2>&1 && [ ! -x "$HOME/go/bin/task" ]; then
  go install github.com/go-task/task/v3/cmd/task@latest
fi

echo "==> Cursor CLI (cursor-agent)"
if ! command -v cursor-agent >/dev/null 2>&1 && [ ! -x "$HOME/.local/bin/cursor-agent" ]; then
  curl https://cursor.com/install -fsS | bash
fi

# Persist PATH for interactive shells (idempotent).
if ! grep -q 'temporalio/bin' "$HOME/.bashrc" 2>/dev/null; then
  echo 'export PATH="$HOME/.local/bin:$HOME/go/bin:$HOME/.temporalio/bin:$PATH"' >> "$HOME/.bashrc"
fi

echo "==> Build check"
go build ./...

echo "Environment ready. Tools:"
command -v go && go version
command -v task && task --version
command -v temporal && temporal --version
command -v cursor-agent && cursor-agent --version || true
