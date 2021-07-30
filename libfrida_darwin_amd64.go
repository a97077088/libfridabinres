package libfridabinres

import _ "embed"

//go:embed frida_shared.dylib
var FridaBinRes []byte
