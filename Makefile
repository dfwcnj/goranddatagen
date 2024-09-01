.DEFAULT_GOAL := build

.PHONY:fmt vet build

fmt:
	go fmt randomdata/*.go
	go fmt *.go

vet: fmt
	go vet ramdomdata/randomdata.go
	go vet *.go

build: vet
	go build -o randomdata *.go randomdata/randomdata.go

clean:
	/bin/rm -f randomdata

