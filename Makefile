# Windows
all: clean build copy

clean:
	del /Q build

.PHONY: build
build:
	set CGO_ENABLED=1 && go build -a -ldflags="-X 'main.StatusMode=release' -s -w" -o=build/BukuTamu.exe main.go
	go build -a -ldflags="-s -w" -o=build/gblib.exe gblib\gblib.go

copy:
	copy binary build
