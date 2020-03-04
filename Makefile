all: clean release

clean:
	rm -rf cognitools
	rm -rf pkged.go

release:
	`go env GOPATH`/bin/pkger -include /webui/dist:./webui/dist
	go build -o cognitools
