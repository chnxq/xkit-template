package modules

import (
	modulehost "xkit-template-v01/shared/modulehost"
)

var modules = []modulehost.Module{
	// Register host modules here.
	// Each module should implement shared/modulehost.Module and expose
	// its own module entry, so the host only maintains this module table.

	// New Module place here
}

func RegisteredHostModules() []modulehost.Module {
	return modules
}
