package platform

import _ "embed"

//go:embed interface.txt
var InterfaceTemplate []byte

//go:embed implementation.txt
var ImplementationTemplate []byte
