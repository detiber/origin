#!/bin/bash

# This script builds the release images for use by the release build and image builds.

STARTTIME=$(date +%s)
source "$(dirname "${BASH_SOURCE}")/lib/init.sh"

# determine the correct tag prefix
tag_prefix="${OS_IMAGE_PREFIX:-"openshift/origin"}"

ret=0

# IF OS_BUILD_ARCHES is not specified, default to the host architecture
image_arches="${OS_BUILD_ARCHES:-$(os::build::host_arch)}"
image_basename="${tag_prefix}-release"
image_tag="golang-${OS_BUILD_ENV_GOLANG}"
image_dir="${OS_ROOT}/images/release/${image_tag}"

os::build::cross_images ${image_basename} ${image_dir} ${image_arches} ${image_tag} || (ret=1 && break)

ret=$?; ENDTIME=$(date +%s); echo "$0 took $(($ENDTIME - $STARTTIME)) seconds"; exit "$ret"
