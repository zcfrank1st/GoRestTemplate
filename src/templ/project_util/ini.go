package project_util

func IniTemplate () string {
    return `[dev]
Connection="""root:djDev123456;@tcp(172.16.8.61:3306)/Item"""
{{ range $index, $value := .IniKeys }}{{ $value | Title }}=
{{ end }}
[prd]
Connection=""
{{ range $index, $value := .IniKeys }}{{ $value | Title }}=
{{ end }}`
}