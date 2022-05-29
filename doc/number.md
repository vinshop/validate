# Number

Validate functions for a Number

## Syntax

```go
Number(...Rule)
```

### Equal

```go
EQ(float64)
```

### Not equal

```go
NEQ(float64)
```

### Less than

```go
LT(float64)
```

### Less or equal

```go
LTE(float64)
```

### Greater than

```go
GT(float64)
```

### Greater or equal

```go
GTE(float64)
```

### Make a change to the number

```go
DoMath(func (float64) float64, ...Rule)
```

### Custom validator

```go
CustomNumber(func (n float64) error {
	// Your logic goes here
})
```