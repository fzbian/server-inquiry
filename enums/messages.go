package enums

var (
	DangerousCommand  = "\nThe %s command can be dangerous to use remotely"
	CantReadOutputCmd = "\nI cannot read the output of the command: %s"
	CantSendOutputCmd = "\nI cannot send the output of the command: %s"
	CantExecCommand   = "\nThe command cannot be executed: %s"
	CantCloseFile     = "\nI cannot close the file: %s"
	CantWriteFile     = "\nI cannot write to the file: %s"
	ParamsRequired    = "\nToken and command required"
	WrongToken        = "\nThe indicated token is invalid"
	CantClearTerminal = "\nI cannot delete the terminal: %s"
)
