package templates

var HaskellTemplate = `module {{ .Pascal }} where

import Extism.PDK
import Extism.PDK.JSON

{{ .Camel }} = do
  -- body of step goes here

foreign export ccall "{{ .Snake }}" {{ .Camel }} ::  IO ()
`
