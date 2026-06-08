#!/usr/bin/env bash

repos=(
  tree-sitter-go
  tree-sitter-python
  tree-sitter-rust
  tree-sitter-javascript
  tree-sitter-typescript
)

for repo in "${repos[@]}"; do
  git clone "https://github.com/tree-sitter/${repo}.git"
done