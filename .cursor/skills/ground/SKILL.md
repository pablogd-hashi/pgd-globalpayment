---
name: ground
description: >-
  Enter an unfamiliar codebase in this repo and produce a reviewable plan before
  editing code. Use when picking up a ticket, when the codebase is unfamiliar,
  or when the user asks to ground, explore, or plan a change.
---

# Ground

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

Write the plan to `<ticket-id>-plan.md` at the repo root (for example
`PAY-1183-plan.md`). Keep it short enough to read on screen in under a minute.
