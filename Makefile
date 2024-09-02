.DEFAULT_GOAL := build

.PHONY:fmt vet build

fmt:
	go fmt *.go

vet: fmt
	go vet *.go

build: vet
	go build -o randomdata main.go randomdata/randomdata.go

clean:
	/bin/rm -f randomdata

