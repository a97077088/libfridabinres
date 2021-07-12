package libfridabinres

import _"embed"

const CRC32Value uint32 = 0x605caa00
//go:embed frida_shared_windows_arm64.dll
var FridaBinRes []byte