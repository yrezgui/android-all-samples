#!/bin/bash

# Find all .kt and .java files recursively in the current directory
# Extract lines that start with "import ", keep only those starting with "android."
# Remove the "import " part, deduplicate, and sort the results

find . -type f \( -name "*.kt" -o -name "*.java" \) -exec grep -hE '^import android\.' {} + | \
    sed 's/^import //' | \
    sort | \
    uniq