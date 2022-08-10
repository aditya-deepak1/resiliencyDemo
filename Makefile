GOPATH=$(shell go env GOPATH)
PROJECT_ROOT=.

polish: tidy fmt vet race staticcheck

vet:
	go vet $(PROJECT_ROOT)/...

staticcheck:
	go install honnef.co/go/tools/cmd/staticcheck@latest
	$(GOPATH)/bin/staticcheck $(PROJECT_ROOT)/...

fmt:
	go fmt $(PROJECT_ROOT)/...

tidy:
	go mod tidy -v

race:
	go build -race main.go

setup:
	export PATH="$(GOPATH)/bin:$PATH"

	[[ -e `which brew` ]] || /bin/bash -c "curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh"
	[[ -e `which pre-commit` ]] || brew install pre-commit

	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/fzipp/gocyclo/cmd/gocyclo@latest

	pre-commit install
