#!/bin/bash -l

# Parse Action inputs into environment variables
export RUN_ALL=${INPUT_RUN_ALL}
export RUN_ALLOWED=${INPUT_RUN_ALLOWED}
export TF_LOG=${INPUT_TF_LOG}
export GITHUB_ORGANIZATION=${INPUT_GITHUB_ORGANIZATION}
export GITHUB_BASE_URL=${INPUT_GITHUB_BASE_URL}
export GITHUB_OWNER=${INPUT_GITHUB_OWNER}
export GITHUB_TEST_USER=${INPUT_GITHUB_TEST_USER}
export GITHUB_TEST_OWNER=${INPUT_GITHUB_TEST_OWNER}
export GITHUB_TEST_ORGANIZATION=${INPUT_GITHUB_TEST_ORGANIZATION}
export GITHUB_TEST_USER_TOKEN=${INPUT_GITHUB_TEST_USER_TOKEN}
export GITHUB_TEST_USER_NAME=${INPUT_GITHUB_TEST_USER_NAME}
export GITHUB_TEST_USER_EMAIL=${INPUT_GITHUB_TEST_USER_EMAIL}
export GITHUB_TEST_COLLABORATOR=${INPUT_GITHUB_TEST_COLLABORATOR}
export GITHUB_TEST_COLLABORATOR_TOKEN=${INPUT_GITHUB_TEST_COLLABORATOR_TOKEN}
export GITHUB_TEMPLATE_REPOSITORY=${INPUT_GITHUB_TEMPLATE_REPOSITORY}
export GITHUB_TEMPLATE_REPOSITORY_RELEASE_ID=${INPUT_GITHUB_TEMPLATE_REPOSITORY_RELEASE_ID}

# Set GITHUB_TOKEN
export GITHUB_TOKEN="${GITHUB_TEST_USER_TOKEN}"

# Acceptance Test Functions

generate_test_fixtures () {
  openssl req -x509 -newkey rsa:4096 -days 1 -nodes \
    -subj "/C=US/ST=CA/L=San Francisco/O=HashiCorp, Inc./CN=untrusted" \
    -keyout github/test-fixtures/key.pem -out github/test-fixtures/cert.pem
}

modified_files () {
  git show --pretty="" --name-only HEAD | tr '\n' ' '
}

test_files_for_modified_files () {
  for f in $(modified_files); do
    find . | grep $(basename -s .go $f)
  done | grep _test | sort | uniq | tr '\n' ' '
}

test_cases_from_modified_files () {
  if [ -z "$(test_files_for_modified_files)" ]; then
    return
  else
    grep -nr "func Test" $(test_files_for_modified_files) | \
    cut -d ' ' -f 2 | cut -d "(" -f 1 | grep -e TestAcc -e TestProvider | \
    tr '\n' ' '
  fi
}

all_test_cases () {
  grep -nr "func Test" . | grep -v vendor | \
  cut -d ' ' -f 2 | cut -d "(" -f 1 | grep -e TestAcc -e TestProvider | \
  tr '\n' ' '
}

test_cases () {
  if ! [ "$RUN_ALLOWED" = "" ]; then
    echo $RUN_ALLOWED
  elif [ "$RUN_ALL" = "true" ]; then
    all_test_cases
  else
    test_cases_from_modified_files
  fi
}

run_test () {
  # FIXME: Running one test case per UNIX process yields less flaky results
  TF_LOG=${INPUT_TF_LOG} TF_ACC=1 go test -v -timeout 30m  ./... -run $1
  return $?
}

failed_test_cases () {
  env | grep "test_case_failed_" | sed 's/test_case_failed_//' | cut -d= -f1 | tr '\n' ' '
}

main () {

  # Exit early if no test cases will run
  if [ -z "$(test_cases)" ]; then
    echo "No test cases eligible to run, exiting."
    return 0
  fi

  # Setup Go Environment
  go mod init
  go mod tidy
  go mod vendor

  # Pre-Sweeper
  go test -v -sweep="gh-region"

  generate_test_fixtures

  for test_case in $(test_cases); do
    unset test_case_failed_${test_case}
    if ! run_test $test_case; then
      export test_case_failed_${test_case}=1
    fi
  done

  # Post-Sweeper
  go test -v -sweep="gh-region"

  # Output failed test cases
  echo "::set-output name=failed::$(failed_test_cases)"

  # Exit with a failure if any test cases failed
  for failed_test_case in $(env | grep "test_case_failed_"); do
    exit 1
  done

}

main $@
