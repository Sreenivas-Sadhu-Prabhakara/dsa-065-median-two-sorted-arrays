# Median of Two Sorted Arrays

> Striver SDE Sheet 065 · Day 11 Binary Search — binary search the cut of the shorter array; the other cut follows, validity is two comparisons. Go, stdlib only, tested.

## Problem

Find the median of the merged order of two sorted arrays in O(log min(n, m)) — the famous hard-tier binary search. Merging is O(n+m); this problem exists to do better.

## Approach

Think in partitions, not elements: cut both arrays so the left pieces together hold exactly half the elements. A cut taking i from the first array forces j = half − i from the second — one degree of freedom, so binary search on i alone (the shorter array keeps j in range). The cut is correct exactly when nothing on a left side exceeds anything on a right side: `aLeft ≤ bRight && bLeft ≤ aRight`. Wrong-way comparisons say which direction to move. ±∞ sentinels for empty sides erase every edge case; the median is then read off the boundary values (max of lefts; for even totals, averaged with min of rights).

**Complexity:** O(log min(n, m)) time, O(1) space.

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
