
sync:
	@echo "sync figma locals"

build:
	web-ext build --ignore-files *.js dictionary.json *.go go.*
