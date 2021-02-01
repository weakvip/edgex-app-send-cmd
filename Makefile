.PHONY: build clean

GO=CGO_ENABLED=1 GO111MODULE=on go

app-name=app-service-demo

build:
	$(GO) build -o $(app-name)

clean:
	rm -f $(app-name)
