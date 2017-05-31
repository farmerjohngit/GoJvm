package cmd

const (
	_1k = 1024
	_1m = _1k * _1k
	_1g = _1k * _1m
)

type Options struct {
	classpath    string
	verboseClass int
	javaHome     string
}

func (self Options) Classpath() string {
	return self.classpath
}

func (self Options) JavaHome() string {
	return self.javaHome
}



func parseOptions(argReader *ArgReader) *Options {
	options := &Options{

	}

	for argReader.hasMoreOptions() {
		optionName := argReader.readOneArg()
		switch optionName {
		case "-cp", "-classpath":
			options.classpath = argReader.readOneArg()
		case "-verbose", "-verbose:class":
			options.verboseClass = 1
		case "-XjavaHome":
			options.javaHome = argReader.readOneArg()
		default:
			panic("Unrecognized option: " + optionName)
		}
	}

	return options
}
