# generate-mocks regenerates all the mocks for unit tests
generate-mocks:
	rm -rf ./mocks
	export PATH=$$PATH:$$(go env GOPATH)/bin && \
	cd helper/mockgeneration && go generate ./...