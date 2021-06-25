#!/usr/bin/env python3
# portcheck is a quick and non-thorough to check for open ports. It attempts to create a simple connection or waits
# until timeout.

import socket
import argparse

parser = argparse.ArgumentParser(
    description="portcheck implements a quick (and non-thorough) check for open ports."
)
parser.add_argument("address")
parser.add_argument("port", type=int)

parser.add_argument(
    "-t",
    "--timeout",
    type=int,
    default=15,
    help="Time (in seconds) to be waited when attempint a connection. Defaults to 15 seconds.",
)

args = parser.parse_args()

with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as sock:
    sock.settimeout(args.timeout)

    try:
        sock.connect((args.address, args.port))
        print("Port is open.")
        sock.close()
    except socket.timeout:
        print("Server connection timed out.")
        exit(1)
