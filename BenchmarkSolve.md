# Benchmark go-goita/search.Solve()

## Benchmark condition

```go
// Input history
"11244556,12234569,11123378,11113457,s3,371,411,115,2p,3p,4p,145,252,3p,4p,124,2p"
```

bench command
```sh
cd search
go test -bench Solve -benchmem -benchtime 10s
```

## Results
rev | result | comment
---------|----------|---------
 0 | 5	2614808707 ns/op	1710808150 B/op	43550518 allocs/op | converted from js
 1 | 5	2655975707 ns/op	1696513310 B/op	40685448 allocs/op | hash mapping move
 A3 | B3 | C3
