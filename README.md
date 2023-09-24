# go_study
Amazing Go study

## Core package
- echo
- sqlc

```
sqlc generate
```

### 테스트 

```
go test -coverprofile=coverage.out

go tool cover -html=coverage.out

go test -coverprofile cover.prof

go tool cover -html=cover.prof -o cover.html
```