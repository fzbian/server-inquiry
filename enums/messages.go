package enums

var (
	DangerousCommand  = "The %s command can be dangerous to use remotely"
	CantReadOutputCmd = "I cannot read the output of the command: %s"
	CantSendOutputCmd = "I cannot send the output of the command: %s"
	CantCloseFile     = "I cannot close the file: %s"
	CantWriteFile     = "I cannot write to the file: %s"
	ParamsRequired    = "Token and command required"
	WrongToken        = "The indicated token is invalid"
	CantClearTerminal = "I cannot delete the terminal: %s"
	CantStartServer   = "I cannot initialize the server: %s"
	CantSaveLog       = "I cannot save log:  %s"
	CommandNotFound   = "Command \"%s\" was not found"
)
