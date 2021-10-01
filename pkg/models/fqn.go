package models

type FQN struct {
	Namespace_ string
	ClassName_ string
	Name_      string
}

func (f FQN) Name() string {
	return f.Name_
}

func (f FQN) Namespace() string {
	return f.Namespace_
}

func (f FQN) ClassName() string {
	return f.ClassName_
}

func (f FQN) FQN() string {
	return f.String()
}

func (f FQN) Func() string {
	return f.Namespace_ + `\` + f.Name_
}

func (f FQN) Method() string {
	return f.Namespace_ + `\` + f.ClassName_ + "::" + f.Name_
}

func (f FQN) String() string {
	if f.ClassName_ != "" {
		return f.Method()
	}
	return f.Func()
}
