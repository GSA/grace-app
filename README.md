# grace-app

## Usage

grace-app provides consistency for GRACE applications by exposing standardized pre-compiled variables. Simply `import "github.com/GSA/grace-app"` inside your application and call `app.Init()` after you have called `flag.Parse()`. An additional flag will be available `-v` which when passed will print information about your application that was passed during compilation.

`
package main

import (
	"flag"

	"github.com/GSA/grace-app"
)

func main() {
	// called here to allow the application to have more flags
	// not just 'v'
	flag.Parse()
	// must be called after flag.Parse to process the 'v' flag
	app.Init()
}
`

## Setting variables during compilation

The variables used by grace-app are declared at the top-scope and unexported, this is to prevent modification after compilation. These can be set when calling `go build` by passing the `-ldflags` parameter. Shown below is this method being used to set the `name` variable inside grace-app to `mytool`.

`
go build -o mytool -ldflags="-X github.com/GSA/grace-app.name=mytool"
`

In addition, the command outputs the resulting binary to a file called `mytool` this can be used in order to search for the fully-qualified names of other variables you may want to set. In order to do so, simply execute `go tool nm PROGRAM` where `PROGRAM` is the name of the Go binary. This will return all variables within the program with its fully-qualified name, so it may be necessary to filter them using `grep` or `findstr` if you're on Windows.

Non-Windows:
`
go tool nm mytool | grep -i name
`

Windows:
`
go tool nm mytool | findstr name
`

The two examples above show filtering the variables of the `mytool` Go binary to find variables whose names contain `name`. The fully-qualified name will generally consist of `package.variable` where package might be `github.com/GSA/grace-app` and the variable `name`. When passing ldflags it's important to ensure you're using the `-X` prior to each variable declaration. In addition, this method only supports string data, so further interpretation may be necessary if you want to use this method to pass an `int`, `[]string`, etc.