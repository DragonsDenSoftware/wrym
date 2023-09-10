package templates

var AssemblyScriptTemplate = `import { Host } from '@extism/as-pdk';

export function {{ .Snake }}(): i32 {
  let str = Host.inputString();

  // body of step goes here

  return 0;
}
`
