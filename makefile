.PHONY: new_problem
new_problem:
	@read -p "Enter problem name: " srcname; \
	name=$$(echo $$srcname | sed -e 's/[\. ]/_/g' -e 's/__/_/g'); \
    mkdir -p ./problems/$$name; \
    cd ./problems/$$name; \
    read -p "need assets folder? (y/n): " isNeedAssetsFolder; \
    if [ "$$isNeedAssetsFolder" = "y" ]; then \
    	mkdir -p ./assets; \
    fi; \
    echo "package problems" > $$name.go; \
    echo "package problems" > $$name"_test.go"; \
	echo "# $$srcname" > README.md; \
    echo "Created directory: $$name and file: $$name/$$name.go"