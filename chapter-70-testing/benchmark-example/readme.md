# what I learned for this chapter
-Benchmark
-Example
-Test

```
BenchmarkJointextManualWay(b *testing.B)
BenchmarkJoinTextFastWay(b *testing.B)
TestJointextManualWay(t *testing.T)
TestJointextFastWay(t *testing.T)
```

# Command

```
go test
go test -bench .
go test -cover
go test -coverprofile c.out
go tool cover -html=c.out 
```
