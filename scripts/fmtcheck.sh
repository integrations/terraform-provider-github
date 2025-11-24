#!/usr/bin/env bash

# Check formatting
echo "==> Checking that code complies with fmt requirements..."
fmt_diff="$(golangci-lint fmt --diff ./...)"
if [[ -n "${fmt_diff}" ]]; then
    echo 'formatting is incorrect:'
    echo "${fmt_diff}"
    echo "You can use the command: \`make fmt\` to reformat code."
    exit 1
fi

exit 0
