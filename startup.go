package main

// commands executed at startup
func startup() {
	doExec(nil, []string{"kitty"})
	doExec(nil, []string{"kitty"})

	screen := screens[0]
	doWorkspaceNew(screen.workspace, nil)
	doWorkspaceMigrate(screen.workspace, nil)

	doExec(nil, []string{"spotify"})
	doExec(nil, []string{"pavucontrol"})
	doExec(nil, []string{"google-chrome-stable"})
}
