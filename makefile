.PHONY: new_problem
new_problem:
	@read -p "Enter problem name: " name; \
	name=$$(echo $$name | sed -e 's/[\. ]/_/g' -e 's/__/_/g'); \
    mkdir -p ./problems/$$name; \
    echo "Created directory: $$name"