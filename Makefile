
EXCLUDE_DIR = cmd|models|tools

# This is bmake (be aware)
TEST_FILES != go list ./... | grep -E -v "$(EXCLUDE_DIR)"

.PHONY: test-list
test-list:
	@echo $(TEST_FILES)

.PHONY: test
test:
	go test $(TEST_FILES)

.PHONY: coverage
coverage:
	go test -coverprofile=cover.out $(TEST_FILES)
	go tool cover -html=cover.out

.PHONY: run-app
run-app: 
	DB_PATH=test.db go run cmd/main.go

.include "build/integration-test.mk"
