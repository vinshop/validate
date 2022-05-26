# Struct
Validate for struct
## Syntax
```go
With(struct{}, Struct(...StructFn))
```
## Struct Function
### With Key
Add a key to the validator ( use for frontend validate purpose )
```go
WithKey(key string)
```
### Field
Validate for field in struct
```go
Field(fiedname string, fns ...Rule)
```