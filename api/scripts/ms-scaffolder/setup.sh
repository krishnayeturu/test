#!/bin/bash
set -e
shopt -s nullglob
shopt -s globstar

TEMP_DIR="$(dirname -- "$(readlink -f -- "$BASH_SOURCE")")"
PROJECT_DIR="$TEMP_DIR/.."
SCRIPT_DIR="$PROJECT_DIR/scripts/ms-scaffolder"

# Ensure git is clean before proceeding.
if [ ! -z "$(git status --porcelain)" ]; then
    echo "Git must be clean before running this command. Clean up your changes and try again."
    exit 1
fi

# If a 1st argument is specified, assume it is the template slug. Otherwise ask user for input.
if [ -z "$1" ]; then
    # Ask user for project template.
    PS3="Please select a project to initialize: "
    OPTIONS=("Go API" "Go Integration" "Java Commandler" "Cancel project initialization")
    select OPT in "${OPTIONS[@]}"; do
        case $OPT in
            "Go API")
                TEMPLATE_DIR="go-api"
                break
                ;;
            "Go Integration")
                TEMPLATE_DIR="go-int"
                break
                ;;
            "Java Commandler")
                TEMPLATE_DIR="java-ch"
                break
                ;;
            "Cancel project initialization")
                echo "Project initialization cancelled."
                exit 0
                ;;
            *)
                echo "$REPLY is an invalid option."
                ;;
        esac
    done
else
    TEMPLATE_DIR="$1"
fi

# If a 2nd argument is specified, assume it is the project slug. Otherwise ask user for input.
if [ -z "$2" ]; then
    # Ask user for project slug.
    echo "What is the hyphenated slug for your project? e.g., ms-api-commander"
    read PROJECT_SLUG
else
    PROJECT_SLUG="$2"
fi
PROJECT_SLUG_NO_HYPHEN="${PROJECT_SLUG//-/}"

# Copy template project to root directory and delete others.
mv $PROJECT_DIR/templates $PROJECT_DIR/.ms-scaffolder-temp # In case there is a directory named "templates" in any of our template projects.
cp $PROJECT_DIR/.ms-scaffolder-temp/$TEMPLATE_DIR/.env.template $PROJECT_DIR/
cp $PROJECT_DIR/.ms-scaffolder-temp/$TEMPLATE_DIR/.gitignore $PROJECT_DIR/
cp $PROJECT_DIR/.ms-scaffolder-temp/$TEMPLATE_DIR/.gitlab-ci.yml.template $PROJECT_DIR/
cp -r $PROJECT_DIR/.ms-scaffolder-temp/$TEMPLATE_DIR/* $PROJECT_DIR/
rm -r $PROJECT_DIR/.ms-scaffolder-temp

# Change helm chart path name to project slug.
mv $PROJECT_DIR/deploy/helm/chart $PROJECT_DIR/deploy/helm/$PROJECT_SLUG

# Load the copied .env variables from the template.
source $PROJECT_DIR/.env.template

# Replace variables in template files and remove the template files.
for FILE in $PROJECT_DIR/**/*.template $PROJECT_DIR/**/.*.template; do
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

# Add setup variables to .env file.
echo "PROJECT_SLUG=$PROJECT_SLUG" >> $PROJECT_DIR/.env
echo "PROJECT_SLUG_NO_HYPHEN=$PROJECT_SLUG_NO_HYPHEN" >> $PROJECT_DIR/.env

# Copy .env back to .env.template so it is committed to the project.
cp $PROJECT_DIR/.env $PROJECT_DIR/.env.template

# Run template-specific setup script.
$PROJECT_DIR/scripts/ms-scaffolder/template-setup.sh
rm $PROJECT_DIR/scripts/ms-scaffolder/template-setup.sh

# Apply do not edit text to appropriate files and remove.
for FILE in $PROJECT_DIR/scripts/ms-scaffolder/*; do
    echo -e "$(cat $PROJECT_DIR/donotedit.txt)\n\n$(cat $FILE)" > $FILE
done
echo -e "$(cat $PROJECT_DIR/donotedit.txt)\n\n$(cat $PROJECT_DIR/Makefile)" > $PROJECT_DIR/Makefile
rm $PROJECT_DIR/donotedit.txt

# Commit scaffolding changes.
git add -A .
git commit -m "Initial project scaffolding."

echo "Repository setup complete and committed to current working branch. If you'd like to push, please run 'git push'."

# TODO: Make calls here to GitLab API to set up repository settings.

# Remove temporary directory
rm -r $TEMP_DIR