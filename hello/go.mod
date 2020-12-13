module hello

go 1.15

// Added in; pre-production unpublished url for module
replace bjml.uk/greetings => ../greetings

// Added automatically by `$ go build` (also makes binary)
require bjml.uk/greetings v0.0.0-00010101000000-000000000000
