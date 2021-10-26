BINDIR ?= $(GOPATH)/bin


ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=game \
	  -X github.com/cosmos/cosmos-sdk/version.AppName=gamed \
	  -X github.com/cosmos/cosmos-sdk/version.Version=0.0.1 
	  
ldflags += $(LDFLAGS)
ldflags := $(strip $(ldflags))

install: go.sum
	go install -mod=readonly -ldflags '$(ldflags)' ./cmd/gamed

go.sum: go.mod
	@go mod verify
	@go mod tidy