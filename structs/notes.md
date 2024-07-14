Section 5: Structs & custom types
When struct is passed to a fucntion a seperate copy is created and then it is passed.

When you use fmt.Println with a pointer, it does not directly print the pointer's address. Instead, it prints the dereferenced value of that pointer.
check chatgpt convo for "pass by value in go".

We first dereference the pointer variable with * and access its value, but in case of structs, directly we can use the pointer variable
to refer to value, however we can also use dereferencing.
so u.firstName and (*u).firstName both will work in case of structs.

81:
In Go, when you define a method with a receiver, the behavior of that method depends on whether the receiver is a value receiver or a pointer receiver.

In code, the method outputUserDetails2 is defined with a value receiver (u user). This means that the method operates on a copy of the user struct, not the original struct itself. Therefore, any modifications made inside outputUserDetails2 do not affect the original appUser struct.

82:
While mutating the value using reciver pattern, if we want to muatate the original value then we have to use pointer reciever becuase normal reciever will work just as variables where copy of variable is worked upon, so if we want the change on original data and not on copy then we should use pointer recivers.

85:
Inside a struct also if the field name is in capitals then only it is available outside the package, otherwise it is not, even if
struct itself is available.

87:
go has no classes and no inheritence.
we can give a struct inside other struct to extend the functionality.