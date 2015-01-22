package data

type AnimationFrame struct {
	Frame string
	Value string
}

type Animation struct {
	Name string
	RefName string
	Target string
	Kind string
	FrameRate string
	FrameValueType string
	Type string
	Frames []AnimationFrame
}