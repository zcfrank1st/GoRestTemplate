package project_util


func IniTemplate () string {
    return `[dev]
DBConnection="""root:djDev123456;@tcp(172.16.8.61:3306)/Item"""

[prd]
DBConnection="user:password@/dbname"
`
}