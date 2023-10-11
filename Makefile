clean:
	rm -f state.json history.json schema_* *.jsonl
	go mod tidy
	go mod vendor
	go build .
