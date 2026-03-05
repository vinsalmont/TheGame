import Cocoa

// Use NSAppleScript to send notification via osascript, but from within the .app
// so the notification shows our app icon.

let message: String
if CommandLine.arguments.count > 1 {
    message = CommandLine.arguments[1]
} else {
    message = "You lost The Game."
}

let escaped = message.replacingOccurrences(of: "\"", with: "\\\"")
let script = """
display notification "\(escaped)" with title "The Game" sound name "Funk"
"""

let appleScript = NSAppleScript(source: script)
appleScript?.executeAndReturnError(nil)
