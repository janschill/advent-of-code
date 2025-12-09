.PHONY: create-day test test-all run get-input

DAY ?= $(shell date +%-d)

COMMAND_TARGETS := create-day test test-all run get-input

ifeq ($(filter $(COMMAND_TARGETS),$(firstword $(MAKECMDGOALS))),$(firstword $(MAKECMDGOALS)))
  DAY_ARG := $(word 2,$(MAKECMDGOALS))
  ifneq ($(DAY_ARG),)
    override DAY := $(DAY_ARG)
    $(eval $(DAY_ARG):;@:)
  endif
endif

PADDED_DAY = $(shell printf "%02d" $(DAY))

create-day:
	@mkdir -p puzzles/day$(PADDED_DAY)
	@cp -r ./template/. puzzles/day$(PADDED_DAY)/
	@$(MAKE) get-input
	@echo "Day $(DAY) created"

test:
	@go test ./puzzles/day$(PADDED_DAY)

test-all:
	@go test ./...

run:
	@go run ./puzzles/day$(PADDED_DAY)/main.go

get-input:
	@mkdir -p ./puzzles/day$(PADDED_DAY)
	@curl --silent --header @cookie.txt --output ./puzzles/day$(PADDED_DAY)/input.txt https://adventofcode.com/2025/day/$(DAY)/input
