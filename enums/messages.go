package enums

var (
	DangerousCommand  = "The %s command can be dangerous to use remotely"
	CantReadOutputCmd = "I cannot read the output of the command: %s"
	CantSendOutputCmd = "I cannot send the output of the command: %s"
	CantExecCommand   = "The command cannot be executed: %s"
	PortAlreadyUsed   = "failed to listen: listen tcp4 :%s: bind: Only one usage of each socket address (protocol/network address/port) is normally permitted."
	CantCloseFile     = "I cannot close the file: %s"
	CantWriteFile     = "I cannot write to the file: %s"
	ParamsRequired    = "Token and command required"
	WrongToken        = "The indicated token is invalid"
)
