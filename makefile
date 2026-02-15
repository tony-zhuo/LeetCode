.PHONY: new_problem test test_problem

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
