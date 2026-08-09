# Lab path

This session leaves the codebase carrying context it did not have before.
Skills (`/ground`, `/verify`) are org infrastructure: already in this repo,
written once and reused across repositories. Starter rules may already exist;
the Persist step adds the convention this task discovers. Rules are
per-codebase and accumulate as architectural decisions get made.

1. Read `slack_thread.md` — the incomplete report that starts the work.
2. Turn the thread into `TASK.md` — requirement and acceptance criteria.
   (Self-serve: use the provided `TASK.md` if you are not drafting it live.)
3. Run `/ground` — explore the codebase and produce a reviewable plan.
4. Review the plan — approve or correct before any code change.
5. Implement the smallest change that meets the requirement.
6. Run `/verify` — prove the change is safe against tests and Temporal history.
7. Persist — add a glob-scoped `.cursor/rules/*.mdc` rule for the boundary you
   discovered, so the next engineer does not rediscover it.
