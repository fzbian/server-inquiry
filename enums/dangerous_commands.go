package enums

var (
	// PowerShell lists commands that can be dangerous when used remotely by code.
	PowerShell = []string{
		"Invoke-Mimikatz",
		"Invoke-Command",
		"Invoke-Expression",
		"Invoke-Shellcode",
		"Start-Process",
		"IEX",
		"Remove-Item",
		"Remove-User",
		"Stop-Computer",
		"Set-ExecutionPolicy",
		"New-Item",
		"New-Service",
		"Start-Service",
		"Stop-Service",
	}

	// Debian lists commands that can be dangerous when used remotely by code.
	Debian = []string{
		"rm",
		"dd",
		"echo",
		"chmod",
		"shutdown",
		"reboot",
		"wget",
		"mv",
		"chown",
		"iptables",
		"kill",
		"dkpg",
	}
)
