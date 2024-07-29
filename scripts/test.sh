echo "Checking the number of commits..."
commit_count=$(git rev-list --count HEAD)
echo "Commit count: $commit_count"
if [ "$commit_count" -gt 1 ]; then
    echo "Fetching list of new directories..."
    git diff --name-status HEAD~2 | grep '^A' | awk '{print $2}' | xargs -I {} dirname {} | sort | uniq > new_dirs.txt
    echo "Contents of new_dirs.txt:"
    cat new_dirs.txt
    if [ ! -s new_dirs.txt ]; then
      echo "No new directories found."
      exit 0
    fi

    while read -r dir; do
      echo "Processing directory: $dir"
      if [ -d "$dir" ] && [ -f "$dir/go.mod" ]; then
        cd "$dir"
        echo "Running tests in $dir"
        go test -v ./...
        cd - > /dev/null
      else
        echo "Skipping $dir: not a Go module or no go.mod file found."
      fi
    done < new_dirs.txt
    rm new_dirs.txt
else
    echo "No new directories to test."
    exit 0
fi
