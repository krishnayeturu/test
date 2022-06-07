#!/bin/bash
set -e
shopt -s nullglob
shopt -s globstar
source .env

SCRIPTS_PATH=./scripts/ms-scaffolder
TEMP_DIR=./.ms-scaffolder

# Ensure git is clean before proceeding.
if [ ! -z "$(git status --porcelain)" ]; then
    echo "Git must be clean before running this command. Clean up your changes and try again."
    exit 1
fi

# Temporarily remove .gitignore for ms-scaffolder scripts.
rm $SCRIPTS_PATH/.gitignore
git add -A .
git commit -m "Temporarily remove $SCRIPTS_PATH/.gitignore for update purposes."

# Copy files over that should be fully managed by ms-scaffolder.
rm -r ./scripts/ms-scaffolder/*
cp $TEMP_DIR/scripts/ms-scaffolder/* ./scripts/ms-scaffolder/
cp $TEMP_DIR/templates/$MS_SCAFFOLDER_TEMPLATE/scripts/ms-scaffolder/* ./scripts/ms-scaffolder/
cp $TEMP_DIR/Makefile ./
cp $TEMP_DIR/README-Makefile.md ./

# Replace variables in template files and remove the template files.
for FILE in $SCRIPTS_PATH/**/*.template $SCRIPTS_PATH/**/.*.template; do
    if [ -z "$PROJECT_SLUG" ] || [ -z "$BUILD_IMAGE" ]; then
        echo "Missing template variable."
        exit 1
    fi

    sed -e "s~{{PROJECT_SLUG}}~$PROJECT_SLUG~g" \
        -e "s~{{BUILD_IMAGE}}~$BUILD_IMAGE~g" \
        $FILE > ${FILE%.*}
    rm $FILE
done

# Commit new versions of scripts and files.
git add -A .
git commit -m "Update ms-scaffolder scripts and files."

# Re-add .gitignore for ms-scaffolder scripts.
cat << EOF > scripts/ms-scaffolder/.gitignore
*
!.gitignore
EOF
git add -A .
git commit -m "Restore $SCRIPTS_PATH/.gitignore file."

# Remove temporary directory
rm -r $TEMP_DIR

