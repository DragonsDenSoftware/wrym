package templates

// TODO: use new Go pdk

var GoTemplate = `package {{ .LowerCase }}

import (
    "strconv"

    "github.com/extism/go-pdk"
)

//export {{ .Snake }}
func {{ .Snake }}() int32 {
	// get input for step
    input := pdk.Input()

	// body of step goes here

	output := ""
    mem := pdk.AllocateString(output)

    // zero-copy output to host
    pdk.OutputMemory(mem)

    return 0
}

func main() {}
}`
