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

## Installing TypeScript 7

This Go project doesn't depend on TypeScript, but if you want to try the
TypeScript 7 toolchain alongside it, here's how to install it.

**Prerequisites:** a current LTS release of [Node.js](https://nodejs.org/)
(which bundles `npm`). Check what you have with:

```bash
node --version
npm --version
```

### Global install

Install the compiler system-wide so `tsc` is on your `PATH`:

```bash
npm install -g typescript@7
```

### Project (dev-dependency) install

Add TypeScript 7 to a single project instead, so the version is pinned in
`package.json`:

```bash
npm install --save-dev typescript@7
```

Invoke the project-local compiler with `npx`:

```bash
npx tsc --init   # create a tsconfig.json
npx tsc          # compile the project
```

### Alternative package managers

```bash
yarn add --dev typescript@7   # Yarn
pnpm add -D typescript@7      # pnpm
```

### Verify the installation

```bash
tsc --version        # global install
npx tsc --version    # project-local install
```

The command should print a `7.x` version string.

> **Note:** "TypeScript 7" refers to the native (Go-based) compiler. While it is
> still in preview, Microsoft publishes it under the
> [`@typescript/native-preview`](https://www.npmjs.com/package/@typescript/native-preview)
> package. If the `typescript@7` tag isn't available for your setup yet, install
> the preview instead (its binary is `tsgo`):
>
> ```bash
> npm install -g @typescript/native-preview
> tsgo --version
> ```

## Roadmap (good first issues to throw at Dispatch)

Each of these is a self-contained feature — file it as a GitHub issue and let
the workflow open the PR:

- `Truncate(s string, n int) string` — cap a string to `n` runes, appending `…`
  when it was shortened.
- `Capitalize(s string) string` — upper-case the first letter of each word.
- Add a `--json` output mode to the CLI.
- `CountVowels(s string) int`.
- Make the CLI read from stdin when no string argument is given.
