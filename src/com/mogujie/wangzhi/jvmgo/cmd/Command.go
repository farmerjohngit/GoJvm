package cmd

type Command struct {
	options *Options
	file    string
	args    []string
}

func (self Command) Options() *Options {
	return self.options
}

func (self Command) File() string {
	return self.file
}

func (self Command) Args() []string {
	return self.args
}

func ParseCommand(osArgs []string) (*Command, error) {
	argReader := &ArgReader{osArgs[1:]}
	options := parseOptions(argReader)
	var class string
	if argReader.hasMoreOptions() {
		jarOp := argReader.readOneArg()
		if jarOp == "-jar" {
			class = argReader.readOneArg()
		} else {
			panic("unknow args: " + jarOp)
		}

	} else {
		class = argReader.readOneArg()
	}

	cmd :=
		&Command{
			options,
			class,
			argReader.args,
		}

	return cmd, nil
}
