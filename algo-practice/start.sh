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

# Create the specified files within the new directory and add "package practice" at the top
for file in binary_search.go binary_search_test.go quicksort.go quicksort_test.go counting_sort.go counting_sort_test.go tree.go tree_test.go; 
do
    echo "package practice" > "$next_dir/$file"
done

cd $next_dir
go mod init practice

echo "Created directory $next_dir with the required files."
