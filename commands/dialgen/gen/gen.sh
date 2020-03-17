#!/usr/bin/env bash

MAINFOLDER=$(realpath $(dirname "$(readlink -f "$0/../..")"))
GENFOLDER=${MAINFOLDER}/commands/dialgen/gen
INFOLDER=${MAINFOLDER}/mavlink-upstream/message_definitions/v1.0
FF=${MAINFOLDER}/message_definitions/v1.0
OUTFOLDER=${MAINFOLDER}/dialects

generate () {
	NAME=${1##*/}
	NAME=${NAME%.*}
	echo "Processing ${NAME}"
	pushd $INFOLDER > /dev/null
	rm -rf ${OUTFOLDER}/${NAME}
	mkdir -p ${OUTFOLDER}/${NAME}
	go run ${GENFOLDER}/main.go ${1%.*}.xml false ${INFOLDER} > ${OUTFOLDER}/${NAME}/dialect.go
	go run ${GENFOLDER}/main.go ${1%.*}.xml true ${INFOLDER} > ${OUTFOLDER}/${NAME}/dialect_test.go
	popd > /dev/null
}

generate ardupilotmega
generate ASLUAV
generate autoquad
generate common
generate icarous
generate matrixpilot
generate minimal
generate paparazzi
generate slugs
generate standard
generate ualberta
generate uAvionix
generate ${FF}/mavlink_freightfish.xml

# ALL DONE.
