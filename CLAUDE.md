# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository. Keep this file up to date when making changes to the project.

## Commands

```bash
# Build
go build -v ./...

# Install locally
go install .

# Run tests
go test -v ./...

# Run a single test
go test -v ./internal/app/... -run TestFunctionName
```

## Architecture

NHL scores CLI built with [Cobra](https://github.com/spf13/cobra) for commands and [go-pretty](https://github.com/jedib0t/go-pretty) for table rendering.

**Layers:**
- `cmd/` — Cobra command handlers (`today`, `yesterday`, `on <date>`). Root command defines `--team` and `--short` flags.
- `internal/api/` — HTTP requests to `https://api-web.nhle.com/v1/score/{date}`, unmarshals JSON into `app.DailyScores`.
- `internal/app/` — Core logic: data models, table rendering, team validation, date validation.

**Data flow:**
```
cmd handler → api.GetScoresFor*() → DailyScores.FilterByTeam() → DailyScores.String() → toTable() → stdout
```

**Table rendering** is in `internal/app/models.go`. The `DailyScores.String()` method calls `toTable()`, which iterates `Game` objects and calls `game.toRow()` — each column is a dedicated method (`gameTypeCol`, `statusCol`, `gwgCol`, etc.). `SuppressEmptyColumns()` hides playoff columns for regular season games. When `DailyScores.Short` is true (`--short`/`-s` flag), `toTable()` renders a compact 3-column table (Teams, Score, Status) instead of the full 8 columns.

**Key structures in `internal/app/models.go`:**
- `DailyScores` — top-level response, holds `[]Game` and `Short` bool for compact output
- `Game` — single game with nested `HomeTeam`, `AwayTeam`, `PeriodDescriptor`, `GameWinner` structs
- `game.toRow()` — returns a `table.Row` with 8 columns; this is the primary place to modify display logic

**Validation:**
- Dates validated in `internal/app/datetime.go` (YYYY-MM-DD, not future, on/after 1917-12-19)
- Teams validated in `internal/app/teams.go` (32 NHL team abbreviations, case-insensitive)
