# Lab — from unfamiliar codebase to verified change

Self-serve path for engineers. Same sequence as the live demo; the 90-minute
workshop stretches the same gates (see [MODULE_OUTLINE.md](MODULE_OUTLINE.md)).
Facilitator cue sheet for a 15-minute excerpt: [RUNBOOK.md](RUNBOOK.md).

**Subject:** agentic SDLC — spec → plan → implement → validate → review →
persist → reuse, with a human approval between phases.

**Not the subject:** learning Temporal. You only need to open Temporal Web to
see whether Withdraw/Deposit Activities ran. That history is evidence, not a
tutorial.

## Prerequisites

- Cursor with Agent chat
- Go 1.23+ (`go.mod`)
- Task (`task` on `PATH`)
- Temporal CLI (`temporal` on `PATH`)

```bash
cd /path/to/pgd-globalpayment
task demo-reset
```

Expect: Temporal comes up, workflow count starts at 0, then a seeded run
`transfer-PAY-1183` completes for amount `-25` with Withdraw and Deposit.
Open http://localhost:8233 and confirm that wrong-success before you continue.

If you are mid-session and only need the stack: `task start`. To prove the
wound again without a full reset: `task repro`.

`demo-reset` switches to `lab/base`, restores the broken app and seeded rules,
cleans session artefacts under `requests/` (keeps `01-spec/TASK.md`), and
fails if lab-relevant paths are dirty. Commit lab infra onto `lab/base` first;
keep any PAY-1183 fix and phase-06 guard rule off that branch.

## What you will produce

| Phase | Skill / action | Artefact | Your gate |
| --- | --- | --- | --- |
| 01-spec | human (or use the file on disk) | `01-spec/TASK.md` | Is “done” checkable? |
| 02-plan | `/spec-to-plan` | `02-plan/plan.md` | Approve placement before code |
| 03-implement | `/implement-change` | the diff | Diff matches the plan only |
| 04-validate | `/validate-change` | `04-validation/report.md` | Evidence vs each criterion |
| 05-review | `/prepare-review` | `05-review/pr.md` | Reviewer can judge without redoing the work |
| 06-convention | `/create-rule` if it recurs | `.cursor/rules/*.mdc` | Next session can inherit the decision |

A convention is worth extracting only when the decision would recur for a
different ticket and someone could reasonably get it wrong. That judgement is
yours; `/create-rule` does the writing.

## Steps

Work in one Cursor Agent chat until the reuse step. Paste the prompts as
written. Stop at each gate before the next skill.

### 0 — See the wound

Read `slack_thread.md`. In Temporal Web, open Workflow ID `transfer-PAY-1183`
(Completed, amount `-25`, Withdraw + Deposit). Completion here is not proof
that forbidden side effects stayed off.

Optional contrast (ungrounded explore — stop it early; do not let it finish a
full freestyle investigation):

```
Transfers with amount <= 0 are completing and money is moving
(Withdraw + Deposit). Workflow ID transfer-PAY-1183.

Where in this codebase is that allowed, and where should it be rejected
so no Withdraw or Deposit Activity runs?
```

That cost is why the lab runs phases instead of improvising.

### 1 — Spec

Self-serve default: use `requests/PAY-1183/01-spec/TASK.md` already on disk.

To draft it yourself:

```
Turn slack_thread.md into requests/PAY-1183/01-spec/TASK.md.

Requirement and acceptance criteria only. Do not explore the codebase
or edit application code.
```

**Gate:** acceptance criteria are checkable before any repository exploration.

### 2 — Plan

```
/spec-to-plan

Ticket PAY-1183. Spec is at requests/PAY-1183/01-spec/TASK.md.
Write the plan to requests/PAY-1183/02-plan/plan.md.
```

**Gate:** read the plan. Approve or correct placement and rejected alternatives
before any code change. Do not skip this gate.

### 3 — Implement

```bash
git switch -c pay-1183
```

```
/implement-change

Ticket PAY-1183. Implement only the approved plan at
requests/PAY-1183/02-plan/plan.md.
```

Then reload the worker so Temporal runs the new code:

```bash
task reload-worker
```

**Gate:** the diff matches the approved plan and nothing more.

### 4 — Validate

```
/validate-change

Ticket PAY-1183.
```

Confirm against the new Temporal run (failed / rejected — not the seeded
COMPLETED wrong-success). Map each acceptance criterion to evidence
(`task verify`, `task repro`, `task transfer`, history).

**Gate:** accept only if side effects never started for the invalid amount and
the ordinary path still works. Green unit tests with a still-COMPLETED Temporal
run usually mean a stale worker — `task reload-worker` and re-validate.

### 5 — Review

```
/prepare-review

Ticket PAY-1183.
```

Artefact: `requests/PAY-1183/05-review/pr.md`. Do not open a GitHub PR unless
you choose to outside this lab.

### 6 — Persist the decision (if it recurs)

```
/create-rule

From requests/PAY-1183/02-plan/plan.md, extract only the recurring placement
decision and rejected alternatives. Write a glob-scoped rule under
.cursor/rules, attaching on workflow.go and activity.go.

State why, not just what: an Activity cannot prevent its own scheduling; a
starter-only check is bypassable by other clients. Do not name the fix — no
amount, error type, threshold, ticket id, or line reference. Do not commit
the rule onto lab/base.
```

### 7 — Reuse (new chat)

Open a **new** Agent chat. Do not re-ask about negative amounts (the fix is
already in the code). Ask a near-transfer question:

```
Transfers where the source and target are the same account should also be
rejected before money moves. Where does that check belong?
```

You are testing whether the session retrieves and applies the committed rule
without re-running `/spec-to-plan`. If it does not, treat that as a failed
inheritance gate, not a silent success.

## Map to the 90-minute outline

| Workshop block ([MODULE_OUTLINE.md](MODULE_OUTLINE.md) §6) | Lab step |
| --- | --- |
| Spec from the report | §1 |
| `/spec-to-plan` + development-entry gate | §2 |
| `/implement-change` | §3 |
| `/validate-change` + acceptance gate | §4 |
| `/prepare-review` | §5 |
| Persist the decision | §6 |
| Near-transfer assessment | §7 |

The live 15-minute proof pre-stages plan/implement and spends the clock on
gates and inheritance; you can run every skill live here.

## Reset and troubleshooting

```bash
task demo-reset
```

| Problem | Fix |
| --- | --- |
| `demo-reset` reports dirty lab paths | Commit or discard WIP on `lab/base` (no Go fix, no guard rule on that branch). |
| Extra rules under `.cursor/rules/` | Reset; remove extras from `lab/base` tip if committed there. |
| Temporal down | `task start` or `task restart` |
| Stale worker (tests green, Temporal still COMPLETED) | `task reload-worker`, then re-run validate |
| Need ordinary happy path | `task transfer` |
