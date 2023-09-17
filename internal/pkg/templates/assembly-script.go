package templates

var AssemblyScriptTemplate = `import { Host } from '@extism/as-pdk';

export function {{ .Snake }}(): i32 {
  let str = Host.inputString();

  // body of step goes here
  // lines 12-14 are just
  // an example

  // write data back to host for use in program
  var out = '{"output": ' + ' "output"}';
  Host.outputString(out)

  return 0;
}
`
