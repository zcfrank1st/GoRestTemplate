package define

func CmdConfTemplate () string {
    return `package define

import "flag"

var (
    Environment string
    Config string
    {{.FlagVars}}
)

func init() {
    flag.StringVar(&Environment, "env","dev","server run environment")
    flag.StringVar(&Config, "conf","/root/app.ini","config path")
    {{.FlagSegments}}
    flag.Parse()
}
`
}
