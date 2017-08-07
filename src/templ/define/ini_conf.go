package define

func IniConfTemplate() string {
    return `package define

import (
    "github.com/go-ini/ini"
    "log"
)

var (
    Connection string
)

func init () {
    conf, err := ini.Load(Config)
    if err != nil {
        log.Panic("init config error")
    }

    Connection = conf.Section(Environment).Key("DBConnection").String()
}
`
}
