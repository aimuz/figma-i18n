
sync:
	@echo "sync figma locals"

build:
	go run sync.go
	web-ext build --ignore-files *.js dictionary.json *.go go.* --overwrite-dest
