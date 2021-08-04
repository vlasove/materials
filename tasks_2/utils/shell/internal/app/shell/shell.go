package shell

type Shell struct {
	Args     []string
	Executor Executor
}

func New() *Shell {
	return &Shell{}
}

func (s *Shell) SetArgs(args []string) {
	s.Args = args
}

func (s *Shell) SetExecutor(e Executor) {
	s.Executor = e
}

func (s *Shell) Start() (string, error) {
	return s.Executor.Execute(s)
}
