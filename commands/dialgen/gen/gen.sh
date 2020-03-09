#!/usr/bin/env bash

MAINFOLDER=$1
INFOLDER=$2
FF=$3
OUTFOLDER=$4


pushd $INFOLDER > /dev/null
go run ${MAINFOLDER}/main.go ardupilotmega.xml > ${OUTFOLDER}/ardupilotmega/dialect.go
go run ${MAINFOLDER}/main.go ASLUAV.xml > ${OUTFOLDER}/ASLUAV/dialect.go
go run ${MAINFOLDER}/main.go autoquad.xml > ${OUTFOLDER}/autoquad/dialect.go
go run ${MAINFOLDER}/main.go common.xml > ${OUTFOLDER}/common/dialect.go
go run ${MAINFOLDER}/main.go icarous.xml > ${OUTFOLDER}/icarous/dialect.go
go run ${MAINFOLDER}/main.go matrixpilot.xml > ${OUTFOLDER}/matrixpilot/dialect.go
go run ${MAINFOLDER}/main.go paparazzi.xml > ${OUTFOLDER}/paparazzi/dialect.go
go run ${MAINFOLDER}/main.go slugs.xml > ${OUTFOLDER}/slugs/dialect.go
go run ${MAINFOLDER}/main.go standard.xml > ${OUTFOLDER}/standard/dialect.go
go run ${MAINFOLDER}/main.go ualberta.xml > ${OUTFOLDER}/ualberta/dialect.go
go run ${MAINFOLDER}/main.go uAvionix.xml > ${OUTFOLDER}/uAvionix/dialect.go
go run ${MAINFOLDER}/main.go ${FF}/mavlink_freightfish.xml > ${OUTFOLDER}/mavlink_freightfish/dialect.go
popd > /dev/null
