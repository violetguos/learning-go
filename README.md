# learning-go

https://youtu.be/iDQAZEJK8lI?si=8UfIbFKzww53ju8r

BASED on https://github.com/matt4biz/go-class-slides/blob/trunk/xmas-2020/

Go facts
- Boolean values and int are not the same (unlike C)
- every new keyword indicates a new scope
```go
if some condition; something {
 // everything inside if is a new scope
}
```
- no circular dependency is allowed
- array length is fixed
- constants are real constants. they can't be changed
