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

cp -rf _template/* $next_dir

cd $next_dir
go get golang.org/x/exp/constraints
go get gotest.tools/assert
go mod tidy

echo "Created directory $next_dir with the required files."
