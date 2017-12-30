package main

// commands executed at startup
func startup() {
	doExec(nil, []string{"spotify"})
	doExec(nil, []string{"pavucontrol"})
	doExec(nil, []string{defaultBrowser})
}
