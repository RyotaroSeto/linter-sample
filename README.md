# linter-sample

## Install

- go install `github.com/quasilyte/go-ruleguard/cmd/ruleguard@latest`
- go install `github.com/quasilyte/go-ruleguard/dsl@latest`

## Lint start
### find HelloWorld
```zsh
% ruleguard -rules rules-helloworld.go -fix helloworld.go
linter-sample/helloworld.go:6:2: helloWorld: HelloWorld発見 (rules-helloworld.go:11)
```

### exploring named and unnamed pattern submatch variables
```zsh
% ruleguard -rules rules-submatch.go submatch.go
linter-sample/submatch.go:9:10: dupArg: suspicious duplicated arg "_" (rules-submatch.go:11)
linter-sample/submatch.go:11:10: dupArg: suspicious duplicated arg part (rules-submatch.go:11)
```

### putting several rules inside one function, making them short-circuit
```zsh
% ruleguard -rules rule-multi.go multi.go
linter-sample/multi.go:7:2: assignOp: can rewrite as v1 += v2 (rule-multi.go:19)
linter-sample/multi.go:8:2: assignOp: can rewrite as v1++ (rule-multi.go:18)
```

### putting several patterns into a single rule
```zsh
% ruleguard -rules rule-multi-pattern.go multi-pattern.go
linter-sample/multi-pattern.go:7:10: boolLiteralInExpr: omit bool literal in expression (rule-multi-pattern.go:12)
linter-sample/multi-pattern.go:8:10: boolLiteralInExpr: omit bool literal in expression (rule-multi-pattern.go:14)
linter-sample/multi-pattern.go:9:10: boolLiteralInExpr: omit bool literal in expression (rule-multi-pattern.go:13)
linter-sample/multi-pattern.go:10:10: boolLiteralInExpr: omit bool literal in expression (rule-multi-pattern.go:15)
```

## Use golangci-lint
```zsh
% golangci-lint run ./...
```

## Ruleguard by example
- https://go-ruleguard.github.io/by-example/
