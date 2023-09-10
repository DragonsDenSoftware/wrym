package templates

var CTemplate = `#include "extism/extism-pdk.h"

#include <stdio.h>

int32_t {{ .Snake }}()
{
  uint64_t length = extism_input_length();

  if (length == 0)
  {
    return 0;
  }

  /*
    body of goes step here
  */

  return 0
}
`
