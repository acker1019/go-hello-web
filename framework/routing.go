package framework

func RouteCommand(core Core, cmds []string) {
	cmd := cmds[0]

	switch cmd {

	case "runserver":
		addr := cmds[1] + ":" + cmds[2]
		RunServer(core, addr)
		break

	default:
		panic("Unsuppored management commands.")

	}
}
