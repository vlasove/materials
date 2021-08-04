package shell

type Shell struct {
	Args     []string
	Executor Executor
}

func New(args []string) *Shell {
	return &Shell{
		Args: args,
	}
}

func (s *Shell) SetExecutor(e Executor) {
	s.Executor = e
}

func (s *Shell) Start() (string, error) {
	return s.Executor.Execute(s)
}
