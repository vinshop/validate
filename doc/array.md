# Array

Validate functions for an Array.

## Syntax

```go
Array(...ArrayFn)
```

## Array function

### Validate each element

Verify for each element in the array.

```go
Each(...Rule)
```

### Validate for the array itself

Verify for the array (like MinSize, MaxSize, etc...).

```go
ArrayHas(...Rule)
```

### Require minimum size

```go
MinSize(l int)
```

### Cap Limit

```go
MaxSize(l int)
```