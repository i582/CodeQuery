package position

// Pos represents node position
type Pos struct {
	StartLine int
	EndLine   int
	StartPos  int
	EndPos    int
}

// NewPos Pos constructor
func NewPos(StartLine int, EndLine int, StartPos int, EndPos int) *Pos {
	return &Pos{
		StartLine: StartLine,
		EndLine:   EndLine,
		StartPos:  StartPos,
		EndPos:    EndPos,
	}
}
