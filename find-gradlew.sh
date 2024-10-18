#!/bin/bash

# Find all gradlew files in subfolders and print their parent directories
find_gradlew_parents() {
    # Use the find command to locate all gradlew files
    find . -type f -name "gradlew" | while read -r filepath; do
        # Get the parent directory of each gradlew file
        parent_dir=$(dirname "$filepath")
        echo "$parent_dir"
    done
}

# Call the function
find_gradlew_parents