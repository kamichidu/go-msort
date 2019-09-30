.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: generate
generate: deps  ## go generate
	go generate ./...

.PHONY: deps
deps:  ## install go generate dependencies
	go get github.com/kamichidu/gobst
