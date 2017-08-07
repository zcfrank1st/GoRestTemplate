package main

import (
    "path"
    "flag"
    "os"
    "templ/bootstrap"
    "templ/define"
    "templ/service"
    "templ/project_util"
    "path/filepath"
    "strings"
    "text/template"
)

var absolute_path string

var services string
var flags string
var inis string

type FlagValueTemplate struct {
    FlagVars []string
    FlagSegments []string
}

func init() {
    flag.StringVar(&absolute_path, "p", "", "set project init path")

    flag.StringVar(&services, "service", "", "set init services")
    flag.StringVar(&flags, "flag", "", "set cmd flags")
    flag.StringVar(&inis, "ini", "", "set ini configs")
    flag.Parse()
}

func main() {
    if !path.IsAbs(absolute_path) {
        os.Exit(1)
    }

    bootstrapPath := filepath.Join(absolute_path, "src", "bootstrap")
    definePath := filepath.Join(absolute_path, "src", "define")
    servicePath := filepath.Join(absolute_path, "src", "service")
    configPath := filepath.Join(absolute_path, "src")

    os.MkdirAll(bootstrapPath, os.ModePerm)
    makeFile(filepath.Join(bootstrapPath, "bootstrap.go"), bootstrap.BootstrapTemplate())

    os.MkdirAll(definePath, os.ModePerm)

    makeFileWithTemplate(filepath.Join(definePath, "cmd.go"), define.CmdConfTemplate(), FlagValueTemplate{strings.Fields(flags), strings.Fields(flags)})
    makeFile(filepath.Join(definePath, "database.go"), define.DatabaseTemplate())
    makeFile(filepath.Join(definePath, "error.go"), define.ErrorTemplate())
    makeFile(filepath.Join(definePath, "ini_loader.go"), define.IniConfTemplate())
    makeFile(filepath.Join(definePath, "response_code.go"), define.ResponseCodeTemplate())
    makeFile(filepath.Join(definePath, "response.go"), define.ResponseTemplate())

    os.MkdirAll(servicePath, os.ModePerm)
    // TODO multi services
    makeFile(filepath.Join(servicePath, "service.go"), service.ServiceTemplate())

    makeFile(filepath.Join(configPath, "glide.yaml"), project_util.GlideTemplate())
    makeFile(filepath.Join(configPath, "app.ini"), project_util.IniTemplate(inis))
    makeFile(filepath.Join(configPath, "Makefile"), project_util.MakefileTemplate())
}

func makeFile(path string, content string) {
    if file, err := os.Create(path); err == nil {
        defer file.Close()
        file.WriteString(content)
    }
}

func makeFileWithTemplate(path string, templateString string, flagValueTemplate FlagValueTemplate) {
    funcMap := template.FuncMap{
        "Title": strings.Title,
    }
    tt := template.Must(template.New(path).Funcs(funcMap).Parse(templateString))

    if file, err := os.Create(path); err == nil {
        defer file.Close()
        tt.Execute(file, flagValueTemplate)
    }
}
