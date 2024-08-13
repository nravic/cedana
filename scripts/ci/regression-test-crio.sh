#!/bin/bash -e
# Used to run regression bats tests (located in tests/regression)

source ./helpers.sh

function start_regression_crio() {
    echo "Running regression tests in cwd: $(pwd)"
    bats test/regression-crio/main.bats
}

main() {
    pushd ../..
    print_env
    source_env
    start_regression_crio

    popd
}

main