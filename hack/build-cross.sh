#!/bin/bash

# Build all cross compile targets and the base binaries
STARTTIME=$(date +%s)
source "$(dirname "${BASH_SOURCE}")/lib/init.sh"

host_platform="$(os::build::host_platform)"

server_platforms=( "${OS_BUILD_SERVER_PLATFORMS:-"${OS_BUILD_SUPPORTED_SERVER_PLATFORMS[@]}"}" )
image_platforms=( "${OS_BUILD_IMAGE_PLATFORMS:-"${OS_BUILD_SUPPORTED_IMAGE_PLATFORMS[@]}"}" )
test_platforms=( "${host_platform}" )
client_platforms=( "${OS_BUILD_CLIENT_PLATFORMS:-"${OS_BUILD_SUPPORTED_CLIENT_PLATFORMS[@]}"}" )

# Build image binaries for a subset of platforms. Image binaries are currently
# linux-only, and a subset of them are compiled with flags to make them static
# for use in Docker images "FROM scratch".
OS_BUILD_PLATFORMS=("${image_platforms[@]+"${image_platforms[@]}"}")
os::build::build_static_binaries "${OS_SCRATCH_IMAGE_COMPILE_TARGETS_LINUX[@]-}"
os::build::build_binaries "${OS_IMAGE_COMPILE_TARGETS_LINUX[@]-}"

# Build the server binaries
OS_BUILD_PLATFORMS=("${server_platforms[@]+"${server_platforms[@]}"}")
os::build::build_binaries "${OS_CROSS_COMPILE_SERVER_TARGETS[@]}"

# Build the client binaries
OS_BUILD_PLATFORMS=("${client_platforms[@]+"${client_platforms[@]}"}")
os::build::build_binaries "${OS_CROSS_COMPILE_CLIENT_TARGETS[@]}"

# Build the test binaries
OS_BUILD_PLATFORMS=("${test_platforms[@]+"${test_platforms[@]}"}")
os::build::build_binaries "${OS_TEST_TARGETS[@]}"

if [[ "${OS_BUILD_RELEASE_ARCHIVES-}" != "n" ]]; then
  # Make the server release.
  OS_BUILD_PLATFORMS=("${server_platforms[@]+"${server_platforms[@]}"}")
  OS_RELEASE_ARCHIVE="openshift-origin" \
    os::build::place_bins "${OS_CROSS_COMPILE_SERVER_BINARIES[@]}"

  # Make the client release.
  OS_BUILD_PLATFORMS=("${client_platforms[@]+"${client_platforms[@]}"}")
  OS_RELEASE_ARCHIVE="openshift-origin" \
    os::build::place_bins "${OS_CROSS_COMPILE_CLIENT_BINARIES[@]}"

  # Make the image binaries release.
  OS_BUILD_PLATFORMS=("${image_platforms[@]+"${image_platforms[@]}"}")
  OS_RELEASE_ARCHIVE="openshift-origin-image" \
    os::build::place_bins "${OS_IMAGE_COMPILE_BINARIES[@]}"

  os::build::release_sha
else
  # Place binaries only
  OS_BUILD_PLATFORMS=("${server_platforms[@]+"${server_platforms[@]}"}") \
    os::build::place_bins "${OS_CROSS_COMPILE_SERVER_BINARIES[@]}"
  OS_BUILD_PLATFORMS=("${client_platforms[@]+"${client_platforms[@]}"}") \
    os::build::place_bins "${OS_CROSS_COMPILE_CLIENT_BINARIES[@]}"
  OS_BUILD_PLATFORMS=("${image_platforms[@]+"${image_platforms[@]}"}") \
    os::build::place_bins "${OS_IMAGE_COMPILE_BINARIES[@]}"
fi

if [[ "${OS_GIT_TREE_STATE:-dirty}" == "clean"  ]]; then
  # only when we are building from a clean state can we claim to
  # have created a valid set of binaries that can resemble a release
  mkdir -p "${OS_OUTPUT_RELEASEPATH}"
  echo "${OS_GIT_COMMIT}" > "${OS_OUTPUT_RELEASEPATH}/.commit"
fi

ret=$?; ENDTIME=$(date +%s); echo "$0 took $(($ENDTIME - $STARTTIME)) seconds"; exit "$ret"
