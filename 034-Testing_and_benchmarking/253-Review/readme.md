# Remember to BET
- Benchmark
- Example
- Test

```
BenchmarkCat(b *testing.B)
ExampleCat()
TestCat(t *testing.T)
```

# Commands

```
godoc -http=localhost:6060

go test
go test -bench .
go test -cover
go test -coverprofile c.out
go tool cover -html=c.out
```