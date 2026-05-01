# 🔧 jkit

[![Go](https://github.com/wuhan005/jkit/actions/workflows/go.yml/badge.svg)](https://github.com/wuhan005/jkit/actions/workflows/go.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/wuhan005/jkit.svg)](https://pkg.go.dev/github.com/wuhan005/jkit)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

A small, fast JSON CLI for the things you actually do every day: pretty-print, fold deeply nested trees, pluck out a single value, or build a JSON array from a list of lines. Reads from stdin **or** the system clipboard — copy a JSON blob, run `jkit f`, done.

## Features

| Command | Alias | What it does |
| --- | --- | --- |
| [`jkit format`](#jkit-f--format-json) | `f` | Pretty-print JSON (4-space indent, keeps numeric precision) |
| [`jkit cut <depth>`](#jkit-c-depth--fold-json) | `c` | Collapse nodes deeper than `<depth>` into compact summaries |
| [`jkit get <path>`](#jkit-g-path--extract-by-path) | `g` | Extract a sub-value by dotted path, e.g. `data.items.0.id` |
| [`jkit maker`](#jkit-m--build-a-json-string-array) | `m` | Turn line-separated text into a JSON string array |

## Install

```bash
go install github.com/wuhan005/jkit/cmd/jkit@latest
```

From source:

```bash
git clone https://github.com/wuhan005/jkit.git
cd jkit
go install ./cmd/jkit
```

Requires Go 1.21+. The binary is a single static file with no runtime dependencies.

## Input handling

`jkit` figures out where to read JSON from automatically:

- **Pipe / redirect** — `echo '{"a":1}' | jkit f` or `jkit f < data.json`
- **Clipboard** — when stdin is a terminal, falls back to the system clipboard. Copy a JSON blob anywhere, then just run `jkit f`.

> If stdin has no data and the clipboard is empty, `jkit` exits with `no input from stdin or clipboard`.

## Commands

### `jkit f` — Format JSON

Pretty-print with 4-space indent. Numbers are decoded with `json.Number`, so big integers like `9223372036854775807` survive round-trip without becoming `9.2233720368e+18`.

```bash
> echo '{"data":{"area":"东京Amazon数据中心","country":"日本","ip":"52.68.96.58"},"error":0,"msg":"success"}' | jkit f

{
    "data": {
        "area": "东京Amazon数据中心",
        "country": "日本",
        "ip": "52.68.96.58"
    },
    "error": 0,
    "msg": "success"
}
```

### `jkit c <depth>` — Fold JSON

Collapse anything deeper than `<depth>` levels into `{ N items dict }` / `[ N items array ]` summaries. Object keys are emitted in lexicographic order, so the same input always produces the same output (handy for diffs).

Omit `<depth>` to keep the full tree.

```bash
> jkit c 2

{
    "code": 0,
    "message": "success",
    "result": {
        "main_section": { 4 items dict },
        "section": [ 1 items array ]
    }
}
```

### `jkit g <path>` — Extract by path

Walk a `.`-separated path. Use integer indices for array elements. String leaves print **without quotes** (so you can pipe them straight into other tools); other leaves print as JSON literals.

```bash
> jkit g result.main_section.episodes.0

{
    "badge": "会员",
    "id": 341208,
    "title": "1"
}

> jkit g result.main_section.episodes.0.title
1

> curl -s https://api.example.com/user | jkit g name | xargs -I{} echo "Hello {}"
```

Errors include the path where things broke:

```bash
> echo '{"a":1}' | jkit g a.b
jkit: cannot descend into json.Number at a.b
```

### `jkit m` — Build a JSON string array

Read line-separated text, emit a JSON array. Empty lines are dropped, each line is trimmed, and `\r\n` is handled. Pass `-u` (or `--unique`) to drop duplicates.

```bash
> cat | jkit m
Hi, I'm E99p1ant. 🍆
🐭 Focus on Golang.
🏠 Blog at github.red.

[
    "Hi, I'm E99p1ant. 🍆",
    "🐭 Focus on Golang.",
    "🏠 Blog at github.red."
]
```

```bash
> printf 'apple\nbanana\napple\n' | jkit m -u

[
    "apple",
    "banana"
]
```

## License

[MIT](LICENSE) © wuhan005
