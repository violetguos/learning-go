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
- no circular dependency is allowed in the package imports
- array length is fixed
- constants are real constants. they can't be changed
- Closure and scope 

you can return the pointer
```go
func some_func() *int {
    var some_var int
    // do something
    return &some_var
}
```
this isn't allowed in C. Go allows some variables declared in one scope to live longer than the scope itself.
