targets = pgrk pgrk-dot pgrk-gen 

all: $(targets)

fmt:
	goimports -w cmd/pgrk/main.go
	goimports -w cmd/pgrk-dot/main.go
	goimports -w cmd/pgrk-gen/main.go

pgrk:
	go build -o pgrk cmd/pgrk/main.go

pgrk-dot:
	go build -o pgrk-dot cmd/pgrk-dot/main.go

pgrk-gen:
	go build -o pgrk-gen cmd/pgrk-gen/main.go

example.png:
	dot -Tpng example.dot -o example.png

clean:
	rm -f $(targets)
	rm -f example.png
