# Hermes Benchmark Suite

## Repository Information

| Field       | Value |
| ----------- | ----- |
| Repository  |       |
| Date        |       |
| Commit      |       |
| Branch      |       |
| Language(s) |       |
| Total LOC   |       |
| Total Files |       |

---

# Test 1: Repository Reduction Test

## Goal

Measure how much Hermes compresses repository structure.

## Inputs

```bash
tree
cloc .
./hermes -input .
```

## Measurements

| Metric              | Raw Repository | Hermes |
| ------------------- | -------------- | ------ |
| Files               |                |        |
| LOC                 |                |        |
| Output Lines        | N/A            |        |
| Output Size (Bytes) | N/A            |        |
| Estimated Tokens    | N/A            |        |

## Formula

Estimated Tokens:

```text
bytes / 4
```

## Success Criteria

* Hermes output significantly smaller than repository LOC.
* Symbol information preserved.

## Notes

*

---


