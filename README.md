Maybe
=====

Simple zero dependency struct-based generic monad implementation in Go.

## Installation

```bash
go get -u github.com/mono83/maybe
```

## Usage

### Instantiation

There are two main ways to create `maybe.Maybe` monadic container - it is either call to `maybe.Nothing()` or 
`maybe.Just(value)`. First will create empty monad, while second one will produce not empty monad with given value. 

Pay attention, that `maybe.Just` does not verify content of passed value, so constructions like 
`maybe.Just[error](nil)` can create not empty monads with nil value within. To avoid this behaviour use 
`maybe.Nilable`, that will perform `nil` check of given value.

In addition to raw instantiation, this library provide mapping constructor `maybe.Map`, able to convert 
`Maybe` from one type to another using mapping function

```go
n := maybe.Nothing[string]()    // String type, but empty
i := maybe.Just(1)              // Int type
s := maybe.Map(i, strconv.Itoa) // String type
```

### JSON

JSON bindings can be accessible in `github.com/mono83/maybe/json` package:

```go
import "github.com/mono83/maybe/json"

type Request struct {
	ID       int
	ParentID json.Maybe[int]
}

```

