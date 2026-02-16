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
- Python only: `make test_py` or `python3 -m pytest -v`
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
