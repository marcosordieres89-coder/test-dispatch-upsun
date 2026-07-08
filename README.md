# test-dispatch-upsun

A deliberately tiny, dependency-free Go project used to exercise
[Dispatch](https://linear.app/platformsh) issue-to-PR workflows against a real
GitHub repository.

`stringkit` is a small package of string helpers with table-driven tests and a
thin CLI. It's intentionally simple: each feature request maps cleanly onto one
change plus one test, so it's an ideal target for automated planning →
implementation → PR flows.

## Library

```go
import stringkit "github.com/marcosordieres89-coder/test-dispatch-upsun"

stringkit.Reverse("hello")        // "olleh"
stringkit.IsPalindrome("RaceCar") // true
stringkit.WordCount("a b  c")     // 3
```

## CLI

```bash
go run ./cmd/stringkit reverse "hello"        # olleh
go run ./cmd/stringkit palindrome "racecar"   # true
go run ./cmd/stringkit words "hello world"    # 2
```

## Develop

```bash
go vet ./...
go test -race ./...
```

CI (`.github/workflows/ci.yml`) runs `go vet` and `go test -race` on every push
to `main` and on every pull request.

## Roadmap (good first issues to throw at Dispatch)

Each of these is a self-contained feature — file it as a GitHub issue and let
the workflow open the PR:

- `Truncate(s string, n int) string` — cap a string to `n` runes, appending `…`
  when it was shortened.
- `Capitalize(s string) string` — upper-case the first letter of each word.
- Add a `--json` output mode to the CLI.
- `CountVowels(s string) int`.
- Make the CLI read from stdin when no string argument is given.
