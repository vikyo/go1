Section 4: Pointers
For pointer code execution, go to pointer folder and then run go run .
64:
By default go creates a copy when passing values to functions.
Advantages of pointers:
- Avoid unnecessary value copies.
    - with pointers only one value is stored in memory, no new copy is created.
- Directly mutate the value of variable.
    - pass the pointer instead of value to a function.
    - the function then directly mutate the underlying value, no return is required.

68:
For a pointer, default value is nil.
nil represents the absence of an address value - i.e., a pointer pointing at no address / no value in memory.