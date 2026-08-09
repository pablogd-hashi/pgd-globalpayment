# PAY-1183

## Context

You've picked up PAY-1183 in a payments codebase you don't work in daily.
You didn't write this code. The codebase runs Temporal Workflows that move
money between accounts. You have this repository, `task repro`, `task verify`,
`task transfer`, and Temporal Web on http://localhost:8233.

## Requirement

Transfers with an amount less than or equal to zero must be rejected before
money movement begins.

## Acceptance criteria

- No withdrawal or deposit Activity runs for a rejected transfer.
- Valid transfers continue to work.
- Existing tests keep passing.
