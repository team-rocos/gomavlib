#!/usr/bin/env bash

MAINFOLDER=$(realpath $(dirname "$(readlink -f "$0/../..")"))
GENFOLDER=${MAINFOLDER}/commands/dialgen
INFOLDER=${MAINFOLDER}/mavlink-upstream/message_definitions/v1.0
FF=${MAINFOLDER}/message_definitions/v1.0
OUTFOLDER=${MAINFOLDER}/dialects

pushd $INFOLDER > /dev/null
go run ${GENFOLDER}/main.go ${GENFOLDER}/defdecoder.go ardupilotmega.xml > ${OUTFOLDER}/ardupilotmega/dialect.go
go run ${GENFOLDER}/main.go ${GENFOLDER}/defdecoder.go ASLUAV.xml > ${OUTFOLDER}/ASLUAV/dialect.go
go run ${GENFOLDER}/main.go ${GENFOLDER}/defdecoder.go autoquad.xml > ${OUTFOLDER}/autoquad/dialect.go
go run ${GENFOLDER}/main.go ${GENFOLDER}/defdecoder.go common.xml > ${OUTFOLDER}/common/dialect.go
go run ${GENFOLDER}/main.go ${GENFOLDER}/defdecoder.go icarous.xml > ${OUTFOLDER}/icarous/dialect.go
go run ${GENFOLDER}/main.go ${GENFOLDER}/defdecoder.go matrixpilot.xml > ${OUTFOLDER}/matrixpilot/dialect.go
go run ${GENFOLDER}/main.go ${GENFOLDER}/defdecoder.go minimal.xml > ${OUTFOLDER}/minimal/dialect.go
go run ${GENFOLDER}/main.go ${GENFOLDER}/defdecoder.go paparazzi.xml > ${OUTFOLDER}/paparazzi/dialect.go
go run ${GENFOLDER}/main.go ${GENFOLDER}/defdecoder.go slugs.xml > ${OUTFOLDER}/slugs/dialect.go
go run ${GENFOLDER}/main.go ${GENFOLDER}/defdecoder.go standard.xml > ${OUTFOLDER}/standard/dialect.go
go run ${GENFOLDER}/main.go ${GENFOLDER}/defdecoder.go ualberta.xml > ${OUTFOLDER}/ualberta/dialect.go
go run ${GENFOLDER}/main.go ${GENFOLDER}/defdecoder.go uAvionix.xml > ${OUTFOLDER}/uAvionix/dialect.go
go run ${GENFOLDER}/main.go ${GENFOLDER}/defdecoder.go ${FF}/mavlink_freightfish.xml > ${OUTFOLDER}/mavlink_freightfish/dialect.go
popd > /dev/null
