.DEFAULT_GOAL := build

.PHONY:fmt vet build

fmt:
	go fmt randomdata.go
	go fmt main.go

vet: fmt
	go vet main.go randomdata.go

build: vet
	go build -o randomdata main.go randomdata.go
#	go build -o randomdata -v -x main.go randomdata.go

clean:
	/bin/rm -f randomdata

