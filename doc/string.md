# String
Validate for string
## Syntax
```go
With(s string, String(fns ...Validate))
```
## String function
### MinLength
```go
MinLength(l int)
```
### MaxLength
```go
MaxLength(l int)
```
### Match
Match regex
```go
Match(regex string)
```
### StringCustom
Custom validate
```go
StringCustom(func(s string) error {
    if s != "abc" {
    return errors.New("not abc")
    }
    return nil
}))
```