# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Structure

This is a Go-based LeetCode problem-solving repository with the following structure:

- `problems/` - Contains individual LeetCode problem solutions, each in its own directory
- `structure/` - Shared data structures and utilities (ListNode, TreeNode, Graph)
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

### Module Information
- Module name: `leet-code`
- Go version: 1.23.0

## Code Architecture

### Package Structure
- All problem solutions use `package problems`
- Shared data structures are in `package structure`
- Common imports: `"leet-code/structure"` for data structures

### Data Structures
The `structure/` package provides:
- `ListNode` - Singly linked list with `Arr2Node()` helper function
- `TreeNode` - Binary tree with `Slice2BinaryTree()` helper function  
- `graph.go` - Graph-related structures

### Testing Patterns
- Uses Go's standard testing framework
- Test functions follow `Test_{functionName}` naming convention
- Tests use table-driven test approach with struct slices
- Common test structure includes `name`, `args`, and `want` fields

### Solution Implementation
- Solutions are implemented as standalone functions (not methods)
- Function names typically match the problem (e.g., `reverseList`)
- Import shared structures from `leet-code/structure` package
- Follow Go naming conventions (camelCase for unexported functions)