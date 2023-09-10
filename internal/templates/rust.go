package templates

var RustTemplate = `use extism_pdk::*;

#[plugin_fn]
pub fn {{ .Snake }}(input: String) -> FnResult<Json<TestOutput>> {
    // body of step goes here
}
`
