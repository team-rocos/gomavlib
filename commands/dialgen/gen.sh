#!/usr/bin/env bash

MAINFOLDER=$1
INFOLDER=$2
FF=$3
OUTFOLDER=$4


pushd $INFOLDER > /dev/null
go run ${MAINFOLDER}/main.go ${MAINFOLDER}/defdecoder.go ardupilotmega.xml > ${OUTFOLDER}/ardupilotmega/dialect.go
go run ${MAINFOLDER}/main.go ${MAINFOLDER}/defdecoder.go ASLUAV.xml > ${OUTFOLDER}/ASLUAV/dialect.go
go run ${MAINFOLDER}/main.go ${MAINFOLDER}/defdecoder.go autoquad.xml > ${OUTFOLDER}/autoquad/dialect.go
go run ${MAINFOLDER}/main.go ${MAINFOLDER}/defdecoder.go common.xml > ${OUTFOLDER}/common/dialect.go
go run ${MAINFOLDER}/main.go ${MAINFOLDER}/defdecoder.go icarous.xml > ${OUTFOLDER}/icarous/dialect.go
go run ${MAINFOLDER}/main.go ${MAINFOLDER}/defdecoder.go matrixpilot.xml > ${OUTFOLDER}/matrixpilot/dialect.go
go run ${MAINFOLDER}/main.go ${MAINFOLDER}/defdecoder.go paparazzi.xml > ${OUTFOLDER}/paparazzi/dialect.go
go run ${MAINFOLDER}/main.go ${MAINFOLDER}/defdecoder.go slugs.xml > ${OUTFOLDER}/slugs/dialect.go
go run ${MAINFOLDER}/main.go ${MAINFOLDER}/defdecoder.go standard.xml > ${OUTFOLDER}/standard/dialect.go
go run ${MAINFOLDER}/main.go ${MAINFOLDER}/defdecoder.go ualberta.xml > ${OUTFOLDER}/ualberta/dialect.go
go run ${MAINFOLDER}/main.go ${MAINFOLDER}/defdecoder.go uAvionix.xml > ${OUTFOLDER}/uAvionix/dialect.go
go run ${MAINFOLDER}/main.go ${MAINFOLDER}/defdecoder.go ${FF}/mavlink_freightfish.xml > ${OUTFOLDER}/mavlink_freightfish/dialect.go
popd > /dev/null
