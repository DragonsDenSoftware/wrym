package templates

var RustTemplate = `use extism_pdk::*;

#[derive(Serialize)]
struct ExampleOutput {
    pub a: String,
}

#[plugin_fn]
pub fn {{ .Snake }}(input: String) -> FnResult<Json<ExampleOutput>> {
    let input = input.chars()
	
	// body of step goes here

	let output = ExampleOutput { a };
	Ok(Json(output))
}
`
