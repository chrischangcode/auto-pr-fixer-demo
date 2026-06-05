# auto-pr-fixer-demo

A sample Go project demonstrating [auto-pr-fixer](https://github.com/chrischangcode/auto-pr-fixer) — a GitHub Action that uses Copilot CLI to automatically fix PR CI failures.

## Scenario

This repo simulates the scenario from [Azure/kubelogin#753](https://github.com/Azure/kubelogin/pull/753):

1. **Main branch** has a working Go project with CI (lint + test + build)
2. **A PR** bumps the Go version and linter version for CVE remediation
3. The new linter version catches issues that the old version didn't — **CI fails**
4. **auto-pr-fixer** detects the failure, fetches the logs, and posts a fix suggestion

## Project structure

```
├── main.go                     # CLI entry point
├── pkg/
│   ├── token/token.go          # OIDC token fetcher (has a gosec-flaggable URL usage)
│   └── version/version.go      # Build version info (uses magic strings)
├── .github/
│   ├── auto-pr-fixer.yml       # auto-pr-fixer config
│   └── workflows/
│       ├── ci.yml              # CI: lint + test + build
│       └── auto-pr-fixer.yml   # Triggers on CI failure
└── go.mod
```

## How to trigger the demo

The PR on branch `fix/cve-remediation-go-bump` bumps the Go version and linter,
which introduces new lint findings that cause CI to fail. auto-pr-fixer then
kicks in and suggests fixes.
