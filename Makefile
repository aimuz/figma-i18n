
sync:
	@echo "sync figma locals"

dist:
dist:
	mkdir dist

build: dist
	go run sync.go
	@cp manifest.json dist/
	@cp LICENSE dist/
	@cp README.md dist/
	@cp -r img dist/
	@cp -r js dist/
	@cp -r locales dist/
	web-ext build -s dist/ --i **/locale.json -o
