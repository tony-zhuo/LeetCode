.PHONY: new_problem
new_problem:
	@read -p "Enter problem name: " srcname; \
	name=$$(echo $$srcname | sed -e 's/[\. ]/_/g' -e 's/__/_/g'); \
    mkdir -p ./problems/$$name; \
    cd ./problems/$$name; \
    touch $$name.go; \
	touch $$name"_test.go"; \
	echo "# $$srcname" > README.md; \
    echo "Created directory: $$name and file: $$name/$$name.go"