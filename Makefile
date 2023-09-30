BINSERVICE := "./build/sysmon"


hello:
	echo "Hello"

build:
	go build -v -o $(BINSERVICE)  ./cmd


rungo: build
	$(BINSERVICE)

.PHONY: rungo

