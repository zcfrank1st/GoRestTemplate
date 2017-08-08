package define

func IniConfTemplate() string {
    return `package define

import (
    "github.com/go-ini/ini"
    "log"
)

var (
    Connection string
    {{ range $index, $value := .IniKeys }}{{ $value | Title }} string
    {{ end }}
)

func init () {
    conf, err := ini.Load(Config)
    if err != nil {
        log.Panic("init config error")
    }

    Connection = conf.Section(Environment).Key("Connection").String()
    {{ range $index, $value := .IniKeys }}{{ $value | Title }} = conf.Section(Environment).Key("{{ $value | Title }}").String
    {{ end }}
}`
}
