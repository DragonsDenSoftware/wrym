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

	lines 22-29 are just
	an example
  */

  int64_t output = 10;

  char out[128];
  int n = snprintf(out, 128, "{\"output\": %lld}", output);

  uint64_t offs_ = extism_alloc(n);
  extism_store(offs_, (const uint8_t *)out, n);
  extism_output_set(offs_, n)

  return 0
}
`
