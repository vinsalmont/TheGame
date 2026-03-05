package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"syscall"
	"text/template"
	"time"
)

var messages = []string{
	// Classic
	"You just lost The Game.",
	"Guess what? You lost The Game.",
	"The Game. You lost it.",
	"lmao you lost The Game.",

	// Dramatic
	"BREAKING NEWS: You lost The Game.",
	"We interrupt your productivity to inform you: you lost The Game.",
	"This just in from the Associated Press: you lost The Game.",
	"NASA has confirmed: you lost The Game.",

	// Passive-aggressive
	"Hey... you lost The Game. Sorry not sorry.",
	"You were doing so well... but you lost The Game.",
	"Friendly reminder: you lost The Game. Have a nice day.",
	"Hope you're having a great day. Anyway, you lost The Game.",
	"Not to be rude, but you lost The Game.",

	// Philosophical
	"If you lose The Game and nobody is around, did you really lose? Yes. You did.",
	"You haven't thought about The Game in a while. Well, now you have.",
	"In an infinite universe, you losing The Game was inevitable.",
	"Schrödinger's Game: you both lost and didn't lose. Just kidding, you lost.",

	// Achievement style
	"Achievement unlocked: Lost The Game.",
	"Your streak of not losing The Game has ended.",
	"NEW PERSONAL RECORD: Fastest time losing The Game today.",
	"Congratulations! You've lost The Game for the 1,000,000th time!",
	"Level up! You are now a Grand Master of losing The Game.",

	// Corporate
	"Per my last notification, you lost The Game.",
	"As per our earlier discussion, you lost The Game. Please advise.",
	"Circling back on this: you lost The Game. Let's sync on next steps.",
	"Just wanted to loop you in — you lost The Game.",
	"Putting a pin in your productivity to remind you: you lost The Game.",

	// Existential
	"Plot twist: you lost The Game.",
	"It's been a good run. But you lost The Game.",
	"Sorry to interrupt, but you lost The Game.",
	"You can close this notification, but you already lost The Game.",
	"This notification will disappear. The fact that you lost The Game won't.",

	// Unhinged
	"The voices told me to tell you: you lost The Game.",
	"Your FBI agent wanted me to let you know — you lost The Game.",
	"I asked ChatGPT who lost The Game. It said you.",
	"Fun fact: every 60 seconds, you lose The Game. And you just did.",
	"Google 'did I lose The Game' for a surprise. Spoiler: yes.",
	"Your mom called. She said you lost The Game.",
	"Roses are red, violets are blue, you lost The Game, and there's nothing you can do.",
}

const plistLabel = "dev.salmont.useless.youlostthegame"

var plistTemplate = template.Must(template.New("plist").Parse(`<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>Label</key>
	<string>{{.Label}}</string>
	<key>ProgramArguments</key>
	<array>
		<string>{{.Binary}}</string>
	</array>
	<key>RunAtLoad</key>
	<true/>
	<key>KeepAlive</key>
	<true/>
</dict>
</plist>
`))

func plistPath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, "Library", "LaunchAgents", plistLabel+".plist")
}

func appPath() string {
	exe, _ := os.Executable()
	resolved, _ := filepath.EvalSymlinks(exe)
	// resolved is inside YouLostTheGame.app/Contents/MacOS/
	macosDir := filepath.Dir(resolved)
	return filepath.Clean(filepath.Join(macosDir, "..", ".."))
}

func notify(msg string) {
	app := appPath()
	if _, err := os.Stat(filepath.Join(app, "Contents", "Info.plist")); err == nil {
		cmd := exec.Command("open", "-a", app, "--args", msg)
		cmd.Run()
		return
	}
	script := fmt.Sprintf(`display notification "%s" with title "The Game" sound name "Funk"`, msg)
	exec.Command("osascript", "-e", script).Run()
}

func install() {
	exe, err := os.Executable()
	if err != nil {
		fmt.Println("Error: could not determine binary path:", err)
		os.Exit(1)
	}
	exe, _ = filepath.Abs(exe)

	f, err := os.Create(plistPath())
	if err != nil {
		fmt.Println("Error creating plist:", err)
		os.Exit(1)
	}
	defer f.Close()

	plistTemplate.Execute(f, map[string]string{
		"Label":  plistLabel,
		"Binary": exe,
	})

	exec.Command("launchctl", "load", plistPath()).Run()
	fmt.Println("Installed and running. You will now lose The Game forever.")
	fmt.Println("To uninstall: you-lost-the-game --uninstall")
}

func uninstall() {
	exec.Command("launchctl", "unload", plistPath()).Run()
	os.Remove(plistPath())
	fmt.Println("Uninstalled. But you still lost The Game.")
}

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "--install":
			install()
			return
		case "--uninstall":
			uninstall()
			return
		case "--now":
			msg := messages[rand.IntN(len(messages))]
			notify(msg)
			fmt.Println("Sent:", msg)
			return
		}
	}

	minMinutes := 15
	maxMinutes := 120

	if len(os.Args) > 1 && os.Args[1] == "--chaos" {
		minMinutes = 1
		maxMinutes = 10
	}

	fmt.Println("You Lost The Game is now running.")
	fmt.Printf("Notifications every %d-%d minutes.\n", minMinutes, maxMinutes)
	fmt.Println("Press Ctrl+C to stop (but you still lost The Game).")
	fmt.Println("Tip: run with --install to make it permanent.")

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for {
			delay := time.Duration(minMinutes+rand.IntN(maxMinutes-minMinutes+1)) * time.Minute
			time.Sleep(delay)
			msg := messages[rand.IntN(len(messages))]
			notify(msg)
		}
	}()

	<-sig
	fmt.Println("\nYou can quit the app, but you still lost The Game.")
}
