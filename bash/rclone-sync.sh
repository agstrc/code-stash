#!/usr/bin/env bash
# This script uses rclone to automatically sync directories to a given rclone remote. It deals with tar files in order
# to avoid long syncing times
set -e

tty -s || { # Only uses a separate log file if we are in not attached to a terminal
    exec &>> "${HOME}/.rclone-sync.log"
    date --rfc-3339=s # Logs the date
}

RCLONE_REMOTE=drive
LOCAL_TMP=/tmp/rclone-sync

mkdir -p "${LOCAL_TMP}"
trap 'rm -rf ${LOCAL_TMP}' exit # cleanup temporary files

# Directories to be synced followed by their desired file names in the remote
DIRECTORIES=("/etc/zsh zsh" "${HOME}/.ssh ssh" "${HOME}/Documents linux-documents")

COMMANDS=(tar rclone ping)

for command in "${COMMANDS[@]}"; do
    # Checks if all required commands are present
    command -v "${command}" &>/dev/null || { echo "Command ${command} required but not found" && exit 1; }
done

until ping -c 1 1.1.1.1 &>/dev/null; do
    # Checks for connectivity by pinging Cloudflare's DNS
    echo 'Waiting for connectivity'
    sleep 1s
done

rclone mkdir -v "${RCLONE_REMOTE}:rclone"

for element in "${DIRECTORIES[@]}"; do
    read -r -a dir_arr <<< "$element"  # Splits the entry into an array, separating each element by spaces
    remote_path="${RCLONE_REMOTE}:rclone/${dir_arr[1]}.tar"
    local_path="${dir_arr[0]}"

    tar_path="${LOCAL_TMP}/${dir_arr[1]}.tar"

    tar cf "${tar_path}" -C "${local_path}" .
    rclone check "${tar_path}" drive:rclone/ || {
        # If the files differ, we upload the new copy
        rclone copyto -v "${tar_path}" "${remote_path}"
    }

    rm "${tar_path}"
done
