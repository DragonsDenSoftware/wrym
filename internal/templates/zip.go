package templates

var ZigTemplate = `const std = @import("std");
const extism_pdk = @import("extism-pdk");
const Plugin = extism_pdk.Plugin;
const http = extism_pdk.http;

pub fn main() void {}
const allocator = std.heap.wasm_allocator;

export fn {{ .Snake }}() i32 {
    // body of step goes here
	return 0
}
`
