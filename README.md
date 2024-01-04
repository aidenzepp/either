# either

*Either* is a Go library that provides types handling "either" functionality: `Either`, `Option`, and `Result`. The library is inspired by the Rust built-in `Option` and `Result` types, as well as the Rust `either` crate.

## Why use it?

- Easy to use
- Allows for simplifying code logic
- Provides various powerful functionalities
- No external dependencies

# Usage

```
go get github.com/aidenzepp/either
```

Then, simply import it into your Go codebase!
```go
import "github.com/aidenzepp/either"
```

# Examples

Let's say, for the sake of simplicity, you have a payload structure (e.g., `Person` seen below) that may or may not contain some data. Using the zero value may be useful, but it could be misconstrued as an actual value.
```go
package main 

type Person struct {
    Name string
    Pet string

    hasPet bool
}

// --snip--

func (p Person) HasPet() bool {
    return p.hasPet
}

```


This is where something like the `Option` type comes in...

```go
package main

import (
    // --snip--
    "github.com/aidenzepp/either"
    // --snip--
)

type Person struct {
    Name string
    Pet either.Option[string]
}

// --snip--

func (p Person) HasPet() bool {
    return p.Pet.IsValue()
}
```

Not only is this more intuitive, but it also creates simpler API structures for users of your library or product.


