# Conditional
Conditional verification
## Conditional Function
### If
```go
If().
	Then(fns ...Rule).
	Else(fns ...Rule)
```
### SwitchCase
```go
Switch().
	Case(value interface{}, fns ...Rule).
	CaseMany(value []interface{}, fns ...Rule).
	Default(fns ...Rule)
```