VERSION:=0.0.1

keymold:
	go build -ldflags "-w -s -X $$(head -n1 go.mod | awk '{print $$2}')/cli.Version=$(VERSION)"

.PHONY: clean
clean:
	rm -rf keymold
