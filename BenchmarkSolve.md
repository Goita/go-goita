# Benchmark go-goita/search.Solve()

## Benchmark condition


Hardware | Info
---------|----------
 System | Linux Mint 18.1 Serena (64 bit gcc: 5.4.0)
 CPU | Quad core Intel Core i5-3570T (-MCP-) cache: 6144 KB clock speeds: min/max: 1600/3300 MHz
Memory | 15.6 GiB


```go
// Input history
"11244556,12234569,11123378,11113457,s3,371,411,115,2p,3p,4p,145,252,3p,4p,124,2p"
```

bench command
```sh
go test ./search -bench Solve -benchmem -benchtime 10s 
```

profile command
```sh
go test ./search -bench Solve -benchmem -cpuprofile cpu.prof
go test ./search -bench Solve -benchmem -memprofile mem.prof
```


## Results
rev | result | comment
---------|----------|---------
 0 | 5	2614808707 ns/op	1710808150 B/op	43550518 allocs/op | converted from js
 1 | 5	2655975707 ns/op	1696513310 B/op	40685448 allocs/op | hash mapping move
 2 | 10	1328204030 ns/op	1652589372 B/op	 5496297 allocs/op | improve koma.String() 
 3 | 10	1127504252 ns/op	2658544395 B/op	 5496547 allocs/op | MoveHashArray
 4 | 20  922853893 ns/op    2639373691 B/op  5460777 allocs/op | improve GetUnique()
 5 | 20  767939206 ns/op    1818866038 B/op  5460578 allocs/op | minimum memory alloc in GetPossibleMoves()
 6 | 20  612600344 ns/op    1139440927 B/op  5504890 allocs/op | reduce memory alloc block-size in GetPossibleMoves()
 7 | 30	 398722771 ns/op	 125237178 B/op	 5504485 allocs/op | no defer
