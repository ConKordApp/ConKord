# ConKord ecosystem contribution guidance

## Scope

- Keep changes focused on the requested task.
- Preserve existing behavior unless the task explicitly changes it.
- Do not modify generated files or dependencies directly.

## Development

- Follow the conventions already used in the affected package.
- Keep React components and Go packages small, typed, and easy to test.
- Handle errors explicitly and avoid committing secrets or local configuration.
- Update relevant documentation when behavior or developer workflows change.

## Validation

- Add or update tests whenever behavior changes.
- Run the applicable formatter, linter, tests, and build before finishing.
- Report the commands run and any checks that could not be completed.

## Git

- Keep commits small and logically scoped.
- Do not rewrite shared history or interact with remotes unless explicitly asked.
- Never commit build outputs, dependency directories, credentials, or editor state.
