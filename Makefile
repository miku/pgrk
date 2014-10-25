targets = pgrk pgrk-dot pgrk-gen 

all: $(targets)

fmt:
	goimports -w cmd/pgrk/pgrk.go
	goimports -w cmd/pgrk-dot/pgrk-dot.go
	goimports -w cmd/pgrk-gen/pgrk-gen.go

pgrk:
	go build -o pgrk cmd/pgrk/pgrk.go

pgrk-dot:
	go build -o pgrk-dot cmd/pgrk-dot/pgrk-dot.go

pgrk-gen:
	go build -o pgrk-gen cmd/pgrk-gen/pgrk-gen.go

example.png:
	dot -Tpng example.dot -o example.png

clean:
	rm -f $(targets)
	rm -f example.png
