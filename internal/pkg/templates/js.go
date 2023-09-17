package templates

var JsTemplate = `function {{ .Snake }}() {
    let input = Host.inputString()
    let output = 0

	// body of step goes here

    Host.outputString(output)
}`
