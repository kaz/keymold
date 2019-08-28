VERSION:=0.1.0

keymold:
	go build -ldflags "-w -s -X $$(head -n1 go.mod | awk '{print $$2}')/cli.Version=$(VERSION)"

.PHONY: install
install: keymold
	install -m 0755 $< /usr/local/bin

.PHONY: package
package: keymold
	tar zcvf keymold-$(VERSION)-x86_64.darwin.tar.gz $<

.PHONY: clean
clean:
	rm -rf keymold *.tar.gz
