package define


func InitTemplate() string {
    return `package define

import (
    "flag"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "github.com/go-ini/ini"
    "log"
)

var (
    Connection string
    {{ range $index, $value := .IniKeys }}{{ $value | Title }} string
    {{ end }}

    Environment string
    Config string
    {{ range $index, $value := .FlagVars }}{{ $value | Title }} string
    {{ end }}

    Db *sql.DB
)

func init () {
    flag.StringVar(&Environment, "env","dev","server run environment")
    flag.StringVar(&Config, "conf","","config path")
    {{ range $index, $value := .FlagVars }}flag.StringVar(&{{ $value | Title }}, "{{ $value }}","","")
    {{ end }}
    flag.Parse()

    conf, err := ini.Load(Config)
    if err != nil {
        log.Panic("init config error")
    }

    Connection = conf.Section(Environment).Key("Connection").String()
    {{ range $index, $value := .IniKeys }}{{ $value | Title }} = conf.Section(Environment).Key("{{ $value | Title }}").String
    {{ end }}

    Db, _ = sql.Open("mysql", Connection)
}`
}
