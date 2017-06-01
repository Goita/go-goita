# Benchmark go-goita/search.Solve()

## Benchmark condition


Hardware | Info
---------|----------
 System | Linux Mint 18.1 Serena (64 bit gcc: 5.4.0)
 CPU | Quad core Intel Core i5-3570T (-MCP-) cache: 6144 KB clock speeds: min/max: 1600/3300 MHz
Memory | 15.6 GiB

bench command
```sh
go test ./search -bench Solve -benchmem -benchtime 10s 
```

profile command
```sh
go test ./search -bench Solve -benchmem -cpuprofile cpu.prof
go tool pprof cpu.prof
go test ./search -bench Solve -benchmem -memprofile mem.prof
go tool pprof --alloc_space mem.prof
```

```go
// Input history-1
"11244556,12234569,11123378,11113457,s3,371,411,115,2p,3p,4p,145,252,3p,4p,124,2p"
```


## Results of history-1
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
 8 | 50	 330169618 ns/op	  49139967 B/op	 1021554 allocs/op | use pre-allocated memory
 9 | 1000 20479849 ns/op       5162924 B/op   105370 allocs/op | alpha-beta negamax search
10 | 1000 22639590 ns/op       8282957 B/op   209756 allocs/op | bugfix: solve result history
11 | 1000 19782379 ns/op      19955525 B/op   191844 allocs/op | manage number of running goroutine
12 | 1000 18128649 ns/op      21660504 B/op   209012 allocs/op | efficient Hand implement

## Results of history-2

```go
// Input history-2
"12222447,11134568,11134557,11133569,s1"
```

rev | result | comment
---------|----------|---------
 0 |  | from history-1.rev0 


## Cut-Off effectiveness with Ordered moves

**Ordered moves: ON**
go run main.go solve -h 11244556,12234569,11123378,11113457,s3,313
&{VisitedNode:3344852555 VisitedLeaf:1359802779 CutOffedNode:2217395681 Routines:10 MaxRoutines:4}
[[35:-20] [34:-20] [37:-30] [31:-20] [p:-10]]
move:[35] score:[-20] 435,151,2p,313,4p,1p,235,3p,4p,1p,262,3p,4p,125,2p,3p,4p,146,284,381,414,1p,2p,3p,417,1p,2p,372

move:[34] score:[-20] 434,141,2p,313,4p,1p,234,3p,4p,1p,262,3p,4p,124,2p,3p,4p,156,285,381,415,1p,2p,3p,417,1p,2p,372

move:[37] score:[-30] 437,1p,2p,373,4p,1p,234,3p,4p,1p,226,3p,4p,164,2p,3p,4p,114,2p,3p,4p,115,2p,3p,4p,125

move:[31] score:[-20] 431,1p,2p,313,4p,1p,234,3p,4p,1p,225,3p,4p,154,2p,3p,447,1p,2p,3p,415,1p,2p,387,4p,1p,2p,312

move:[p] score:[-10] 4p,1p,232,3p,4p,121,2p,313,434,1p,2p,3p,417,1p,2p,3p,411,114,2p,3p,4p,154,245,387,4p,1p,282,321

execution time: 19m23.810449732s

**Ordered moves: OFF**
go run main.go solve -h 11244556,12234569,11123378,11113457,s3,313
search begin on 5 moves[p 31 34 35 37]
&{VisitedNode:3218663571 VisitedLeaf:1312547654 CutOffedNode:2133737818 Routines:4 MaxRoutines:4}
[[34:-20] [35:-20] [37:-30] [31:-20] [p:-10]]
move:[34] score:[-20] 434,141,2p,313,4p,1p,234,3p,4p,1p,262,3p,4p,124,2p,3p,4p,156,285,381,415,1p,2p,3p,417,1p,2p,372

move:[35] score:[-20] 435,151,2p,313,4p,1p,235,3p,4p,1p,262,3p,4p,125,2p,3p,4p,146,284,381,414,1p,2p,3p,417,1p,2p,372

move:[37] score:[-30] 437,1p,2p,373,4p,1p,234,3p,4p,1p,226,3p,4p,164,2p,3p,4p,114,2p,3p,4p,115,2p,3p,4p,125

move:[31] score:[-20] 431,1p,2p,313,4p,1p,234,3p,4p,1p,225,3p,4p,154,2p,3p,447,1p,2p,3p,415,1p,2p,387,4p,1p,2p,312

move:[p] score:[-10] 4p,1p,232,3p,4p,121,2p,313,434,1p,2p,3p,417,1p,2p,3p,411,114,2p,3p,4p,154,245,387,4p,1p,282,321

execution time: 13m16.19928656s
