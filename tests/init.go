package test

import (

	"errors"
	"flag"
	"testing"
)

var (
	ENV = ""
)

func init() {
	var _ = func() bool {
		testing.Init()
		return true
	}()
	env := flag.String("env", "", "environment to point integration test at")
	flag.Parse()
	if env != nil {
		ENV = *env
	} else {
		panic(errors.New("env flag was not parsed"))
	}
}
