#!/bin/bash
set -e
shopt -s nullglob
shopt -s globstar

TEMP_DIR="$(dirname -- "$(readlink -f -- "$BASH_SOURCE")")/.."
PROJECT_DIR="$TEMP_DIR/.."
SCRIPT_DIR="$PROJECT_DIR/scripts/ms-scaffolder"

# Load env.
if [ ! -f "$PROJECT_DIR/.env" ]; then
    cp $PROJECT_DIR/.env.template $PROJECT_DIR/.env
fi
source $PROJECT_DIR/.env

# Ensure git is clean before proceeding.
if [ ! -z "$(git status --porcelain)" ]; then
    echo "Git must be clean before running this command. Clean up your changes and try again."
    exit 1
fi

# Copy files over that should be fully managed by ms-scaffolder.
rm -r $PROJECT_DIR/scripts/ms-scaffolder/*
cp $TEMP_DIR/scripts/ms-scaffolder/* $PROJECT_DIR/scripts/ms-scaffolder/
cp $TEMP_DIR/templates/$MS_SCAFFOLDER_TEMPLATE/scripts/ms-scaffolder/* $PROJECT_DIR/scripts/ms-scaffolder/
cp $TEMP_DIR/Makefile $PROJECT_DIR/
cp $TEMP_DIR/README-HelmChart.md $PROJECT_DIR/
cp $TEMP_DIR/README-Makefile.md $PROJECT_DIR/

# Replace variables in template files and remove the template files.
for FILE in $SCRIPT_DIR/**/*.template $SCRIPT_DIR/**/.*.template; do
    if [ -z "$PROJECT_SLUG" ] || [ -z "$BUILD_IMAGE" ]; then
        echo "Missing template variable."
        exit 1
    fi

    NEW_FILE=${FILE%.*}
    
    # Copy the file first to preserve permissions.
    cp $FILE $NEW_FILE

    sed -e "s~{{PROJECT_SLUG}}~$PROJECT_SLUG~g" \
        -e "s~{{PROJECT_SLUG_NO_HYPHEN}}~$PROJECT_SLUG_NO_HYPHEN~g" \
        -e "s~{{BUILD_IMAGE}}~$BUILD_IMAGE~g" \
        $FILE > $NEW_FILE

    rm $FILE
done

# Apply do not edit text to appropriate files.
for FILE in $PROJECT_DIR/scripts/ms-scaffolder/*; do
    echo -e "$(cat $TEMP_DIR/donotedit.txt)\n\n$(cat $FILE)" > $FILE
done
echo -e "$(cat $TEMP_DIR/donotedit.txt)\n\n$(cat $PROJECT_DIR/Makefile)" > $PROJECT_DIR/Makefile

# Commit new versions of scripts and files.
git add -A .
git commit -m "Update ms-scaffolder scripts and files."

echo "ms-scaffolder update complete and committed to current working branch. If you'd like to push, please run 'git push'."

# Remove temporary directory
rm -r $TEMP_DIR

