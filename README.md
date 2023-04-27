# Design Pattern in Go
### What is this?
This is my personal project to learn design pattern in Go. I will try to implement all design pattern in Go. I will also try to benchmark each pattern to see which one is the best. You can give me feedback if you think my implementation is wrong or you have better implementation.

### Benchmark result
Benchmark result for each pattern.

Im only using my laptop for benchmarking, so the result may not be accurate. You can see my laptop spec below.
```txt
OS: EndeavourOS Linux x86_64 
Host: Swift SF314-41 V1.10 
Kernel: 5.15.94-1-lts
CPU: AMD Ryzen 5 3500U with Radeon Vega Mobile Gfx (8) @ 2.100GHz 
Memory: 12GiB 
GO: go version go1.20.1 linux/amd64
```

#### Singleton pattern
SOON

#### Factory pattern
SOON

#### Builder pattern

command :
```bash
export FUNCTION=BenchmarkCarWithoutBuilder  &&  go test -benchmem -run=^$ -bench ^$FUNCTION$ sdp/builder -count 10 > "$FUNCTION.txt"   
```

```txt
goos: linux
goarch: amd64
pkg: sdp/builder
cpu: AMD Ryzen 5 3500U with Radeon Vega Mobile Gfx  
BenchmarkCarWithoutBuilder-8   	1000000000	         0.3093 ns/op	       0 B/op	       0 allocs/op
BenchmarkCarWithoutBuilder-8   	1000000000	         0.3334 ns/op	       0 B/op	       0 allocs/op
BenchmarkCarWithoutBuilder-8   	1000000000	         0.3245 ns/op	       0 B/op	       0 allocs/op
BenchmarkCarWithoutBuilder-8   	1000000000	         0.4177 ns/op	       0 B/op	       0 allocs/op
BenchmarkCarWithoutBuilder-8   	1000000000	         0.3192 ns/op	       0 B/op	       0 allocs/op
BenchmarkCarWithoutBuilder-8   	1000000000	         0.3203 ns/op	       0 B/op	       0 allocs/op
BenchmarkCarWithoutBuilder-8   	1000000000	         0.3244 ns/op	       0 B/op	       0 allocs/op
BenchmarkCarWithoutBuilder-8   	1000000000	         0.3697 ns/op	       0 B/op	       0 allocs/op
BenchmarkCarWithoutBuilder-8   	1000000000	         0.4043 ns/op	       0 B/op	       0 allocs/op
BenchmarkCarWithoutBuilder-8   	1000000000	         0.3064 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	sdp/builder	3.998s
```

command :
```bash
export FUNCTION=BenchmarkCarWithBuilder  &&  go test -benchmem -run=^$ -bench ^$FUNCTION$ sdp/builder -count 10 > "$FUNCTION.txt"   
```

```txt
goos: linux
goarch: amd64
pkg: sdp/builder
cpu: AMD Ryzen 5 3500U with Radeon Vega Mobile Gfx  
BenchmarkCarWithBuilder-8   	46566160	        23.65 ns/op	       0 B/op	       0 allocs/op
BenchmarkCarWithBuilder-8   	52180687	        26.55 ns/op	       0 B/op	       0 allocs/op
BenchmarkCarWithBuilder-8   	45210944	        25.83 ns/op	       0 B/op	       0 allocs/op
BenchmarkCarWithBuilder-8   	47239432	        26.71 ns/op	       0 B/op	       0 allocs/op
BenchmarkCarWithBuilder-8   	42893767	        24.94 ns/op	       0 B/op	       0 allocs/op
BenchmarkCarWithBuilder-8   	48100722	        24.36 ns/op	       0 B/op	       0 allocs/op
BenchmarkCarWithBuilder-8   	47479564	        25.28 ns/op	       0 B/op	       0 allocs/op
BenchmarkCarWithBuilder-8   	48045500	        27.48 ns/op	       0 B/op	       0 allocs/op
BenchmarkCarWithBuilder-8   	45597368	        26.65 ns/op	       0 B/op	       0 allocs/op
BenchmarkCarWithBuilder-8   	45343317	        26.15 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	sdp/builder	13.496s
```

command :
```bash
export FUNCTION=BenchmarkManyCarWithoutBuilder  &&  go test -benchmem -run=^$ -bench ^$FUNCTION$ sdp/builder -count 10 > "$FUNCTION.txt"   
```

```txt
goos: linux
goarch: amd64
pkg: sdp/builder
cpu: AMD Ryzen 5 3500U with Radeon Vega Mobile Gfx  
BenchmarkManyCarWithoutBuilder-8   	1000000000	         0.3319 ns/op	       0 B/op	       0 allocs/op
BenchmarkManyCarWithoutBuilder-8   	1000000000	         0.3169 ns/op	       0 B/op	       0 allocs/op
BenchmarkManyCarWithoutBuilder-8   	1000000000	         0.3900 ns/op	       0 B/op	       0 allocs/op
BenchmarkManyCarWithoutBuilder-8   	1000000000	         0.3307 ns/op	       0 B/op	       0 allocs/op
BenchmarkManyCarWithoutBuilder-8   	1000000000	         0.4076 ns/op	       0 B/op	       0 allocs/op
BenchmarkManyCarWithoutBuilder-8   	1000000000	         0.3329 ns/op	       0 B/op	       0 allocs/op
BenchmarkManyCarWithoutBuilder-8   	1000000000	         0.3573 ns/op	       0 B/op	       0 allocs/op
BenchmarkManyCarWithoutBuilder-8   	1000000000	         0.3387 ns/op	       0 B/op	       0 allocs/op
BenchmarkManyCarWithoutBuilder-8   	1000000000	         0.3631 ns/op	       0 B/op	       0 allocs/op
BenchmarkManyCarWithoutBuilder-8   	1000000000	         0.3221 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	sdp/builder	4.042s
```

command :
```bash
export FUNCTION=BenchmarkManyCarWithBuilder  &&  go test -benchmem -run=^$ -bench ^$FUNCTION$ sdp/builder -count 10 > "$FUNCTION.txt"   
```

```txt
goos: linux
goarch: amd64
pkg: sdp/builder
cpu: AMD Ryzen 5 3500U with Radeon Vega Mobile Gfx  
BenchmarkManyCarWithBuilder-8   	16300284	        88.28 ns/op	       0 B/op	       0 allocs/op
BenchmarkManyCarWithBuilder-8   	16317159	        75.00 ns/op	       0 B/op	       0 allocs/op
BenchmarkManyCarWithBuilder-8   	16777719	       103.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkManyCarWithBuilder-8   	16601606	        85.30 ns/op	       0 B/op	       0 allocs/op
BenchmarkManyCarWithBuilder-8   	 8810316	       335.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkManyCarWithBuilder-8   	 1830415	       644.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkManyCarWithBuilder-8   	 1922127	       626.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkManyCarWithBuilder-8   	16493972	        85.44 ns/op	       0 B/op	       0 allocs/op
BenchmarkManyCarWithBuilder-8   	15629968	        75.30 ns/op	       0 B/op	       0 allocs/op
BenchmarkManyCarWithBuilder-8   	13927608	        76.60 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	sdp/builder	19.016s
```

command :
```bash
export FUNCTION=BenchmarkManyCarWithoutBuilderWithAlloc  &&  go test -benchmem -run=^$ -bench ^$FUNCTION$ sdp/builder -count 10 > "$FUNCTION.txt"   
```

```txt
goos: linux
goarch: amd64
pkg: sdp/builder
cpu: AMD Ryzen 5 3500U with Radeon Vega Mobile Gfx  
BenchmarkManyCarWithoutBuilderWithAlloc-8   	  551187	      2052 ns/op	     676 B/op	       6 allocs/op
BenchmarkManyCarWithoutBuilderWithAlloc-8   	  474135	      2216 ns/op	     747 B/op	       6 allocs/op
BenchmarkManyCarWithoutBuilderWithAlloc-8   	  715899	      2590 ns/op	     953 B/op	       6 allocs/op
BenchmarkManyCarWithoutBuilderWithAlloc-8   	  812655	      2229 ns/op	     848 B/op	       6 allocs/op
BenchmarkManyCarWithoutBuilderWithAlloc-8   	  751066	      2425 ns/op	     911 B/op	       6 allocs/op
BenchmarkManyCarWithoutBuilderWithAlloc-8   	  774258	      2383 ns/op	     886 B/op	       6 allocs/op
BenchmarkManyCarWithoutBuilderWithAlloc-8   	  638835	      2658 ns/op	    1059 B/op	       6 allocs/op
BenchmarkManyCarWithoutBuilderWithAlloc-8   	  681963	      2585 ns/op	     996 B/op	       6 allocs/op
BenchmarkManyCarWithoutBuilderWithAlloc-8   	  832257	      2412 ns/op	     829 B/op	       6 allocs/op
BenchmarkManyCarWithoutBuilderWithAlloc-8   	  503432	      2141 ns/op	     716 B/op	       6 allocs/op
PASS
ok  	sdp/builder	16.339s
```

command :
```bash
export FUNCTION=BenchmarkManyCarWithBuilderWithAlloc  &&  go test -benchmem -run=^$ -bench ^$FUNCTION$ sdp/builder -count 10 > "$FUNCTION.txt"   
```

```txt
goos: linux
goarch: amd64
pkg: sdp/builder
cpu: AMD Ryzen 5 3500U with Radeon Vega Mobile Gfx  
BenchmarkManyCarWithBuilderWithAlloc-8   	  484632	      2141 ns/op	     735 B/op	       6 allocs/op
BenchmarkManyCarWithBuilderWithAlloc-8   	  492708	      2144 ns/op	     726 B/op	       6 allocs/op
BenchmarkManyCarWithBuilderWithAlloc-8   	  577790	      2430 ns/op	    1163 B/op	       6 allocs/op
BenchmarkManyCarWithBuilderWithAlloc-8   	  754068	      2208 ns/op	     908 B/op	       6 allocs/op
BenchmarkManyCarWithBuilderWithAlloc-8   	  580905	      2344 ns/op	    1157 B/op	       6 allocs/op
BenchmarkManyCarWithBuilderWithAlloc-8   	  752586	      2309 ns/op	     909 B/op	       6 allocs/op
BenchmarkManyCarWithBuilderWithAlloc-8   	  657351	      2641 ns/op	    1031 B/op	       6 allocs/op
BenchmarkManyCarWithBuilderWithAlloc-8   	  634466	      2670 ns/op	    1065 B/op	       6 allocs/op
BenchmarkManyCarWithBuilderWithAlloc-8   	  617529	      2509 ns/op	    1093 B/op	       6 allocs/op
BenchmarkManyCarWithBuilderWithAlloc-8   	  751978	      2371 ns/op	     910 B/op	       6 allocs/op
PASS
ok  	sdp/builder	17.259s
```