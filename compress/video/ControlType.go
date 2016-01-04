package video

// ControlType defines how to interpret a ControlWord
type ControlType byte

const (
	CtrlParamTwoPixel = ControlType(0)
	CtrlParamTwoIndex = ControlType(1)

	CtrlParamOffset4Pixel = ControlType(2)
	CtrlParamOffset8Pixel = ControlType(3)
	CtrlParamOfset16Pixel = ControlType(4)

	CtrlSkip = ControlType(5)

	CtrlRepeatPreviousA = ControlType(6)
	Unknown             = ControlType(7)
)
