.DEFAULT_GOAL := build

.PHONY:fmt vet build

fmt:
	go fmt randomdata.go

vet: fmt
	go vet main.go randomdata.go

build: vet
	go build -o randomdata main.go randomdata.go

clean: randomdata
	/bin/rm randomdata

