package app

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

var (
	name   = "empty"
	env    = "development"
	branch = "empty"
	commit = "empty"
	vPtr   *bool
)

func init() {
	vPtr = flag.Bool("v", false, "prints application data, then exits")
}

// Init ... checks the version flag, prints version, if set
func Init() {
	if !flag.Parsed() {
		fmt.Println("flag.Parse must be called before app.Init()")
		os.Exit(1)
	}
	if !*vPtr {
		return
	}
	version()
	os.Exit(0)
}

func version() {
	v := struct {
		FullName string
		Branch   string
		Commit   string
	}{
		FullName: FullName(),
		Branch:   branch,
		Commit:   commit,
	}
	data, err := json.MarshalIndent(&v, "", "\t")
	if err != nil {
		fmt.Printf("failed to marshal version data: %v", err)
		os.Exit(1)
	}
	fmt.Println(string(data))
}

// FullName ... returns grace-env-name where
// env and name are values passed into the linker at compile time
func FullName() string {
	return fmt.Sprintf("grace-%s-%s", env, name)
}

// Name ... returns the value of name
// default is 'empty'
func Name() string {
	return name
}

// Env ... returns the value of env
// default is 'empty'
func Env() string {
	return env
}
