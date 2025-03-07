MOCK_DIR=./internal/test/mocks
DAO_DIR=./internal/dao

DAO_FILES=$(wildcard $(DAO_DIR)/*.go)
MOCK_FILES=$(patsubst $(DAO_DIR)/%.go, $(MOCK_DIR)/%_mock.go, $(DAO_FILES))

$(MOCK_DIR)/%_mock.go: $(DAO_DIR)/%.go
	@mkdir -p $(MOCK_DIR)
	mockgen -source=$< -destination=$@ -package=tests
	@echo "Generated mock: $@"

mocks: $(MOCK_FILES)

clean:
	rm -rf $(MOCK_DIR)