package main

import (
    "github.com/c-bata/go-prompt"
    "github.com/logrusorgru/aurora"
    "time"
    "strings"
    "os"
    "fmt"
    "errors"
    "produce"
)

const (
    Logo = `
 ____            ____                    __    ______
/\  _ \         /\  _ \                 /\ \__/\__  _\
\ \ \L\_\    ___\ \ \L\ \     __    ____\ \ ,_\/_/\ \/
 \ \ \L_L   / __ \ \ ,  /   /'__ \ /',__\\ \ \/  \ \ \
  \ \ \/, \/\ \L\ \ \ \\ \ /\  __//\__,  \\ \ \_  \ \ \
   \ \____/\ \____/\ \_\ \_\ \____\/\____/ \ \__\  \ \_\
    \/___/  \/___/  \/_/\/ /\/____/\/___/   \/__/   \/_/
                                                    v1.0`

    Keyword_Dir     = "dir"
    Keyword_Name    = "name"
    Keyword_Service = "service"
    Keyword_Flag    = "flag"
    Keyword_Ini     = "ini"


    Keyword_Help = "help"
    Keyword_Quit = "quit"
    Keyword_Exit = "exit"

    SemicolonSplit = ";"
)

var (
    dir     string
    name    string = "demo"
    service string = "demo"
    flag    string
    ini     string
)

type GoRestTCommand struct {
    verb string
    args []string
}

func (c *GoRestTCommand) run() {
    switch strings.ToLower(c.verb) {
    case Keyword_Dir:
        dir = c.args[0]
    case Keyword_Name:
        name = c.args[0]
    case Keyword_Service:
        service = strings.Join(c.args, " ")
    case Keyword_Flag:
        flag = strings.Join(c.args, " ")
    case Keyword_Ini:
        ini = strings.Join(c.args, " ")
    }
}

func parseCommands(command_string string) ([]*GoRestTCommand, error) {
    var commands []*GoRestTCommand

    commands_segments := strings.Split(command_string, SemicolonSplit)
    for _, command_segments := range commands_segments {
        command_fields := strings.Fields(strings.Trim(command_segments, " "))
        if !checkCommandIfValid(command_fields) {
            return nil, errors.New("[GoRestT] invalid command")
        }
        commands = append(commands, &GoRestTCommand{command_fields[0], command_fields[1: ]})
    }

    return commands, nil
}

func checkCommandIfValid(fields []string) bool {
    verb := fields[0]
    args_length := len(fields[1:])
    switch verb {
    case Keyword_Dir, Keyword_Name:
        if args_length != 1 {
            return false
        }
    case Keyword_Service, Keyword_Flag, Keyword_Ini:
        if args_length == 0 {
            return false
        }
    default:
        return false
    }
    return true
}

func completer(d prompt.Document) []prompt.Suggest {
    s := []prompt.Suggest{
        {Text: "dir", Description: "Rest absolute project path"},
        {Text: "name", Description: "Rest project name which use to set root url segment"},
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
    fmt.Println(aurora.Brown(Logo))
    for {
        t := prompt.Input(
            "GoRestT ["+ time.Now().String()[:19]+"] >>> ",
            completer,
            prompt.OptionPrefixTextColor(prompt.Blue),
            prompt.OptionMaxSuggestion(9))

        if t != "" {
            command_string := strings.ToLower(t)

            switch command_string {
            case Keyword_Help:
                fmt.Println(aurora.Brown(Logo + `
dir     [absolute filepath] "absolute project path: /usr/demo"
name    [project name]      "project name which use to set root url segment: demo"
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
                commands, err := parseCommands(command_string)
                if err == nil {
                    for _, command := range commands {
                        command.run()
                    }
                    build.GenerateSkeleton(dir, name, service, flag, ini)
                } else {
                    fmt.Println(aurora.Red(err))
                }
            }
        }
    }
}
