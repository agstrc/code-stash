#!/usr/bin/env bash
# This script uses rclone to automatically upload the given directories to a specific rclone remote.
set -e

tty -s || { # Only uses a separate log file if we are in not attached to a terminal
    exec &>> "${HOME}/.rclone-sync.log"
    date --rfc-3339=s # Logs the date
}

RCLONE_REMOTE=drive
DIRECTORIES=(
    "/etc/zsh zsh"
    "${HOME}/.ssh ssh"
    "${HOME}/Documents linux-documents"
)
COMMANDS=(
    zip
    rclone
    ping
)

for command in "${COMMANDS[@]}"; do
    command -v "${command}" &>/dev/null || { echo "Command ${command} required but not found" && exit 1; }
done

until ping -c 1 1.1.1.1 &>/dev/null; do # Checks for connectivity by pinging Cloudflare's DNS
    echo 'Waiting for connectivity'
    sleep 1s
done

rclone mkdir -v "${RCLONE_REMOTE}:rclone"

for element in "${DIRECTORIES[@]}"; do
    read -r -a dir_arr <<< "$element"  # Splits the entry into an array, separating each element by spaces
    local_path="${dir_arr[0]}"
    remote_path="${RCLONE_REMOTE}:rclone/${dir_arr[1]}.zip"

    cd "${local_path}" || { echo "Could not change to directory ${local_path}" && exit 1; }
    echo "Zipping ${local_path} into ${remote_path}"
    zip -r - ./ 2>/dev/null | rclone rcat "${remote_path}"
done
