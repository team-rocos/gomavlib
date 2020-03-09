#!/usr/bin/env bash

MAINFOLDER=$(realpath $(dirname "$(readlink -f "$0/../..")"))
GENFOLDER=${MAINFOLDER}/commands/dialgen/gen
INFOLDER=${MAINFOLDER}/mavlink-upstream/message_definitions/v1.0
FF=${MAINFOLDER}/message_definitions/v1.0
OUTFOLDER=${MAINFOLDER}/dialects

pushd $INFOLDER > /dev/null
go run ${GENFOLDER}/main.go ardupilotmega.xml > ${OUTFOLDER}/ardupilotmega/dialect.go
go run ${GENFOLDER}/main.go ASLUAV.xml > ${OUTFOLDER}/ASLUAV/dialect.go
go run ${GENFOLDER}/main.go autoquad.xml > ${OUTFOLDER}/autoquad/dialect.go
go run ${GENFOLDER}/main.go common.xml > ${OUTFOLDER}/common/dialect.go
go run ${GENFOLDER}/main.go icarous.xml > ${OUTFOLDER}/icarous/dialect.go
go run ${GENFOLDER}/main.go matrixpilot.xml > ${OUTFOLDER}/matrixpilot/dialect.go
go run ${GENFOLDER}/main.go minimal.xml > ${OUTFOLDER}/minimal/dialect.go
go run ${GENFOLDER}/main.go paparazzi.xml > ${OUTFOLDER}/paparazzi/dialect.go
go run ${GENFOLDER}/main.go slugs.xml > ${OUTFOLDER}/slugs/dialect.go
go run ${GENFOLDER}/main.go standard.xml > ${OUTFOLDER}/standard/dialect.go
go run ${GENFOLDER}/main.go ualberta.xml > ${OUTFOLDER}/ualberta/dialect.go
go run ${GENFOLDER}/main.go uAvionix.xml > ${OUTFOLDER}/uAvionix/dialect.go
go run ${GENFOLDER}/main.go ${FF}/mavlink_freightfish.xml ${INFOLDER}/common.xml > ${OUTFOLDER}/mavlink_freightfish/dialect.go
popd > /dev/null
