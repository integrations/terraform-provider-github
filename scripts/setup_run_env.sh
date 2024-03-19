#!/bin/bash

# Check if launch.json already exists
if [ -f ".vscode/launch.json" ]; then
    read -p "launch.json already exists. Overwrite? (y/n): " overwrite
    if [[ $overwrite != "y" ]]; then
        echo "Exiting without creating launch.json."
        exit 1
    fi
fi

# Prompt for environment variables
read -p "Enter TF_TEST_FILE: " TF_TEST_FILE
read -p "Enter TF_TEST_FUNCTION: " TF_TEST_FUNCTION
read -p "Enter GITHUB_TEST_COLLABORATOR: " GITHUB_TEST_COLLABORATOR
read -p "Enter GITHUB_TEST_COLLABORATOR_TOKEN: " GITHUB_TEST_COLLABORATOR_TOKEN
read -p "Enter GITHUB_TEST_USER: " GITHUB_TEST_USER
read -p "Enter GITHUB_TOKEN: " GITHUB_TOKEN
read -p "Enter GITHUB_TEMPLATE_REPOSITORY: " GITHUB_TEMPLATE_REPOSITORY
read -p "Enter GITHUB_TEMPLATE_REPOSITORY_RELEASE_ID: " GITHUB_TEMPLATE_REPOSITORY_RELEASE_ID
read -p "Enter GITHUB_ORGANIZATION: " GITHUB_ORGANIZATION
read -p "Enter TF_CLI_CONFIG_FILE: " TF_CLI_CONFIG_FILE
read -p "Enter TF_ACC: " TF_ACC
read -p "Enter TF_LOG: " TF_LOG
read -p "Enter APP_INSTALLATION_ID: " APP_INSTALLATION_ID

# Export environment variables
export TF_TEST_FILE="$TF_TEST_FILE"
export TF_TEST_FUNCTION="$TF_TEST_FUNCTION"
export GITHUB_TEST_COLLABORATOR="$GITHUB_TEST_COLLABORATOR"
export GITHUB_TEST_COLLABORATOR_TOKEN="$GITHUB_TEST_COLLABORATOR_TOKEN"
export GITHUB_TEST_USER="$GITHUB_TEST_USER"
export GITHUB_TOKEN="$GITHUB_TOKEN"
export GITHUB_TEMPLATE_REPOSITORY="$GITHUB_TEMPLATE_REPOSITORY"
export GITHUB_TEMPLATE_REPOSITORY_RELEASE_ID="$GITHUB_TEMPLATE_REPOSITORY_RELEASE_ID"
export GITHUB_ORGANIZATION="$GITHUB_ORGANIZATION"
export TF_CLI_CONFIG_FILE="$TF_CLI_CONFIG_FILE"
export TF_ACC="$TF_ACC"
export TF_LOG="$TF_LOG"
export APP_INSTALLATION_ID="$APP_INSTALLATION_ID"

# Create the launch.json file
cat << EOF > .vscode/launch.json
{
  "version": "0.2.0",
  "configurations": [
    {
      "name": "Launch test function",
      "type": "go",
      "request": "launch",
      "mode": "test",
      "program": "\${workspaceFolder}/github/\${env:TF_TEST_FILE}",
      "args": [
        "-test.v",
        "-test.run",
        "^\${env:TF_TEST_FILE}$"
      ],
      "env": {
        "GITHUB_TEST_COLLABORATOR": "\${env:GITHUB_TEST_COLLABORATOR}",
        "GITHUB_TEST_COLLABORATOR_TOKEN": "\${env:GITHUB_TEST_COLLABORATOR_TOKEN}",
        "GITHUB_TEST_USER": "\${env:GITHUB_TEST_USER}",
        "GITHUB_TOKEN": "\${env:GITHUB_TOKEN}",
        "GITHUB_TEMPLATE_REPOSITORY": "\${env:GITHUB_TEMPLATE_REPOSITORY}",
        "GITHUB_TEMPLATE_REPOSITORY_RELEASE_ID": "\${env:GITHUB_TEMPLATE_REPOSITORY_RELEASE_ID}",
        "GITHUB_ORGANIZATION": "\${env:GITHUB_ORGANIZATION}",
        "TF_CLI_CONFIG_FILE": "\${env:TF_CLI_CONFIG_FILE}",
        "TF_ACC": "\${env:TF_ACC}",
        "TF_LOG": "\${env:TF_LOG}",
        "APP_INSTALLATION_ID": "\${env:APP_INSTALLATION_ID}"
      }
    },
    {
      "name": "Attach to Process",
      "type": "go",
      "request": "attach",
      "mode": "local",
      "processId": "\${command:AskForProcessId}"
    }
  ]
}
EOF

echo "launch.json has been created."