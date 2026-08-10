# Architecture

## Problem

Teams adopting agents in an unfamiliar codebase cannot verify output cheaply.
When a change is hard to check, engineers revert to doing the work themselves.
The cost of rebuilding context every session compounds that hesitation.

## Approach

Work proceeds in specification-driven phases with human gates between them.

- Skills under `.cursor/skills/` are deterministic procedures for each phase.
- Rules under `.cursor/rules/` encode conventions that should survive a session.
- Each phase produces a committed artefact. A human approves it before the next
  phase begins.

The agent explores, proposes, and implements within an approved boundary. The
human still owns what “done” means, whether placement is correct, and whether
the evidence is enough to accept the change.

## Phase model

Ticket work lives under `requests/<ticket>/`. Skills under `.cursor/skills/`
match the phase names. This lab goes deep on plan, validation, and (when the
decision would recur) convention persistence. Spec, implement, and review are
present but thinner.

| Phase | Folder / skill | Artefact | Depth |
| --- | --- | --- | --- |
| 01-spec | *(human)* | `01-spec/TASK.md` | sketch |
| 02-plan | `/spec-to-plan` | `02-plan/plan.md` | deep |
| 03-implement | `/implement-change` | the diff | sketch |
| 04-validate | `/validate-change` | `04-validation/report.md` | deep |
| 05-review | `/prepare-review` | `05-review/pr.md` | sketch |
| 06-convention | `/create-rule` — only if it recurs | `.cursor/rules/*.mdc` | deep |

## Artefact table

| Artefact | Reviewer question |
| --- | --- |
| `01-spec/TASK.md` | What does done mean, in checkable terms? |
| `02-plan/plan.md` | Where should the change live, and why not elsewhere? |
| the diff | Does the change match the approved plan and nothing more? |
| `04-validation/report.md` | Is there evidence each acceptance criterion is met? |
| `05-review/pr.md` | Can a reviewer judge the change without redoing the investigation? |
| `.cursor/rules/*.mdc` | Will the next session inherit this decision without re-deriving it? |

## Why this repo

The host application is a real public Temporal money-transfer template with a
genuine gap: transfers with amount ≤ 0 complete and move money. The codebase is
small enough to understand in about twenty minutes, so learners spend time on
placement and evidence rather than navigation.

Temporal is the runtime that makes side effects observable. It is not the
subject of the lab.
