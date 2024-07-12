#!/bin/bash

# Find the highest existing "day-XXX" directory
last_dir=$(ls -d day-* 2>/dev/null | sort -V | tail -n 1)

# Extract the number from the last directory and increment it
if [[ -z "$last_dir" ]]; then
    next_num=1
else
    last_num=${last_dir#day-}
    next_num=$((10#$last_num + 1))
fi

# Format the new directory name
next_dir=$(printf "day-%03d" "$next_num")

# Create the new directory
mkdir "$next_dir"

# Create the specified files within the new directory
touch "$next_dir/binary_search.go" "$next_dir/binary_search_test.go" "$next_dir/quicksort.go" "$next_dir/quicksort_test.go" "$next_dir/counting_sort.go" "$next_dir/counting_sort_test.go"

echo "Created directory $next_dir with the required files."
