all: test lint build run

test:
	go test -v -cover -race ./...

lint:
	command -v golangci-lint || (cd /usr/local ; wget -O - -q https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s latest)
	golangci-lint --version
	golangci-lint run --disable-all \
	--deadline=10m \
	--skip-dirs vendor \
	--skip-files \.*_mock\.*\.go \
	-E errcheck \
	-E govet \
	-E unused \
	-E gocyclo \
	-E golint \
	-E varcheck \
	-E structcheck \
	-E maligned \
	-E ineffassign \
	-E interfacer \
	-E unconvert \
	-E goconst \
	-E gosimple \
	-E staticcheck \
	-E gosec

build:
	GO111MODULES=on CGO_ENABLED=0 go build -o mango


