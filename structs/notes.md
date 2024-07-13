Section 5: Structs & custom types
When struct is passed to a fucntion a seperate copy is created and then it is passed.

When you use fmt.Println with a pointer, it does not directly print the pointer's address. Instead, it prints the dereferenced value of that pointer.
check chatgpt convo for "pass by value in go".

We first dereference the pointer variable with * and access its value, but in case of structs, directly we can use the pointer variable
to refer to value, however we can also use dereferencing.
so u.firstName and (*u).firstName both will work in case of structs.
