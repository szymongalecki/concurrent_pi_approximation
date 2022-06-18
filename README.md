### Concurrent implementation of approximating π using Leibniz formula
This minimal program encompasses concurrent design patterns used in the Go programming language 

concurrency:
- separate channels for positive and negative terms of the series
- addition, subtraction and quiting calculations is handled by the select statement
- value of π is updated and read by one goroutine only, this eliminates race condition
- waitgroups guarantee execution of every goroutine
- separate channel for quiting is used after all terms are sent and consumed
- terms are consumed before quiting because they are sent over non-buffered channels which enforce immediate synchronisation

usage:
- number of terms is always even
- pair is equal to two terms
- change number of **pairs** and run
- observe the output

resolution:
- six numbers after comma for pairs <= 1000
- change number of places after comma for bigger number of pairs

accuracy:
- order of operations is not predictable and will influence the result when calculated for a big number of terms
- error propagation of addition and subtraction combined with representation of float numbers

references:   
- Leibniz formula for π - https://en.wikipedia.org/wiki/Leibniz_formula_for_%CF%80
- IEEE 754 - https://en.wikipedia.org/wiki/IEEE_754
