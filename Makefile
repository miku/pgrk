sources = $(wildcard *.go)
targets = $(basename $(sources))

all: $(targets)

%: %.go
	gofmt -w -tabs=false -tabwidth=4 $<
	go build -o $@ $<

example.png:
	dot -Tpng example.dot -oexample.png

clean:
	rm -f $(targets)
	rm -f example.png
