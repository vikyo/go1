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