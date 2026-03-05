<p align="center">
  <img src="icon.png" alt="You Lost The Game" width="128" height="128">
</p>

<h1 align="center">You Lost The Game</h1>

<p align="center">
  <strong>A zero-configuration, enterprise-grade psychological warfare daemon for macOS.</strong>
</p>

<p align="center">
  <a href="#installation">Installation</a> •
  <a href="#usage">Usage</a> •
  <a href="#modes">Modes</a> •
  <a href="#how-it-works">How It Works</a> •
  <a href="#uninstallation">Uninstallation</a> •
  <a href="#faq">FAQ</a> •
  <a href="#license">License</a>
</p>

<p align="center">
  <img src="https://img.shields.io/badge/platform-macOS-blue" alt="Platform">
  <img src="https://img.shields.io/badge/go-1.22+-00ADD8?logo=go" alt="Go Version">
  <img src="https://img.shields.io/badge/status-you%20lost-red" alt="Status">
  <img src="https://img.shields.io/badge/license-MIT-green" alt="License">
  <img src="https://img.shields.io/badge/dependencies-0-brightgreen" alt="Dependencies">
  <img src="https://img.shields.io/badge/usefulness-0-red" alt="Usefulness">
</p>

---

## What is this?

[The Game](https://en.wikipedia.org/wiki/The_Game_(mind_game)) is a mental game where the objective is to **not think about The Game**. Whenever you think about The Game, you lose, and you must announce your loss.

**You Lost The Game** automates this experience by delivering beautifully crafted, randomized notifications to your Mac at unpredictable intervals. Install it once. Forget about it. Then remember it at the worst possible moment.

> *"I was in the middle of a production deploy and a notification popped up saying 'Per my last notification, you lost The Game.' I lost it."* — No one, but it could be you.

## Installation

### From source

```bash
git clone https://github.com/devsalmont/you-lost-the-game.git
cd you-lost-the-game
bash build.sh
```

This produces `YouLostTheGame.app` — a self-contained macOS app bundle.

### Quick install (run on boot)

```bash
bash build.sh
./YouLostTheGame.app/Contents/MacOS/you-lost-the-game --install
```

That's it. It will now run silently on every login. You will lose The Game forever.

## Usage

```bash
# Run manually (notifications every 15–120 minutes)
open YouLostTheGame.app

# Fire a single notification right now
./YouLostTheGame.app/Contents/MacOS/you-lost-the-game --now

# Chaos mode — notifications every 1–10 minutes
./YouLostTheGame.app/Contents/MacOS/you-lost-the-game --chaos

# Install as a persistent background service
./YouLostTheGame.app/Contents/MacOS/you-lost-the-game --install

# Uninstall the background service
./YouLostTheGame.app/Contents/MacOS/you-lost-the-game --uninstall
```

## Modes

| Mode | Interval | Flag | Vibe |
|------|----------|------|------|
| **Default** | 15–120 min | *(none)* | Subtle psychological torment |
| **Chaos** | 1–10 min | `--chaos` | Unhinged. For enemies. |
| **One-shot** | Instant | `--now` | Immediate regret |
| **Persistent** | 15–120 min | `--install` | Survives reboots. Survives you. |

## How It Works

```
┌─────────────────────────────┐
│     YouLostTheGame.app      │
│                             │
│  ┌───────────┐  ┌────────┐  │
│  │  Go daemon│──│ Swift  │  │
│  │  (timer)  │  │notifier│  │
│  └───────────┘  └────┬───┘  │
│                      │      │
│              ┌───────▼────┐ │
│              │  macOS     │ │
│              │notification│ │
│              │  center    │ │
│              └────────────┘ │
└─────────────────────────────┘
```

1. A **Go binary** runs in the background, sleeping for a random interval between notifications.
2. When the timer fires, it picks a random message from **38 handcrafted messages** across 7 categories.
3. It launches a **Swift notifier** embedded in the `.app` bundle, which delivers the notification through macOS Notification Center with the app's icon.
4. The `--install` flag registers a **launchd agent** (`dev.salmont.useless.youlostthegame`) that starts the daemon on login automatically.

## Message Categories

| Category | Example |
|----------|---------|
| Classic | *"lmao you lost The Game."* |
| Dramatic | *"NASA has confirmed: you lost The Game."* |
| Passive-aggressive | *"Hope you're having a great day. Anyway, you lost The Game."* |
| Philosophical | *"Schrödinger's Game: you both lost and didn't lose. Just kidding, you lost."* |
| Achievement | *"NEW PERSONAL RECORD: Fastest time losing The Game today."* |
| Corporate | *"Circling back on this: you lost The Game. Let's sync on next steps."* |
| Unhinged | *"Your FBI agent wanted me to let you know — you lost The Game."* |

## Requirements

- macOS 12+
- Go 1.22+
- A will to live (optional)

## Uninstallation

```bash
./YouLostTheGame.app/Contents/MacOS/you-lost-the-game --uninstall
```

This removes the launchd agent. You can also delete the `.app` bundle.

> **Note:** Uninstalling the app does not mean you won The Game. You cannot win The Game. You can only lose it.

## FAQ

**Q: Why?**
A: You lost The Game.

**Q: How do I win?**
A: You don't.

**Q: Can I add custom messages?**
A: Fork it. Add your worst. Open a PR. We accept all forms of suffering.

**Q: Is this malware?**
A: Technically no. Emotionally? Debatable.

**Q: Does this work on Linux/Windows?**
A: No. This is a macOS-exclusive L.

**Q: I uninstalled it but I'm still losing The Game.**
A: That's how The Game works. You're welcome.

## Contributing

Contributions are welcome. Please ensure all new messages meet our high standards of psychological damage.

```bash
# Run a quick test
./YouLostTheGame.app/Contents/MacOS/you-lost-the-game --now
```

If the notification made you feel something, the message is ready.

## License

MIT License — Do whatever you want with it. You already lost The Game anyway.

---

<p align="center">
  <sub>By reading this README, you have lost The Game.</sub>
</p>
