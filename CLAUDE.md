# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Structure

This is a multi-language LeetCode problem-solving repository organized **by problem** (not by language). Each problem or data structure directory contains language-specific subdirectories with shared README and assets.

```
problems/{number}_{Problem_Name}/
├── README.md              ← shared
├── assets/                ← shared (if any)
├── go/
│   ├── {number}_{Problem_Name}.go
│   └── {number}_{Problem_Name}_test.go
└── python/
    ├── solution.py
    └── test_solution.py

data_structures/{name}/
├── README.md              ← shared
├── go/
│   ├── {name}.go
│   └── {name}_test.go
└── python/
    ├── {name}.py
    └── test_{name}.py
```

**Supported languages:** Go, Python 3

## Common Commands

### Running Tests
- Run all tests (both languages): `make test`
- Go only: `make test_go` or `go test ./...`
- Python only: `make test_py` (**always use Makefile targets for Python tests**; do NOT run `python3 -m pytest` directly)
- Single problem (Go): `make test_problem_go` (prompts for name)
- Single problem (Python): `make test_problem_py` (prompts for name)

### Creating New Problems
- Both languages: `make new_problem`
- Go only: `make new_problem_go`
- Python only: `make new_problem_py`
- Add Python to existing Go problem: `make add_python`

### Data Structure Commands
- Test all (Go): `make test_ds_go` or `go test -v ./data_structures/...`
- Test all (Python): `make test_ds_py`
- Test single (Go): `make test_ds_single_go`
- Test single (Python): `make test_ds_single_py`
- Scaffold new: `make new_ds` (both), `make new_ds_go`, `make new_ds_py`
- Add Python to existing: `make add_python_ds`

### Module Information
- Go module name: `leet-code` (Go 1.23.0)
- Python: `>=3.10`, pytest for testing (config in `pyproject.toml`)

## Code Architecture

### Go

#### Package Structure
- All problem solutions use `package problems`
- All data structure implementations use `package datastructures`
- Import paths include the `/go` subdirectory:
  - `"leet-code/data_structures/linked_list/go"` for `ListNode` and `Arr2Node()`
  - `"leet-code/data_structures/binary_tree/go"` for `TreeNode` and `Slice2BinaryTree()`
  - `"leet-code/data_structures/graph/go"` for `Node` (graph node with Neighbors)

#### Testing Patterns
- Uses Go's standard testing framework
- Test functions follow `Test_{functionName}` naming convention
- Tests use table-driven test approach with struct slices
- Common test structure includes `name`, `args`, and `want` fields

#### Solution Implementation
- Solutions are implemented as standalone functions (not methods)
- Function names typically match the problem (e.g., `reverseList`)
- Import shared structures aliased as `datastructures`
- Follow Go naming conventions (camelCase for unexported functions)

### Python

#### File Naming
- Solution: `solution.py` with a `class Solution` containing the method
- Tests: `test_solution.py` using pytest

#### Data Structures
- Define data structures (ListNode, TreeNode, etc.) **inline in `solution.py`** — matches LeetCode's style
- Do NOT import from shared data_structures; each solution is self-contained

#### Testing Patterns
- Uses pytest with `class Test*` and `def test_*` conventions
- `conftest.py` at root handles `sys.path` for `from solution import ...`
- Config in `pyproject.toml` with `--import-mode=importlib` to avoid module name collisions

## Multiple Solutions Convention

When recording multiple approaches for the same problem, keep all solutions in the same file and name each with the pattern **`{functionName}_{approach}`**.

### Python

Add multiple methods to the same `Solution` class. Tests share a `CASES` list and parametrize each method separately.

```python
# solution.py
class Solution:
    def twoSum_bruteforce(self, nums, target):
        ...

    def twoSum_hashmap(self, nums, target):
        ...
```

```python
# test_solution.py
import pytest
from solution import Solution

CASES = [
    ([2, 7, 11, 15], 9, [0, 1]),
]

class TestTwoSum:
    def setup_method(self):
        self.sol = Solution()

    @pytest.mark.parametrize("nums,target,want", CASES)
    def test_bruteforce(self, nums, target, want):
        assert self.sol.twoSum_bruteforce(nums, target) == want

    @pytest.mark.parametrize("nums,target,want", CASES)
    def test_hashmap(self, nums, target, want):
        assert self.sol.twoSum_hashmap(nums, target) == want
```

### Go

Add multiple functions to the same `.go` file. Tests share a `cases` variable and have one `Test_` function per approach.

```go
// solution.go
func twoSum_bruteforce(nums []int, target int) []int { ... }
func twoSum_hashmap(nums []int, target int) []int { ... }
```

```go
// solution_test.go
var cases = []struct {
    name   string
    nums   []int
    target int
    want   []int
}{
    {name: "example 1", nums: []int{2, 7, 11, 15}, target: 9, want: []int{0, 1}},
}

func Test_twoSum_bruteforce(t *testing.T) {
    for _, tt := range cases {
        t.Run(tt.name, func(t *testing.T) { ... })
    }
}

func Test_twoSum_hashmap(t *testing.T) {
    for _, tt := range cases {
        t.Run(tt.name, func(t *testing.T) { ... })
    }
}
```

### Common Approach Suffixes

`_bruteforce`, `_hashmap`, `_sorting`, `_twopointer`, `_dp`, `_greedy`, `_bfs`, `_dfs`, `_stack`, `_binarysearch`
