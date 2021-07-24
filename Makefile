LNSC=lnsc

help:
	@echo make build
	@echo make test
	@echo make test-all
	@echo make test-trans
	@echo make access-token

build:
	find -iname '*.lns' | $(LNSC) @- save -langGo --main main --package main
	$(LNSC) main.lns mkmain entry.go
	go build
test: build
	./trasorg test.org -c ../trans-org-accesstoken.json -v
	./trasorg -m mkreq test.org -c ../trans-org-accesstoken.json
#	./trasorg -m trans test.org -c ../trans-org-accesstoken.json

test-all: build
	./trasorg ../ifritJP.github.io/hugo/content/LuneScript/all.org \
			-c ../trans-org-accesstoken.json > dump
	./trasorg -m mkreq ../ifritJP.github.io/hugo/content/LuneScript/all.org \
			-c ../trans-org-accesstoken.json > req.json


test-trans:
	cat testreq.json | \
		curl -X POST -H "Authorization: Bearer "$$(gcloud auth application-default print-access-token) \
			-H "Content-Type: application/json; charset=utf-8" \
			-d @- https://translation.googleapis.com/language/translate/v2

access-token:
	gcloud auth application-default print-access-token


