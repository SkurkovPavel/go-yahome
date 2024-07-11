.PHONY: test
test:
	go test -v -tags safe ./... -covermode=atomic -coverprofile=coverage.out
	$(GOPATH)/bin/goveralls -coverprofile=coverage.out -repotoken $(COVERALLS_TOKEN)

.DEFAULT_GOAL := test