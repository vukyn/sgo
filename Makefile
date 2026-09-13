#!make
# ⚠️ `-include`, not `include`, and the sed is guarded. .env is gitignored, so a
# fresh clone does not have one — with the bare forms every target failed with
# "No such file or directory. Stop." before reaching a recipe, which made the
# Makefile unusable anywhere but a developer's own checkout. Only `tag` actually
# reads a value out of .env, and it already errors on its own when VERSION is
# unset.
-include ./.env
export $(shell [ -f ./.env ] && sed 's/=.*//' ./.env)

PRJ=

build:
	@echo "Building $(PRJ)..."
	@go build -o bin/ ./$(PRJ)
	@echo "Build complete"

install:
	@echo "Installing $(PRJ)..."
	@go install ./$(PRJ)
	@echo "Install complete"

tag:
	git tag -a v$(VERSION) -m "Release version $(VERSION)"
	git push origin v$(VERSION)