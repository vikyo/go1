Section 5: Structs & custom types
When struct is passed to a fucntion a seperate copy is created and then it is passed.

When you use fmt.Println with a pointer, it does not directly print the pointer's address. Instead, it prints the dereferenced value of that pointer.
check chatgpt convo for "pass by value in go".

We first dereference the pointer variable with * and access its value, but in case of structs, directly we can use the pointer variable
to refer to value, however we can also use dereferencing.
so u.firstName and (*u).firstName both will work in case of structs.

In Go, when you define a method with a receiver, the behavior of that method depends on whether the receiver is a value receiver or a pointer receiver.

In code, the method outputUserDetails2 is defined with a value receiver (u user). This means that the method operates on a copy of the user struct, not the original struct itself. Therefore, any modifications made inside outputUserDetails2 do not affect the original appUser struct.
