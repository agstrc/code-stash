#!/bin/env bash
# This script allows you to copy the output of a given command to your X clipboard. It accepts either CLI arguments
# or reads the desired command from stdin.

# Exiting via sigint is expected within this script
trap 'exit 0' INT

function empty_args {
    read COMMAND
    eval "${COMMAND}" | xclip -f -sel clip
}

function cli_args {
    eval "$@" | xclip -f -sel clip
}

if [[ $# != 0 ]]; then
    cli_args "$@"
    exit "$?"
fi

empty_args
