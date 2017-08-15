package project_util


func GitIgnoreTemplate () string {
    return `.idea
pkg
bin/
vendor/
*.lock
`
}