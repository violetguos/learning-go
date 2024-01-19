# learning-go

https://youtu.be/iDQAZEJK8lI?si=8UfIbFKzww53ju8r

BASED on https://github.com/matt4biz/go-class-slides/blob/trunk/xmas-2020/

# Go facts 
## go-02-example1-slides
- `go mod init <module name>` to generate the go.mod file
- no circular dependency is allowed in the package imports

## go-03-basics
varaible initialization
```go
var a int
var ( 
    b=2
    f = 2.01 
)
```

only in functions. type is implied. must be explicitly cast into another type
```go
some_var := 2.3
```

- pointers
    - are allowed (more on this later)
    - cannot do pointer manipulation unless using a speicifc library (thank god!!!)
- all initialization gets a 0, 00.0, nil, false, default value
- Boolean values and int are not the same (unlike C)

## go-04-strings
• byte: a synonym for uint8
• rune: a synonym for int32 for characters
• string: an immutable sequence of “characters” 
• physically a sequence of bytes (UTF-8 encoding) 
• logically a sequence of (unicode) runes
a non english character has length > 1, `len(é)` is 2

## go-05-complex
### array
- fixed length
- can only copy if type and length match 
- passed by value
### slice
- length not fixed
- feels like C++ vector
- passed by reference
### map
example decalaration
```go
// JUST DON'T do this
var m map[string]int // nil, no storage

// instead, do this
p := make(map[string]int) //non nil but empty

// look up words that don't exist
p["the"] // returns 0
m["the"] // returns 0

// insert new key value pair
m["and"] = 1 // ERROR!
p["and"] = 1 // updates the map safely

```
see https://stackoverflow.com/questions/37568346/map-types-are-reference-types-var-m-mapstringint-doesnt-point-to-an-initiali


## Misc
- every new keyword indicates a new scope
```go
if some condition; something {
 // everything inside if is a new scope
}
```
- array length is fixed
- constants are real constants. they can't be changed
- Only numbers, strings, and booleans can be constants (immutable)
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
