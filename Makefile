.PHONY: test
test:
	go test -v -race -tags safe ./... -covermode=atomic -coverprofile=coverage.out

	$(GOPATH)/bin/goveralls -coverprofile=coverage.out -service=travis-ci -repotoken $(COVERALLS_TOKEN)

.DEFAULT_GOAL := test