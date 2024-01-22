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

## go-06-control
`;` as a short hand declaration
```go

if err := doSomething(); err != nil { 
    return err
}
```
- packaging
  - Every name that’s capitalized is exported
- typing `type [name your type] [existing built in type]` 
    - e.g. `type x int`  -> int is now called `x`. it might make sense in some cases

## go-07-io
fmt, Args, files

## go-08-funcs

## go-09-closure

## go-10-slices-in-detail
empty vs Nil
- a slice can be empty with length 0 but it points to a speical memory location
- or it points to nil / nowhere
    -   you can append to a nil slice
- always use `len(some_slice) == 0` check instead of `some_slice == nil` check for empty slice
- append starts to overwrite the first empty item in the slice. 
- `make([]int, some_len 0, some_capacity 5)` -> a slice, no values
- `make([]int, some_len 5)` -> a slice with 5 zeros

length vs capacity
- when you append to a slice that has reached the max capacity, it *reallocs* a slice in a different address
- if a slice was declared by slicing an explicitly defined array, it points at the same memory address. appending to the slice may overwrite the original array item value.

```go
some_array = [3]int{1,2,3}
some_slice = some_array[0:2]
// some array {1, 2, 3}
// some_slice {1, 2}, capacity is 3 by default

some_slice= append(some_slice, 5)

// some_array {1, 2, 5}
// some_slice {1, 2, 5}
```

- use triple slice operator `[start_index : length_end_index : capcity_end_index]` for fine grained control when slicing from an array

```go
some_array = [3]int{1,2,3}
some_slice = some_array[0:2:2]
// some array {1, 2, 3}
// some_slice {1, 2}

some_slice= append(some_slice, 5)

// some_array {1, 2, 3}
// some_slice {1, 2, 5}
```

## go-12-struct
structs are passed by value
- the function will create a copy of it. you need to explicitly return if you want to keep the changes
dereferencing
```go
func soldAnother(a *album){
    a.copies++
    // equal
    (*a).copies++
}
```
Go is smart enought to know this is accessing a memory address (YES!!!) so we don't have to write `a->copies++` like C++

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
