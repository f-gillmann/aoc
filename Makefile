YEAR 		:= 2025
BIN_DIR		:= ./bin
SRC_DIR		:= ./aoc

CMD_DIR 	= $(SRC_DIR)/$(YEAR)/day$(day)
BINARY  	= aoc-$(YEAR)-$(day)

.PHONY: all build run clean

all: build

build:
ifndef day
	$(error Day is undefined. Usage: make run day=01)
endif
	@echo "Building Year: $(YEAR), Day: $(day)..."
	mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(BINARY) $(CMD_DIR)
	chmod +x $(BIN_DIR)/$(BINARY)

run: build
	@echo "Running Year: $(YEAR), Day: $(day)..."
	$(BIN_DIR)/$(BINARY)

clean:
	rm -rf $(BIN_DIR)