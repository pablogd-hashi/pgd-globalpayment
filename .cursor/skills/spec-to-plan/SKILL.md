---
name: spec-to-plan
description: >-
  Phase 02 (requests/<ticket>/02-plan). Enter an unfamiliar codebase in this
  repo and produce a reviewable plan before editing code. Use when picking up a
  ticket, when the codebase is unfamiliar, or when the user asks to turn a spec
  into a plan.
---

# Spec to plan

## When to Use

- Picking up a ticket in this repository
- The codebase is unfamiliar
- The user asks to turn a spec into a plan

## Instructions

Produce a reviewable plan. Do not edit application code.

1. **Trace the path** — From the entry point that starts work to every side
   effect that can run. Name the functions and files you read.
2. **Place responsibility** — Say where the change should live, citing
   repository evidence (call sites, types, existing error handling, tests).
3. **Alternatives rejected** — Name the other places this could reasonably go
   and say why each is worse. If you can't argue against an alternative, say
   so.
4. **Preserve** — Note conventions and tests that must keep working.
5. **Propose the smallest change** — Implementation sketch and how you will
   verify it (`task verify`, repro/happy-path tasks, Temporal Web as relevant).
6. Use the ask-questions tool if the spec is ambiguous rather than assuming.

Read the spec from `requests/<ticket>/01-spec/TASK.md`. Write the plan to
`requests/<ticket>/02-plan/plan.md`. Keep it short enough to read on screen in
under a minute.
