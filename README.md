# SOA - Structure of Arrays

[![Go Report Card](https://goreportcard.com/badge/github.com/andygeiss/soa)](https://goreportcard.com/report/github.com/andygeiss/soa)

Allocate, Grow, Pack or Unpack Golang slices of basic types to optimize memory access for Data-Driven design.
The primary motivation is easier manipulation with packed SIMD instructions like SSE2.

![](soa.png)

##### Table of Contents

- [Features](README.md#features)
- [Installation](README.md#installation)
- [Benchmark](README.md#benchmark)
- [Usage](README.md#usage)
    - [Allocate](README.md#allocate)
    - [Grow](README.md#grow)
    - [Pack](README.md#pack)
    - [Unpack](README.md#unpack)
    - [Add](README.md#add)

## Freatures

* **Allocate** a new slice of a given basic type which uses one page of memory.
* **Grow** a slice of a given basic type to the next multiple of a page size.
* **Pack** all non-zero (0) values of a slice without changing the order.
* **Unpack** the values of a slice to the next multiple of a page.
* **Add** two slices using SIMD instructions.

## Installation

    go get -u github.com/andygeiss/soa

## Benchmark

According to the following benchmark, the performance gain is between factor 3 for 64-bit and factor 6 for 32-bit by using SSE2.
Theoretically, AVX would be able to double this increase:

    go test -benchmem -run=^$ github.com/andygeiss/soa -bench ^(BenchmarkAddFloat32s|BenchmarkAddFloat64s|BenchmarkAddInt32s|BenchmarkAddInt64s)$

    goos: linux
    goarch: amd64
    pkg: github.com/andygeiss/soa
    BenchmarkAddFloat32s/Native-4         	 1806871	       614 ns/op	       0 B/op	       0 allocs/op
    BenchmarkAddFloat32s/Assembly_SSE-4   	10943130	       113 ns/op	       0 B/op	       0 allocs/op
    BenchmarkAddFloat64s/Native-4         	 3843424	       310 ns/op	       0 B/op	       0 allocs/op
    BenchmarkAddFloat64s/Assembly_SSE-4   	10939322	       109 ns/op	       0 B/op	       0 allocs/op
    BenchmarkAddInt32s/Native-4           	 1790896	       669 ns/op	       0 B/op	       0 allocs/op
    BenchmarkAddInt32s/Assembly_SSE-4     	 5618310	       217 ns/op	       0 B/op	       0 allocs/op
    BenchmarkAddInt64s/Native-4           	 1948944	       612 ns/op	       0 B/op	       0 allocs/op
    BenchmarkAddInt64s/Assembly_SSE-4     	11035382	       110 ns/op	       0 B/op	       0 allocs/op
    Benchmark/Alloc-4                     	 1775888	       680 ns/op	    4160 B/op	       3 allocs/op
    Benchmark/Grow-4                      	  583726	      1980 ns/op	   14464 B/op	       6 allocs/op
    Benchmark/Pack-4                      	  105363	     11345 ns/op	      64 B/op	       2 allocs/op
    Benchmark/Unpack-4                    	  732098	      1505 ns/op	    8320 B/op	       6 allocs/op

## Usage

#### Allocate
The following function creates new entities with a position (p) and velocity (v),
which fits into one page size (ex. 4096 bytes on linux).

```go
func createEntities() (p, v []int32) {
    p = soa.DefaultManager.Allocate([]int32{}).([]int32)
    v = soa.DefaultManager.Allocate([]int32{}).([]int32)
    return
}
```

#### Grow

Next we want to ensure that we have enough space to add a new entity to the world, by using the following code. 

```go
func ensureSpace(offset, p, v) (pn []int32, vn []int32) {
    // no more space left for a new entity? then grow ... 
    if offset + 1 >= len(p) {
       pn = soa.DefaultManager.Grow(p)
       vn = soa.DefaultManager.Grow(v)
    }
    return p, v
}
```

#### Pack

For serialization we dont need to save unused / initialized zero values.
Thus, we pack the slices.

```go
type World struct {
    p []int32
    v []int32
}

func saveWorld(w *World, filename string) (err error) {
    // open a new file
    file, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
    if err != nil {
	return err
    }
    defer file.Close()
    // create packed world data
    var tmp World
    tmp.p = soa.DefaultManager.Pack(w.p).([]int32)
    tmp.v = soa.DefaultManager.Pack(w.v).([]int32)
    // encode data to JSON
    if err := json.NewEncoder(file).Encode(&tmp); err != nil {
	return err
    }
    return nil
}
```

#### Unpack

Next time we load the world from the file and unpack the data:

```go
func loadWorld(filename string) (w *World, err error) {
    // open a new file
    file, err := os.Open(filename)
    if err != nil {
	return nil, err
    }
    defer file.Close()
    // decode data from JSON
    var tmp World
    if err := json.NewDecoder(file).Decode(&tmp); err != nil {
	return nil, err
    }
    tmp.p = soa.DefaultManager.Unpack(tmp.p).([]int32)
    tmp.v = soa.DefaultManager.Unpack(tmp.v).([]int32)
    return &tmp, nil
}
```

#### Add

Two slices of float64s which are equal in length (multiple of page size) could be easily added by using SSE2 (as an example).

```go
func moveEntities(pos, velocity float64[]) {
    // add the position and velocity vectors together and overwrite the old positions.
    soa.AddFloat64s(pos, velocity, pos)
}
```
