# pgd-globalpayment

A teaching lab for **agentic SDLC patterns**: how to take a ticket from an
incomplete request to a verified, reviewable change, with human gates between
phases. Skills, rules, and committed artefacts are the curriculum. The host
application is only a small payments-shaped codebase so side effects are
observable.

You do not need to master Temporal to complete the lab. Temporal here is the
runtime that shows whether money-moving work actually ran.

- Architecture (problem, approach, artefacts): [docs/architecture.md](docs/architecture.md)
- Step-by-step lab (self-serve, same path as the demo): [LAB.md](LAB.md)
- 90-minute workshop outline: [MODULE_OUTLINE.md](MODULE_OUTLINE.md)
- Facilitator 15-minute proof cue sheet: [RUNBOOK.md](RUNBOOK.md)

## Prerequisites

- [Cursor](https://cursor.com/) (Agent chat; project skills under `.cursor/skills/`)
- Go 1.23+ (see `go.mod`; toolchain may fetch 1.24.x)
- [Task](https://taskfile.dev/) (`task` on your `PATH`)
- [Temporal CLI](https://docs.temporal.io/cli) (`temporal` on your `PATH`)

## Phase model

Ticket work lives under `requests/<ticket>/`. Skills under `.cursor/skills/`
match the phase names. Each phase produces a committed artefact a human
approves before the next begins.

| Phase | Folder / skill | Artefact | Depth |
| --- | --- | --- | --- |
| 01-spec | *(human)* | `01-spec/TASK.md` | sketch |
| 02-plan | `/spec-to-plan` | `02-plan/plan.md` | deep |
| 03-implement | `/implement-change` | the diff | sketch |
| 04-validate | `/validate-change` | `04-validation/report.md` | deep |
| 05-review | `/prepare-review` | `05-review/pr.md` | sketch |
| 06-convention | `/create-rule` — only if it recurs | `.cursor/rules/*.mdc` | deep |

Follow [LAB.md](LAB.md) end to end. The 90-minute cohort path is the same loop
with more time on gates; see [MODULE_OUTLINE.md](MODULE_OUTLINE.md).

## Quick environment check

```bash
task start     # local server, Web UI, worker
task repro     # baseline bug: amount -25 completes and moves money
task verify    # go vet ./... && go test ./...
```

Temporal Web: http://localhost:8233

`task repro` should complete Workflow ID `transfer-PAY-1183` with Withdraw and
Deposit Activities — that wrong-success is the starting wound. After a fix,
validation must show the opposite.

## Attribution

Host application forked from
[temporalio/money-transfer-project-template-go](https://github.com/temporalio/money-transfer-project-template-go).
The upstream [LICENSE](LICENSE) is unchanged. Upstream docs and tutorials stay
there; this repository is about the agentic workflow layered on top.
