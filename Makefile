SRC_MAIN = cmd/main.go
SRC_AUTMCS = cmd/authorization/auth.go

server: go-run
autmcs: go-run-autmcs

test: go-test go-tool
	cat cover | grep -v "mock" | grep -v  "easyjson" | grep -v "proto" > cover.out

clean:
	rm -f /tmp/session_*

.PHONY: server test clean

go-run: ${SRC_MAIN}
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go run ${SRC_MAIN}

go-run-autmcs: ${SRC_AUTMCS}
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go run ${SRC_AUTMCS}
	
go-test: 
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go test -coverpkg=./... -coverprofile=cover ./...

go-tool:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go tool cover -func=cover.out
set-gopath:
	export GOPATH=$HOME/go && export PATH=$PATH:$GOROOT/bin:$GOPATH/bin

create-easyjson:
	easyjson -all internal/pkg/domain/*.go
easyjson:
	rm internal/pkg/domain/*easyjson.go && easyjson -all internal/pkg/domain/*.go
# Команда для запуска тестов: go test -coverpkg=./... -coverprofile=cover ./... && cat cover | grep -v "mock" | grep -v  "easyjson" | grep -v "proto" > cover.out && go tool cover -func=cover.out
