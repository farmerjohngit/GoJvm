package cmd

type ArgReader struct {
	args []string
}

func (self *ArgReader) hasMoreOptions() bool {
	return len(self.args) > 0 && self.args[0][0] == '-'
}

func (self *ArgReader) readOneArg() string {
	arg := self.args[0]
	self.args = self.args[1:]
	return arg
}
