# String
Validate functions for a String
## Syntax
```go
With(s string, String(fns ...Rule))
```
## String functions
### Require min length
```go
MinLength(int)
```
### Cap max length
```go
MaxLength(int)
```
### Match regex
```go
Match(regex string)
```
### Custom validator
```go
StringCustom(func(s string) error {
    // your logic goes here
})
```