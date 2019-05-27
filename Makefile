.PHONY: jarvis test

GOBIN = $(shell pwd)/build/bin

jarvis:
	$(shell pwd)/build/env.sh go build -o $(GOBIN)/jarvis
	@echo "Done building."
	@echo "Run \"$(GOBIN)/jarvis\" to launch."

test: testadd testsub testmulti

testadd:
	$(shell pwd)/build/env.sh go test github.com/wylzabc/jarvis/add -v -cover
	@echo "Done test add."

testsub:
	$(shell pwd)/build/env.sh go test github.com/wylzabc/jarvis/subtraction -v -cover
	@echo "Done test sub."

testmulti:
	$(shell pwd)/build/env.sh go test github.com/wylzabc/jarvis/multiplication -v -cover
	@echo "Done test multi."

