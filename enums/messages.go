package enums

var (
	DangerousCommand  = "The %s command can be dangerous to use remotely"
	CantReadOutputCmd = "I cannot read the output of the command: %s"
	CantSendOutputCmd = "I cannot send the output of the command: %s"
	CantExecCommand   = "The command cannot be executed: %s"
	PortAlreadyUsed   = "failed to listen: listen tcp4 :%s: bind: address already in use"
)
