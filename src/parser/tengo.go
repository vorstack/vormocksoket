package parser

import (
	"context"
	"github.com/d5/tengo/v2"
	"github.com/d5/tengo/v2/stdlib"
)

const VARIABLE_NAME_TO_GET = "result"

func tengoParse(raw []byte) []interface{} {
	script := tengo.NewScript(raw)
	script.SetImports(stdlib.GetModuleMap("json", "fmt"))
	compiled, err := script.RunContext(context.Background())
	if err != nil {
		panic(err)
	}
	variable := compiled.Get(VARIABLE_NAME_TO_GET)
	return variable.Array()
}
