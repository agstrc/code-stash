#!/usr/bin/env bash
# This script automatically syncs the given directories to a configured rclone remote.
set -e

exec &>> "${HOME}/.rclone-sync.log"

date --rfc-3339=s # Logs the date

until ping -c 1 1.1.1.1 &>/dev/null; do # Checks for connectivity by pinging Cloudflare's DNS
    echo 'Waiting for connectivity'
    sleep 1s
done


RCLONE_REMOTE=drive
DIRECTORIES=(
    "/etc/zsh zsh"
    "${HOME}/.ssh ssh"
    "${HOME}/Documents linux-box.documents"
)

for element in "${DIRECTORIES[@]}"; do
    read -r -a dir_arr <<< "$element"  # Splits the entry into an array, separating each element by spaces
    local_path="${dir_arr[0]}"
    remote_path="${RCLONE_REMOTE}:rclone/${dir_arr[1]}"
    echo "Syncing ${local_path} into ${remote_path}"

    rclone mkdir -v "${remote_path}"
    rclone sync -v -L "${local_path}" "${remote_path}"
done
