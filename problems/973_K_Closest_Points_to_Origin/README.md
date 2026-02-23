# 973. K Closest Points to Origin

## Problem

Given an array of `points` where `points[i] = [xi, yi]` represents a point on the X-Y plane and an integer `k`, return the `k` closest points to the origin `(0, 0)`.

The distance between two points on the X-Y plane is the Euclidean distance (`sqrt(x1^2 + y1^2)`).

You may return the answer in any order.

## Solutions

### Go

#### Sorting (`kClosest`)

Sort all points by their squared Euclidean distance to the origin, then return the first `k` points. No need to compute the actual square root since we only compare relative distances.

- Time: O(n log n)
- Space: O(1) (in-place sort, excluding output)