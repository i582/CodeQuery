.PHONY: compile build

build: compile
	go build

release:
	go run ./_script/release.go

compile: ./pkg/querylang/internal/query/query.go ./pkg/querylang/internal/query/scanner.go
	sed -i -e 's/yyErrorVerbose = false/yyErrorVerbose = true/g' ./pkg/querylang/internal/query/query.go

./pkg/querylang/internal/query/query.go: ./pkg/querylang/internal/query/query.y
	cp ./pkg/querylang/internal/query/query.y ./pkg/querylang/internal/query/query-p.y
	sed -i -e "s/\@b/yylex\.\(\*Parser\)\.builder/g" ./pkg/querylang/internal/query/query-p.y
	~/go/bin/goyacc -o $@ ./pkg/querylang/internal/query/query-p.y

./pkg/querylang/internal/query/scanner.go: ./pkg/querylang/internal/query/scanner.rl
	ragel -Z -G2 -o $@ $<
