test:
	go test ./... -tags testing -v

coverage:
	go test ./... -v -coverprofile=coverage.out