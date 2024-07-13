13:*******
main is the special package name, it tells go that this is package is the main entry package for the program.

14:*******
`go mod init` command turns project into a module, a module can have multiple packages.
go mod init initiatize a module, adds go.mod,
then run go build to create a executable file,
this creates a file same as the path given in mod, if 'module example.com/first-app' then executable file name is first-app.
This executable can be run with ./first-app command, should be in same folder or give full path.

So package main is required to know the entry point of the program and and required otherwise go build will not work.

15:*******
Along with package main, func main is also mandatory to start the execution when application starts.
There should only be one main function, or there will be error.
If we build a library then this main function is not required as we are not running it as application.

22:*******
var is used to declare variable, but if we use inference := then, var not required.
const for constant variables.

31:
fmt.Sprintf returns the formatted string, printf outputs string in condolr, does not return anything.

Section 3: Packages
58:
Every package must go into its own folder/subfolder, the folder name should be same as the package name.

60:
The functions that start with uppercase are only exported from the package, only those functions we can use
outside the package into another package.

61:
go package discovery: https://pkg.go.dev/google.golang.org/api/discovery/v1
`go get` is used to get all the third party packages used by your project.
To use a third party library:
1 - go to discovery url, search and run go get <package github link>.
2 - This adds require statement in go.mod.
3 - Run go get to get the dependency and then add in the import statement.