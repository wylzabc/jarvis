.PHONY: jarvis test

GOBIN = $(shell pwd)/build/bin

jarvis:
	$(shell pwd)/build/env.sh go build -o $(GOBIN)/jarvis
	@echo "Done building."
	@echo "Run \"$(GOBIN)/jarvis\" to launch."

test: testadd

testadd:
	$(shell pwd)/build/env.sh go test github.com/wylzabc/jarvis/add -v -cover
	@echo "Done test add."

