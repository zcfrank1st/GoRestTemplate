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

var (
    absolute_path string
    project string

    services string
    flags string
    iniKeys string

    funcMap = template.FuncMap {
        "Title": strings.Title,
        "ToLower": strings.ToLower,
    }
)

type (
    FlagValueTemplate struct {
        FlagVars []string
    }

    IniValueTemplate struct {
        IniKeys []string
    }

    BootstrapValueTemplate struct {
        Project string
        Services []string
    }
)

func init() {
    flag.StringVar(&absolute_path, "p", "", "set project init path")

    flag.StringVar(&project, "project", "", "set project name")
    flag.StringVar(&services, "service", "", "set init services")
    flag.StringVar(&flags, "flag", "", "set cmd flags")
    flag.StringVar(&iniKeys, "ini", "", "set ini configs")
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
    os.MkdirAll(definePath, os.ModePerm)
    os.MkdirAll(servicePath, os.ModePerm)

    flags := strings.Fields(flags)
    makeFileWithTemplate(define.CmdConfTemplate(), FlagValueTemplate{flags}, []string{filepath.Join(definePath, "cmd.go")})
    iniKeys := strings.Fields(iniKeys)
    iniValueTemplate := IniValueTemplate{iniKeys}
    makeFileWithTemplate(define.IniConfTemplate(), iniValueTemplate, []string{filepath.Join(definePath, "ini_loader.go")})
    makeFileWithTemplate(project_util.IniTemplate(), iniValueTemplate, []string{filepath.Join(configPath, "app.ini")})

    servicesArray := strings.Fields(services)
    var servicePaths []string
    for _, serv := range servicesArray {
        servicePaths = append(servicePaths, filepath.Join(servicePath, serv + "_service.go"))
    }

    makeFileWithTemplate(service.SimpleServiceTemplate(), servicesArray, servicePaths)
    makeFileWithTemplate(bootstrap.SimpleBootstrapTemplate(), BootstrapValueTemplate{project, servicesArray} , []string{filepath.Join(bootstrapPath, "bootstrap.go")})

    makeFile(filepath.Join(definePath, "database.go"), define.DatabaseTemplate())
    makeFile(filepath.Join(definePath, "error.go"), define.ErrorTemplate())
    makeFile(filepath.Join(definePath, "response_code.go"), define.ResponseCodeTemplate())
    makeFile(filepath.Join(definePath, "response.go"), define.ResponseTemplate())
    makeFile(filepath.Join(configPath, "glide.yaml"), project_util.GlideTemplate())
    makeFile(filepath.Join(configPath, "Makefile"), project_util.MakefileTemplate())
}

func makeFile(path string, content string) {
    if file, err := os.Create(path); err == nil {
        defer file.Close()
        file.WriteString(content)
    }
}

func makeFileWithTemplate(templateString string, valueTemplate interface{}, s []string) {
    for idx, value := range s {
        tt := template.Must(template.New(value).Funcs(funcMap).Parse(templateString))
        if file, err := os.Create(value); err == nil {
            if r, ok := valueTemplate.([]string); ok {
                tt.Execute(file, r[idx])
            } else {
                tt.Execute(file, valueTemplate)
            }
            file.Close()
        }
    }
}
