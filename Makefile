# This version-strategy uses git refs to set the version string
# Get the following from left to right: tag > branch > branch of detached HEAD commit
VERSION = $(shell (git describe --tags --exact-match 2>/dev/null || git symbolic-ref -q --short HEAD 2>/dev/null || git name-rev --name-only "$$( git rev-parse --short HEAD )" | sed 's@.*/@@') | tr '/' '-' | head -c10)

# Get the short SHA
SHA_SHORT = $(shell git rev-parse --short HEAD)

# Get the date of build
BUILD_DT:=$(shell date +%F-%T)

build:
	go build -ldflags "-s -w -extldflags \"-static\" -X $(go list -m)/pkg/version.COMMIT_SHA1=$(git rev-parse --short HEAD) -X $(go list -m)/pkg/version.BUILD_DATE=$(date +%F-%T)" ./cmd/docker_stats_exporter

