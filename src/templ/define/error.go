package define

func ErrorTemplate() string {
    return `package define

import "errors"

var (
    SystemError = errors.New("system error ...")
)
`
}
