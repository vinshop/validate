# Struct
Validate functions for a Struct
## Syntax
```go
With(struct{}, Struct(...StructFn))
```
## Struct Functions
### Add key if an error happen
```go
WithKey(string)
```
### Add validate functions to a field
```go
Field(string, ...Rule)
```