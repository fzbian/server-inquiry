package enums

var (
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
