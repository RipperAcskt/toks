module main

go 1.18

replace toks/serialDriver => ../serialDriver

require toks/serialDriver v0.0.0

require (
	github.com/tarm/serial v0.0.0-20180830185346-98f6abe2eb07 // indirect
	golang.org/x/sys v0.0.0-20220907062415-87db552b00fd // indirect
)
