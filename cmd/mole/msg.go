package main

const (
	msgDebugEnabled = `Debug output enabled.`

	msgErrGainRoot = `Could not gain root privileges (%v) to execute "%s".
Mole uses root euid only during operations that require it, otherwise falling
back to the regular user id. To give mole root access, execute it using sudo.`

	msgErrNoVPN = `No VPN provider for "%s" available.`

	msgErrIncorrectFwd     = `Badly formatted fwd command %q.`
	msgErrIncorrectFwdSrc  = `Badly formatted fwd source %q.`
	msgErrIncorrectFwdDst  = `Badly formatted fwd destination %q.`
	msgErrIncorrectFwdIP   = `Cannot forward from non-existent local IP %q.`
	msgErrIncorrectFwdPriv = `Cannot forward from privileged port %q (<1024).`
	msgErrNoSuchCommand    = `No such command %q. Try "help".`

	msgErrNoHome = `No home directory that I could find; cannot proceed.`

	msgWarnSetrlimit = `Warning: setrlimit: %v`

	msgErrPEMNoKey = `No ssh key found after PEM decode.`

	msgSshFirst = `ssh: Dial %s@%s`
	msgSshVia   = `ssh: Tunnel to %s@%s`

	msgVpncStart     = `vpnc: Started (pid %d).`
	msgVpncStopping  = `vpnc: Stopping (pid %d).`
	msgVpncWait      = `vpnc: Waiting for connect...`
	msgVpncConnected = `vpnc: Connected.`
	msgVpncStopped   = `vpnc: Stopped.`

	msgStatistics = ` -- %d bytes in, %d bytes out`
)

var (
	msgOk = bold(green("\nOK"))
)