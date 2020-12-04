setup:
	go install github.com/ozanh/ugo/cmd/ugo
	go install github.com/d5/tengo/v2/cmd/tengo
	go install github.com/d5/tengobench/cmd/gofib
	go install github.com/d5/tengobench/cmd/gofibtc
	go install github.com/yuin/gopher-lua/cmd/glua
	go install github.com/robertkrimen/otto/otto
	go install github.com/d5/tengobench/cmd/go-lua
	go install github.com/mattn/anko
	go install go.starlark.net/cmd/starlark
	go install github.com/go-python/gpython
	go install github.com/dop251/goja/goja
	go install github.com/traefik/yaegi/cmd/yaegi

start:
	@go run ./cmd/bench ./code

get:
	go get -u github.com/ozanh/ugo
	go get -u github.com/d5/tengo/v2
	go get -u github.com/Shopify/go-lua
	go get -u github.com/traefik/yaegi
	go get -u github.com/dop251/goja
	go get -u github.com/go-python/gpython
	go get -u github.com/mattn/anko
	go get -u github.com/robertkrimen/otto
	go get -u github.com/yuin/gopher-lua
	go get -u go.starlark.net

fmt:
	@go fmt ./...
