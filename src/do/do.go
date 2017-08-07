package main

import (
    "path"
    "flag"
    "os"
    "templ/bootstrap"
    "templ/define"
    "templ/service"
    "templ/project_util"
)

var absolute_path string

func init() {
    flag.StringVar(&absolute_path, "path", "", "set project init path")
    flag.Parse()
}

func main() {
    if !path.IsAbs(absolute_path) {
        os.Exit(1)
    }

    bootstrapPath := absolute_path + "/src/bootstrap"
    definePath := absolute_path + "/src/define"
    servicePath := absolute_path + "/src/service"
    configPath := absolute_path + "/src"

    os.MkdirAll(bootstrapPath, os.ModePerm)
    makeFile(bootstrapPath + "/bootstrap.go", bootstrap.BootstrapTemplate())

    os.MkdirAll(definePath, os.ModePerm)
    makeFile(definePath + "/cmd.go", define.CmdConfTemplate())
    makeFile(definePath + "/database.go", define.DatabaseTemplate())
    makeFile(definePath + "/error.go", define.ErrorTemplate())
    makeFile(definePath + "/ini_loader.go", define.IniConfTemplate())
    makeFile(definePath + "/response_code.go", define.ResponseCodeTemplate())
    makeFile(definePath + "/response.go", define.ResponseTemplate())

    os.MkdirAll(servicePath, os.ModePerm)
    makeFile(servicePath + "/service.go", service.ServiceTemplate())

    makeFile(configPath + "/glide.yaml", project_util.GlideTemplate())
    makeFile(configPath + "/app.ini", project_util.IniTemplate())
    makeFile(configPath + "/Makefile", project_util.MakefileTemplate())
}

func makeFile(path string, content string) {
    if file, err := os.Create(path); err == nil {
        defer file.Close()
        file.WriteString(content)
    }
}
