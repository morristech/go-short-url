#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

my_dir="$(dirname "$0")"

docker pull meio/go-short-url-server:latest
docker run --rm --name go-short-url-server -u 0 -p 5000:5000 -it meio/go-short-url-server:latest
