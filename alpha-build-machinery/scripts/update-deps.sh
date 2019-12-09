#!/bin/bash -e

readonly GO_MINOR_VERSION=12
readonly REQUIRED_GO_VERSION="1.$GO_MINOR_VERSION"

function verify_go_version() {
	if ! command -v go &> /dev/null; then
		echo "[FATAL] go was not found in \$PATH. Please install version ${REQUIRED_GO_VERSION} or newer."
		exit 1
	fi

	local go_version
    go_version=($(go version))
	if ! echo "${go_version[2]#go}" | awk -F. -v min=$GO_MINOR_VERSION '{ exit $2 < min }'; then
		echo "Detected go version: ${go_version[*]}."
		echo "Please install go version ${REQUIRED_GO_VERSION} or newer."
		exit 1
	fi
}

verify_go_version

go mod vendor
