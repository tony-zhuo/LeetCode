.PHONY: new_problem test test_problem new_ds test_ds test_ds_single

new_problem:
	@read -p "Enter problem name: " srcname; \
	name=$$(echo $$srcname | sed -e 's/[\. ]/_/g' -e 's/__/_/g'); \
	dir=./problems/$$name; \
	mkdir -p $$dir; \
	echo "package problems" > $$dir/$$name.go; \
	echo "package problems" > $$dir/$${name}_test.go; \
	echo "# $$srcname" > $$dir/README.md; \
	echo "Created: $$dir"

test:
	go test ./...

test_problem:
	@read -p "Enter problem directory name: " name; \
	go test -v ./problems/$$name/

new_ds:
	@read -p "Enter data structure name (snake_case): " name; \
	dir=./data_structures/$$name; \
	mkdir -p $$dir; \
	echo "package datastructures" > $$dir/$$name.go; \
	echo "package datastructures" > $$dir/$${name}_test.go; \
	echo "# $$name" > $$dir/README.md; \
	echo "Created: $$dir"

test_ds:
	go test -v ./data_structures/...

test_ds_single:
	@read -p "Enter data structure directory name: " name; \
	go test -v ./data_structures/$$name/
