#!/bin/bash

for i in $(seq -w 1 20); do
  dir="ch$i"
  mkdir -p "$dir/python"
  mkdir -p "$dir/go"

  # NOTES.md
  cat > "$dir/NOTES.md" << EOF
# Chapter $i Notes

## Key Concepts

## Questions

## Observations
EOF

  # Python stubs
  cat > "$dir/python/examples.py" << EOF
"""
Chapter $i: Examples
"""

EOF

  cat > "$dir/python/exercises.py" << EOF
"""
Chapter $i: Exercises
"""

if __name__ == "__main__":
    print("exercises.py: no exercises yet")
EOF

  # Go stubs
  cat > "$dir/go/go.mod" << EOF
module github.com/djviodes/wengrow-dsa/$dir

go 1.21
EOF

  cat > "$dir/go/examples.go" << EOF
package ch$i

// TODO: implement as you read
EOF

  cat > "$dir/go/exercises_test.go" << EOF
package ch$i

import "testing"

func TestPlaceholder(t *testing.T) {
	// remove when first real test is added
}
EOF

done

echo "All 20 chapters scaffolded."