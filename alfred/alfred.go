package alfred

// Alfred will handle all of our setup to get our tasks running
type Alfred struct {
	cli          *CLI
	Name         string
	RootDir      string
	Instructions []byte
}

// New will create a new Alfred object
func New(args []string) *Alfred {
	m := &Alfred{
		cli: NewCLI(args),
	}
	// set some defaults up
	m.Name = m.cli.name

	// load our instructions
	// TODO -> error handling?
	m.Instructions, _ = m.load(m.cli.file)
	return m
}
