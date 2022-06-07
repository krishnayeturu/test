#!/bin/bash
set -e
shopt -s nullglob
shopt -s globstar

# Ensure git is clean before proceeding.
if [ ! -z "$(git status --porcelain)" ]; then
    echo "Git must be clean before running this command. Clean up your changes and try again."
    exit 1
fi

# Select a project template
PS3="Please select a project to initialize: "
OPTIONS=("Go API" "Java Commandler" "Cancel project initialization")
select OPT in "${OPTIONS[@]}"; do
    case $OPT in
        "Go API")
            PROJECT_DIR="go-api"
            break
            ;;
        "Java Commandler")
            PROJECT_DIR="java-ch"
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

# Collect parameters from user.
echo "What is the hyphenated slug for your project? e.g., ms-api-commander-writer"
read PROJECT_SLUG

# Copy template project to root directory and delete others.
mv templates .ms-scaffolder-temp # In case there is a directory named "templates" in any of our template projects.
cp .ms-scaffolder-temp/$PROJECT_DIR/.env ./
cp .ms-scaffolder-temp/$PROJECT_DIR/.gitignore ./
cp .ms-scaffolder-temp/$PROJECT_DIR/.gitlab-ci.yml.template ./
cp -r .ms-scaffolder-temp/$PROJECT_DIR/* ./
rm -r .ms-scaffolder-temp

# Load the copied .env variables from the template.
source .env

# Replace variables in template files and remove the template files.
for FILE in **/*.template **/.*.template; do
    if [ -z "$PROJECT_SLUG" ] || [ -z "$BUILD_IMAGE" ]; then
        echo "Missing template variable."
        exit 1
    fi

    sed -e "s~{{PROJECT_SLUG}}~$PROJECT_SLUG~g" \
        -e "s~{{BUILD_IMAGE}}~$BUILD_IMAGE~g" \
        $FILE > ${FILE%.*}
    rm $FILE
done

# Add setup variables to .env file.
echo "" >> .env # Ensure newline at end of .env
echo "PROJECT_SLUG=$PROJECT_SLUG" >> .env

# Commit scaffolding changes.
git add -A .
git commit -m "Initial project scaffolding."

# Add files that shouldn't be modified to .gitignore
cat << EOF > scripts/ms-scaffolder/.gitignore
*
!.gitignore
EOF

# Commit changes .gitignore changes.
git add -A .
git commit -m "Update .gitignore with appropriate files."

echo "Repository setup complete and committed to current working branch. If you'd like to push, please run 'git push'."
# TODO: Make calls here to GitLab API to set up repository settings.