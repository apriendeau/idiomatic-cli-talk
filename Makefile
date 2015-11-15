default: build

build:
	node-sass --output-style compressed style.scss style.css
	cleaver --style style.css ./presentation.md

.PHONY: build
