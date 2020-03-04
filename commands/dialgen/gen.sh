#!/usr/bin/env bash

MAINFOLDER=$1
INFOLDER=$2
OUTFOLDER=$3

pushd $INFOLDER
go run ${MAINFOLDER}/main.go ${MAINFOLDER}/defdecoder.go ardupilotmega.xml > ${OUTFOLDER}/ardupilotmega.go
go run ${MAINFOLDER}/main.go ${MAINFOLDER}/defdecoder.go ASLUAV.xml > ${OUTFOLDER}/ASLUAV.go
go run ${MAINFOLDER}/main.go ${MAINFOLDER}/defdecoder.go autoquad.xml > ${OUTFOLDER}/autoquad.go
go run ${MAINFOLDER}/main.go ${MAINFOLDER}/defdecoder.go common.xml > ${OUTFOLDER}/common.go
go run ${MAINFOLDER}/main.go ${MAINFOLDER}/defdecoder.go icarous.xml > ${OUTFOLDER}/icarous.go
go run ${MAINFOLDER}/main.go ${MAINFOLDER}/defdecoder.go matrixpilot.xml > ${OUTFOLDER}/matrixpilot.go
go run ${MAINFOLDER}/main.go ${MAINFOLDER}/defdecoder.go paparazzi.xml > ${OUTFOLDER}/paparazzi.go
go run ${MAINFOLDER}/main.go ${MAINFOLDER}/defdecoder.go python_array_test.xml > ${OUTFOLDER}/python_array_test.go
go run ${MAINFOLDER}/main.go ${MAINFOLDER}/defdecoder.go slugs.xml > ${OUTFOLDER}/slugs.go
go run ${MAINFOLDER}/main.go ${MAINFOLDER}/defdecoder.go standard.xml > ${OUTFOLDER}/standard.go
go run ${MAINFOLDER}/main.go ${MAINFOLDER}/defdecoder.go test.xml > ${OUTFOLDER}/test.go
go run ${MAINFOLDER}/main.go ${MAINFOLDER}/defdecoder.go ualberta.xml > ${OUTFOLDER}/ualberta.go
go run ${MAINFOLDER}/main.go ${MAINFOLDER}/defdecoder.go uAvionix.xml > ${OUTFOLDER}/uAvionix.go
go run ${MAINFOLDER}/main.go ${MAINFOLDER}/defdecoder.go ../../../message_definitions/v1.0/mavlink_freightfish.xml > ${OUTFOLDER}/mavlink_freightfish.go
popd
