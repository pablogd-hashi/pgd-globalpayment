# Lab path

This session leaves the codebase carrying context it did not have before.
Skills under `.cursor/skills/` are numbered to match `requests/<ticket>/`
phases. Starter rules may already exist; phase 06 adds the convention this
task discovers — only when it would recur. Rules are per-codebase and
accumulate as architectural decisions get made.

Each phase produces a committed artefact a human approves before the next
begins. Reviewers can scope to a single phase without reading unrelated ones.

| Phase | Folder / skill | Artefact | Depth |
| --- | --- | --- | --- |
| 01-spec | *(human)* | `01-spec/TASK.md` | sketch |
| 02-plan | `/02-spec-to-plan` | `02-plan/plan.md` | deep |
| 03-implement | `/03-implement-change` | the diff | sketch |
| 04-validate | `/04-validate-change` | `04-validation/report.md` | deep |
| 05-review | `/05-prepare-review` | `05-review/pr.md` | sketch |
| 06-convention | `/create-rule` — only if it recurs | `.cursor/rules/*.mdc` | deep |

A convention is worth extracting when the decision would recur for a
different ticket in this codebase and someone could reasonably get it wrong.
If the fix was specific to this ticket, write nothing. The judgement is the
human's; `/create-rule` does the writing.

1. Read `slack_thread.md` — the incomplete report that starts the work.
2. Turn the thread into `requests/<ticket>/01-spec/TASK.md` — requirement and
   acceptance criteria. (Self-serve: use the provided
   `requests/PAY-1183/01-spec/TASK.md` if you are not drafting it live.)
3. Run `/02-spec-to-plan` — explore the codebase and produce a reviewable plan
   at `requests/<ticket>/02-plan/plan.md`.
4. Review the plan — approve or correct before any code change.
5. Run `/03-implement-change` — implement only the approved plan.
6. Run `/04-validate-change` — prove the change is safe; write
   `requests/<ticket>/04-validation/report.md`.
7. Run `/05-prepare-review` — draft `requests/<ticket>/05-review/pr.md` for
   the reviewer.
8. Phase 06 — if the decision would recur, use `/create-rule` to persist a
   glob-scoped `.cursor/rules/*.mdc` rule for the boundary you discovered.

To re-run the lab from the broken baseline (wrong-success Temporal history),
run `task demo-reset`. That switches to `lab/base`, restores the three app
files and `.cursor/rules/` from the branch tip, runs
`git clean -fdx -- requests/ .cursor/rules/` (ignored session dirs included),
then **fails** if `git status --porcelain` is non-empty or if
`.cursor/rules/` is not exactly the two seeded rules (`change-scope.mdc`,
`go-conventions.mdc`). Commit lab infra onto `lab/base` first; keep the
PAY-1183 fix and any phase-06 guard rule off that branch.
