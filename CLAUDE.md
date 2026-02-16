# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Structure

This is a Go-based LeetCode problem-solving repository with the following structure:

- `problems/` - Contains individual LeetCode problem solutions, each in its own directory
- `data_structures/` - From-scratch implementations of common data structures, organized by category
- `libs/` - Additional helper libraries (if any)

Each problem directory follows this naming pattern: `{number}_{Problem_Name}/` and contains:
- `{number}_{Problem_Name}.go` - Main solution implementation
- `{number}_{Problem_Name}_test.go` - Test cases
- `README.md` - Problem description

## Common Commands

### Running Tests
- Run all tests: `go test ./...`
- Run tests for a specific problem: `cd problems/{problem_name} && go test -v`
- Run with verbose output: `go test -v ./...`

### Creating New Problems
Use the makefile to scaffold new problem directories:
```bash
make new_problem
```
This will prompt for the problem name and create the appropriate directory structure with boilerplate Go files.

### Data Structure Tests
- Run all data structure tests: `make test_ds` or `go test -v ./data_structures/...`
- Run specific: `make test_ds_single` or `go test -v ./data_structures/{name}/`
- Scaffold new: `make new_ds`

### Module Information
- Module name: `leet-code`
- Go version: 1.23.0

## Code Architecture

### Package Structure
- All problem solutions use `package problems`
- All data structure implementations use `package datastructures`
- Common imports for problem solutions:
  - `"leet-code/data_structures/linked_list"` for `ListNode` and `Arr2Node()`
  - `"leet-code/data_structures/binary_tree"` for `TreeNode` and `Slice2BinaryTree()`
  - `"leet-code/data_structures/graph"` for `Node` (graph node with Neighbors)

### Testing Patterns
- Uses Go's standard testing framework
- Test functions follow `Test_{functionName}` naming convention
- Tests use table-driven test approach with struct slices
- Common test structure includes `name`, `args`, and `want` fields

### Solution Implementation
- Solutions are implemented as standalone functions (not methods)
- Function names typically match the problem (e.g., `reverseList`)
- Import shared structures from `leet-code/data_structures/` subpackages (aliased as `datastructures`)
- Follow Go naming conventions (camelCase for unexported functions)