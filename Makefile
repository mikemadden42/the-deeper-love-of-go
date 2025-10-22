# Get all package directories, excluding vendor. Used for file-based tools.
GO_DIRS := $(shell go list -f '{{.Dir}}' ./... | grep -v /vendor/)

# Get all packages, excluding vendor. Used for package-based tools.
GO_PKGS := $(shell go list ./... | grep -v /vendor/)

.PHONY: check
check:
	@echo "--- Running Formatter (goimports) ---"
	@# Check formatting. test -z exits 0 if string is empty (no diffs)
	@test -z "$$(goimports -d $(GO_DIRS))" || \
		(echo "FAIL: goimports found files to format. Run 'make fmt' to fix:"; goimports -d $(GO_DIRS); exit 1)
	
	@echo "\n--- Running Linters (golangci-lint) ---"
	@# Run go vet (included) and other linters on all packages. 'run' defaults to './...'
	golangci-lint run
	
	@echo "\n--- Running Security Scanner (gosec) ---"
	@# Run gosec on all packages
	gosec $(GO_PKGS)
	
	@echo "\nAll checks passed."

# Add a simple 'format' target for convenience
.PHONY: fmt
fmt:
	@echo "Running goimports to format code..."
	@goimports -w $(GO_DIRS)
	@echo "Formatting complete."


