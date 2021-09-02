ut:
	CGO_ENABLED=0 go test -v -coverprofile=coverage.txt $$(go list ./...)