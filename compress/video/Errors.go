package video

import (
	"fmt"
)

var FormatError = fmt.Errorf("Format Error")
var BitstreamEndError = fmt.Errorf("Reaching beyond end-of-bitstream")
