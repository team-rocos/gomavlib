#!/usr/bin/env bash

MAINFOLDER=$(realpath $(dirname "$(readlink -f "$0/../..")"))
GENFOLDER=${MAINFOLDER}/commands/dialgen/gen
INFOLDER=${MAINFOLDER}/mavlink-upstream/message_definitions/v1.0
FF=${MAINFOLDER}/message_definitions/v1.0
OUTFOLDER=${MAINFOLDER}/dialects

pushd $INFOLDER > /dev/null
go run ${GENFOLDER}/main.go ardupilotmega.xml ${INFOLDER} > ${OUTFOLDER}/ardupilotmega/dialect.go
go run ${GENFOLDER}/main.go ASLUAV.xml ${INFOLDER} > ${OUTFOLDER}/ASLUAV/dialect.go
go run ${GENFOLDER}/main.go autoquad.xml ${INFOLDER} > ${OUTFOLDER}/autoquad/dialect.go
go run ${GENFOLDER}/main.go common.xml ${INFOLDER} > ${OUTFOLDER}/common/dialect.go
go run ${GENFOLDER}/main.go icarous.xml ${INFOLDER} > ${OUTFOLDER}/icarous/dialect.go
go run ${GENFOLDER}/main.go matrixpilot.xml ${INFOLDER} > ${OUTFOLDER}/matrixpilot/dialect.go
go run ${GENFOLDER}/main.go minimal.xml ${INFOLDER} > ${OUTFOLDER}/minimal/dialect.go
go run ${GENFOLDER}/main.go paparazzi.xml ${INFOLDER} > ${OUTFOLDER}/paparazzi/dialect.go
go run ${GENFOLDER}/main.go slugs.xml ${INFOLDER} > ${OUTFOLDER}/slugs/dialect.go
go run ${GENFOLDER}/main.go standard.xml ${INFOLDER} > ${OUTFOLDER}/standard/dialect.go
go run ${GENFOLDER}/main.go ualberta.xml ${INFOLDER} > ${OUTFOLDER}/ualberta/dialect.go
go run ${GENFOLDER}/main.go uAvionix.xml ${INFOLDER} > ${OUTFOLDER}/uAvionix/dialect.go
go run ${GENFOLDER}/main.go ${FF}/mavlink_freightfish.xml ${INFOLDER} > ${OUTFOLDER}/mavlink_freightfish/dialect.go
popd > /dev/null
