package main

import (
    "github.com/c-bata/go-prompt"
    "time"
    "strings"
    "os"
    "fmt"
    "github.com/logrusorgru/aurora"
)

const (
    Keyword_Help = "help"
    Keyword_Quit = "quit"
    Keyword_Exit = "exit"
)

var (
    dir     string
    project string = "demo"
    service string = "demo"
    flag    string
    ini     string
)

func completer(d prompt.Document) []prompt.Suggest {
    s := []prompt.Suggest{
        {Text: "dir", Description: "Rest absolute project path"},
        {Text: "project", Description: "Rest project name which use to set root url segment"},
        {Text: "service", Description: "Rest service names"},
        {Text: "flag", Description: "Rest flag keys"},
        {Text: "ini", Description: "Rest config ini keys"},

        {Text: "help", Description: "Help with the command shell"},
        {Text: "quit/exit", Description: "Quit the shell"},
        {Text: "", Description: "[Tips] use ; to split commands, for example: dir /usr/local; project hello;..."},
    }

    return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func main() {
    for {
        t := prompt.Input(
            "GoRestT ["+ time.Now().String()[:19]+"] >>> ",
            completer,
            prompt.OptionPrefixTextColor(prompt.Blue),
            prompt.OptionMaxSuggestion(9))

        switch strings.ToLower(t) {
        case Keyword_Help:
            fmt.Println(aurora.Brown(`
 ____            ____                    __    ______
/\  _ \         /\  _ \                 /\ \__/\__  _\
\ \ \L\_\    ___\ \ \L\ \     __    ____\ \ ,_\/_/\ \/
 \ \ \L_L   / __ \ \ ,  /   /'__ \ /',__\\ \ \/  \ \ \
  \ \ \/, \/\ \L\ \ \ \\ \ /\  __//\__,  \\ \ \_  \ \ \
   \ \____/\ \____/\ \_\ \_\ \____\/\____/ \ \__\  \ \_\
    \/___/  \/___/  \/_/\/ /\/____/\/___/   \/__/   \/_/
                                                    v1.0

dir     [absolute filepath] "absolute project path: /usr/demo"
project [name]              "project name which use to set root url segment: demo"
service [service names]     "service names: hello word haha ..."
flag    [flag names]        "flag keys: Env Secret ..."
ini     [ini names]         "config ini keys: Config Connect ..."

help                        "Help with the command shell"
quit/exit                   "Quit the shell"
[Tips] use ; to split commands, for example: dir /usr/local; project hello;...`))

        case Keyword_Quit, Keyword_Exit:
            fmt.Println(aurora.Red("[GoRestT] quiting ..."))
            os.Exit(1)
        default:


        }
    }
}
