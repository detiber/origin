#!/bin/bash

# This script builds the base and release images for use by the release build and image builds.

STARTTIME=$(date +%s)
source "$(dirname "${BASH_SOURCE}")/lib/init.sh"

# determine the correct tag prefix
tag_prefix="${OS_IMAGE_PREFIX:-"openshift/origin"}"

os::util::ensure::gopath_binary_exists imagebuilder

orig_build_image_args="${OS_BUILD_IMAGE_ARGS:-}"
host_arch=$(arch)

case $host_arch in
aarch64)
  OS_BUILD_IMAGE_ARGS+=' --from centos/aarch64:7'
  ;;
ppc64le)
  OS_BUILD_IMAGE_ARGS+=' --from ppc64le/centos:7'
  ;;
s390x)
  OS_BUILD_IMAGE_ARGS+=' --from sinenomine/clefos-base-s390x:latest'
  ;;
esac

# Build the base image without the default image args
os::build::image "${tag_prefix}-source" "${OS_ROOT}/images/source"

OS_BUILD_IMAGE_ARGS=${orig_build_image_args:-}

os::build::image "${tag_prefix}-base"   "${OS_ROOT}/images/base"

ret=$?; ENDTIME=$(date +%s); echo "$0 took $(($ENDTIME - $STARTTIME)) seconds"; exit "$ret"
