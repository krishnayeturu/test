# Microservice Scaffolder

This project is intended to be used as a template when creating any new microservice projects. By
using this template, we will be able to keep many things consistent across our many microservices,
including build processes, CI/CD, project structure, etc.

# System Dependencies

- `docker` >= 20.10.x
- `git` >= 2.x
- `make` >= 4.x

# Usage

Note for Windows Users: You must run `make` from PowerShell.

1. Create a new project in GitLab using this project as the template. Do not configure anything in
the project settings after creation, the setup process will handle this.
1. Clone your newly created project to your local machine.
1. Run `make setup <template>` (a full list of templates can be seen by running `make setup help`).
1. Enter any parameters if prompted.
1. Congratulations, your project is fully setup for developing!

# What Does Setup Actually Do?

The `make setup` command performs several operations to set up the initial project:

- It gathers any parameters it needs from you and replaces them where necessary throughout the
project files.
- It commits all of the changes it makes to your repository automatically.

# Template Requirements
- .env file specifying BUILD_IMAGE
- .gitlab-ci.yml.template for CI/CD
- .gitignore
