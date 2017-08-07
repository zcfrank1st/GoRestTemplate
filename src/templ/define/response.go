package define

func ResponseTemplate() string {
    return `package define

type Message struct {
    code int
    body interface{}
}
`
}