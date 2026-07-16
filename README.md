# A Common-Sense Guide to Data Structures and Algorithms — Implementation Journal

Working through Jay Wengrow's *A Common-Sense Guide to Data Structures and Algorithms* (2nd Edition) with full implementations in Python and Go.

## Why This Exists

This repo is a deliberate study artifact — not a reference implementation, but a record of working through the material from first principles. Every exercise is attempted cold before any implementation is written. The goal is interview-ready fluency, not a cheat sheet.

## Structure

```
chXX/
  NOTES.md             # raw notes taken while reading
  python/
    examples.py        # in-chapter code examples with assertions
    exercises.py       # end-of-chapter exercises with assertions
  go/
    examples.go        # same examples in Go
    exercises_test.go  # same exercises as Go tests
    go.mod
```

## Running the Code

### Python
```bash
cd chXX/python
python exercises.py
python examples.py
```

### Go
```bash
cd chXX/go
go test ./...
```

## Progress

| Chapter | Topic | Python | Go |
|---------|-------|--------|----|
| 01 | Why Data Structures Matter | ✅ | ✅ |
| 02 | Why Algorithms Matter | ✅ | ✅ |
| 03 | O Yes! Big O Notation | ✅ | ✅ |
| 04 | Speeding Up Your Code With Big O | ✅ | ✅ |
| 05 | Optimizing Code With and Without Big O | ✅ | ✅ |
| 06 | Optimizing for Optimistic Scenarios | ✅ | ✅ |
| 07 | Big O in Everyday Code | 🔄 | 🔄 |
| 08 | | | |
| 09 | | | |
| 10 | | | |
| 11 | | | |
| 12 | | | |
| 13 | | | |
| 14 | | | |
| 15 | | | |
| 16 | | | |
| 17 | | | |
| 18 | | | |
| 19 | | | |
| 20 | | | |

## Languages

- **Python** — reference implementation, closest to Wengrow's pseudocode
- **Go** — forces explicit thinking about types, slices, and memory layout
- **Rust** — planned post-MVP; ownership model will make memory behavior explicit