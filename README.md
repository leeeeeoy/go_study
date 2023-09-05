# go_study
개쩌는 Go study

## Core package
- echo
- ent

### ent 명령어

- 새로운 schema 생성



```
go run -mod=mod entgo.io/ent/cmd/ent new Scheme

go generate ./ent
```

### 테스트 

```
go test -coverprofile=coverage.out

go tool cover -html=coverage.out

go test -coverprofile cover.prof

go tool cover -html=cover.prof -o cover.html
```