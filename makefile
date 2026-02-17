.PHONY: new_problem new_problem_go new_problem_py add_python \
       test test_go test_py test_problem_go test_problem_py \
       new_ds new_ds_go new_ds_py add_python_ds \
       test_ds test_ds_go test_ds_py test_ds_single_go test_ds_single_py

# ── Problem targets ──────────────────────────────────────────

new_problem:
	@read -p "Enter problem name: " srcname; \
	name=$$(echo $$srcname | sed -e 's/[\. ]/_/g' -e 's/__/_/g'); \
	dir=./problems/$$name; \
	mkdir -p $$dir/go $$dir/python; \
	echo "package problems" > $$dir/go/$$name.go; \
	echo "package problems" > $$dir/go/$${name}_test.go; \
	echo "from solution import Solution" > $$dir/python/test_solution.py; \
	echo "class Solution:" > $$dir/python/solution.py; \
	echo "    pass" >> $$dir/python/solution.py; \
	echo "# $$srcname" > $$dir/README.md; \
	echo "Created: $$dir (go + python)"

new_problem_go:
	@read -p "Enter problem name: " srcname; \
	name=$$(echo $$srcname | sed -e 's/[\. ]/_/g' -e 's/__/_/g'); \
	dir=./problems/$$name; \
	mkdir -p $$dir/go; \
	echo "package problems" > $$dir/go/$$name.go; \
	echo "package problems" > $$dir/go/$${name}_test.go; \
	[ -f $$dir/README.md ] || echo "# $$srcname" > $$dir/README.md; \
	echo "Created: $$dir/go"

new_problem_py:
	@read -p "Enter problem name: " srcname; \
	name=$$(echo $$srcname | sed -e 's/[\. ]/_/g' -e 's/__/_/g'); \
	dir=./problems/$$name; \
	mkdir -p $$dir/python; \
	echo "from solution import Solution" > $$dir/python/test_solution.py; \
	echo "class Solution:" > $$dir/python/solution.py; \
	echo "    pass" >> $$dir/python/solution.py; \
	[ -f $$dir/README.md ] || echo "# $$srcname" > $$dir/README.md; \
	echo "Created: $$dir/python"

add_python:
	@read -p "Enter problem directory name: " name; \
	dir=./problems/$$name/python; \
	mkdir -p $$dir; \
	echo "from solution import Solution" > $$dir/test_solution.py; \
	echo "class Solution:" > $$dir/solution.py; \
	echo "    pass" >> $$dir/solution.py; \
	echo "Created: $$dir"

# ── Test targets ─────────────────────────────────────────────

test: test_go test_py

test_go:
	go test ./...

test_py:
	.venv/bin/python -m pytest -v

test_problem_go:
	@read -p "Enter problem directory name: " name; \
	go test -v ./problems/$$name/go/

test_problem_py:
	@read -p "Enter problem directory name: " name; \
	.venv/bin/python -m pytest -v ./problems/$$name/python/

# ── Data structure targets ───────────────────────────────────

new_ds:
	@read -p "Enter data structure name (snake_case): " name; \
	dir=./data_structures/$$name; \
	mkdir -p $$dir/go $$dir/python; \
	echo "package datastructures" > $$dir/go/$$name.go; \
	echo "package datastructures" > $$dir/go/$${name}_test.go; \
	touch $$dir/python/$$name.py; \
	touch $$dir/python/test_$$name.py; \
	echo "# $$name" > $$dir/README.md; \
	echo "Created: $$dir (go + python)"

new_ds_go:
	@read -p "Enter data structure name (snake_case): " name; \
	dir=./data_structures/$$name; \
	mkdir -p $$dir/go; \
	echo "package datastructures" > $$dir/go/$$name.go; \
	echo "package datastructures" > $$dir/go/$${name}_test.go; \
	[ -f $$dir/README.md ] || echo "# $$name" > $$dir/README.md; \
	echo "Created: $$dir/go"

new_ds_py:
	@read -p "Enter data structure name (snake_case): " name; \
	dir=./data_structures/$$name; \
	mkdir -p $$dir/python; \
	touch $$dir/python/$$name.py; \
	touch $$dir/python/test_$$name.py; \
	[ -f $$dir/README.md ] || echo "# $$name" > $$dir/README.md; \
	echo "Created: $$dir/python"

add_python_ds:
	@read -p "Enter data structure directory name: " name; \
	dir=./data_structures/$$name/python; \
	mkdir -p $$dir; \
	touch $$dir/$$name.py; \
	touch $$dir/test_$$name.py; \
	echo "Created: $$dir"

test_ds: test_ds_go test_ds_py

test_ds_go:
	go test -v ./data_structures/...

test_ds_py:
	.venv/bin/python -m pytest -v ./data_structures/

test_ds_single_go:
	@read -p "Enter data structure directory name: " name; \
	go test -v ./data_structures/$$name/go/

test_ds_single_py:
	@read -p "Enter data structure directory name: " name; \
	.venv/bin/python -m pytest -v ./data_structures/$$name/python/
