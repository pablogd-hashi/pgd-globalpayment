---
name: prepare-review
description: >-
  Phase 05 (requests/<ticket>/05-review). Draft a reviewer-facing summary for a
  ticket from the spec and validation report. Use after validation, when
  preparing review material, or when the user asks for a PR write-up without
  creating a pull request.
---

# Prepare review

## When to Use

- After validation
- When preparing review material
- When the user asks for a PR write-up without creating a pull request

## Instructions

Do not change application code. 

Read `requests/<ticket>/01-spec/TASK.md` and
`requests/<ticket>/04-validation/report.md`. Write
`requests/<ticket>/05-review/pr.md` containing:

1. A title for the change.
2. A one-paragraph summary of the change.
3. A table mapping each acceptance criterion from the spec to the evidence
   from the validation report.

The artefact is what a reviewer reads.
