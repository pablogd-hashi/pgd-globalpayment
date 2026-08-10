---
name: validate-change
description: >-
  Phase 04 (requests/<ticket>/04-validation). Prove a change is safe in this
  side-effecting Temporal codebase. Use after implementing a fix, before
  treating work as done, or when the user asks to validate, check, or confirm
  acceptance criteria from the spec.
paths:
  - workflow.go
  - "**/*_test.go"
---

# Validate change

## When to Use

- After implementing a fix
- Before treating work as done
- When the user asks to validate, check, or confirm acceptance criteria from
  the spec

## Instructions

Do not change application code while verifying.

1. Run `task verify`. Identify which tests cover the invalid input and state
   what each one actually proves. A test that only asserts an error was
   returned does not prove the side effect never started — say so plainly if
   that is what you find.
2. Run `task repro`. Report the Temporal event count and Activity count for
   that run.
3. Run `task transfer`. Confirm the ordinary path still completes.
4. For each acceptance criterion in `requests/<ticket>/01-spec/TASK.md`, state
   whether it is met and cite the evidence (test output, Temporal history,
   command results).

Write the report to `requests/<ticket>/04-validation/report.md`.
