package templates

var HaskellTemplate = `module {{ .Pascal }} where

import Extism.PDK
import Extism.PDK.JSON

{{ .Camel }} = do
  -- Get input string from host
  s <- inputString

  -- replace lines 14-17 with step body

  -- Calculate the number of vowels
  let count = length (filter isVowel s)
  -- Return a JSON object {"count": count} back to the host
  outputJSON $ object ["count" .= count]


foreign export ccall "{{ .Snake }}" {{ .Camel }} ::  IO ()
`
