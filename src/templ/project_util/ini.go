package project_util

import "strings"

var devDefaultConfig = `[dev]
DBConnection="""root:djDev123456;@tcp(172.16.8.61:3306)/Item"""`

var prdDefaultConfig = `
[prd]
DBConnection=""
`

func IniTemplate (flags string) string {
    fields := strings.Fields(flags)

    var fieldsConfig string
    for _, value := range fields {
        fieldsConfig += "\n" + value + "="
    }

    return devDefaultConfig + fieldsConfig +
        prdDefaultConfig + fieldsConfig[1:]
}