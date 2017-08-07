package define

func CmdConfTemplate () string {
    return `package define

import "flag"

var (
    Environment string
    Config string
    {{ range $index, $value := .FlagVars }}{{ $value | Title }} string
    {{ end }}
)

func init() {
    flag.StringVar(&Environment, "env","dev","server run environment")
    flag.StringVar(&Config, "conf","","config path")
    {{ range $index, $value := .FlagSegments }}flag.StringVar(&{{ $value | Title }}, "{{ $value }}","","")
    {{ end }}
    flag.Parse()
}
`
}
