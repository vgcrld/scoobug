# Generics in Go

This sample demonstrates the use of generics in Go. The `generics.go` file contains functions that sum the values of maps with different types of values (int64 and float64) using both non-generic and generic approaches.

## Functions

### `SumInts`

```go
func SumInts(m map[string]int64) int64
```

This function takes a map with string keys and int64 values and returns the sum of the values.

### `SumFloats`

```go
func SumFloats(m map[string]float64) float64
```

This function takes a map with string keys and float64 values and returns the sum of the values.

### `SumIntsOrFloats`

```go
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V
```

This generic function takes a map with keys of any comparable type and values of either int64 or float64, and returns the sum of the values.

## Usage

The `main` function initializes two maps, one with int64 values and one with float64 values, and demonstrates the use of both non-generic and generic sum functions.

```go
func main() {
	// ...existing code...
}
```

The output will show the sums calculated by both non-generic and generic functions.
