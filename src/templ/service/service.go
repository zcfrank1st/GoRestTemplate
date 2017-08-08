package service

func SimpleServiceTemplate() string{
    return `package service

func GetAll{{ . | Title }}s() {

}

func Get{{ . | Title }}() {

}
func Add{{ . | Title }}() {

}
func Update{{ . | Title }}() {

}

func Delete{{ . | Title }}() {

}
`
}