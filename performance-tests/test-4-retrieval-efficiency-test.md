
# Test 4: Retrieval Efficiency Test

## Goal

Measure how many files must be opened before locating the implementation.

## Procedure

Perform the same tasks:

```text
Add upload timeout
Add logging
Modify output schema
Add new symbol type
```

Run twice:

### Baseline

Repository Tree Only

### Hermes

Repository Tree + Hermes Map

## Measurements

| Task | Files Opened (Tree Only) | Files Opened (Hermes) |
| ---- | ------------------------ | --------------------- |
|      |                          |                       |
|      |                          |                       |
|      |                          |                       |

## Score

```text
Reduction Percentage
```

## Formula

```text
(TreeFiles - HermesFiles) / TreeFiles * 100
```

## Success Criteria

50%+ reduction.

## Notes

*


