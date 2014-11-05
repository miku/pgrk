TARGETS = pgrk pgrk-dot pgrk-gen

all: $(TARGETS)

fmt:
	goimports -w cmd/pgrk/pgrk.go
	goimports -w cmd/pgrk-dot/pgrk-dot.go
	goimports -w cmd/pgrk-gen/pgrk-gen.go

deps:
	go get ./...

pgrk: deps
	go build -o pgrk cmd/pgrk/pgrk.go

pgrk-dot: deps
	go build -o pgrk-dot cmd/pgrk-dot/pgrk-dot.go

pgrk-gen: deps
	go build -o pgrk-gen cmd/pgrk-gen/pgrk-gen.go

example.png:
	dot -Tpng example.dot -o example.png

clean:
	rm -f $(TARGETS)
	rm -f example.png
	rm -f *.rpm
	rm -f debian/*.deb

deb: $(TARGETS)
	mkdir -p debian/pgrk/usr/sbin
	cp $(TARGETS) debian/pgrk/usr/sbin
	cd debian && fakeroot dpkg-deb --build pgrk .

rpm: $(TARGETS)
	mkdir -p $(HOME)/rpmbuild/{BUILD,SOURCES,SPECS,RPMS}
	cp ./packaging/pgrk.spec $(HOME)/rpmbuild/SPECS
	cp $(TARGETS) $(HOME)/rpmbuild/BUILD
	./packaging/buildrpm.sh pgrk
	cp $(HOME)/rpmbuild/RPMS/x86_64/pgrk*rpm .
