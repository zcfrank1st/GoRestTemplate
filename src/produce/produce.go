package build

import (
    "os"
    "strings"
    "templ/define"
    "path/filepath"
    "templ/project_util"
    "templ/bootstrap"
    "os/exec"
    "text/template"
    "templ/service"
    "fmt"
    "github.com/logrusorgru/aurora"
)

var funcMap = template.FuncMap {
    "Title": strings.Title,
    "ToLower": strings.ToLower,
}

type (
    InitValueTemplate struct {
        FlagVars []string
        IniKeys  []string
    }

    BootstrapValueTemplate struct {
        Project string
        Services []string
    }
)

func GenerateSkeleton(absolute_path string, project string, services string, flags string, inis string) {
    fmt.Println(aurora.Blue("[GoRestT] {{absPath}} : " + absolute_path))
    fmt.Println(aurora.Blue("[GoRestT] {{project}} : " +  project))
    fmt.Println(aurora.Blue("[GoRestT] {{service}} : " + services))
    fmt.Println(aurora.Blue("[GoRestT] {{flag}} : " + flags))
    fmt.Println(aurora.Blue("[GoRestT] {{ini}} : " + inis))

    bootstrapPath := filepath.Join(absolute_path, "src", "bootstrap")
    definePath := filepath.Join(absolute_path, "src", "define")
    servicePath := filepath.Join(absolute_path, "src", "service")
    configPath := filepath.Join(absolute_path, "src")

    // create dirs
    fmt.Println(aurora.Cyan("[GoRestT] building project dirs ..."))
    if err:= os.MkdirAll(bootstrapPath, os.ModePerm); err != nil {
        fmt.Println(aurora.Red("create bootstrap dir failed"))
        os.Exit(1)
    }
    if err:= os.MkdirAll(definePath, os.ModePerm); err != nil {
        fmt.Println(aurora.Red("[GoRestT] create define dir failed"))
        os.Exit(1)
    }
    if err:= os.MkdirAll(servicePath, os.ModePerm); err != nil {
        fmt.Println(aurora.Red("[GoRestT] create service dir failed"))
        os.Exit(1)
    }
    fmt.Println(aurora.Cyan("[GoRestT] dirs build done ..."))

    // create files
    fmt.Println(aurora.Cyan("[GoRestT] building project files ..."))
    iniKeys := strings.Fields(inis)
    flagVars := strings.Fields(flags)
    initValue := InitValueTemplate{flagVars, iniKeys}

    makeFileWithTemplate(define.InitTemplate(), initValue, []string{filepath.Join(definePath, "init.go")})
    makeFileWithTemplate(project_util.IniTemplate(), initValue, []string{filepath.Join(configPath, "app.ini")})

    servicesArray := strings.Fields(services)
    var servicePaths []string
    for _, serv := range servicesArray {
        servicePaths = append(servicePaths, filepath.Join(servicePath, serv + "_service.go"))
    }

    makeFileWithTemplate(service.SimpleServiceTemplate(), servicesArray, servicePaths)
    makeFileWithTemplate(bootstrap.SimpleBootstrapTemplate(), BootstrapValueTemplate{project, servicesArray} , []string{filepath.Join(bootstrapPath, "bootstrap.go")})

    makeFile(filepath.Join(definePath, "error.go"), define.ErrorTemplate())
    makeFile(filepath.Join(definePath, "response_code.go"), define.ResponseCodeTemplate())
    makeFile(filepath.Join(configPath, "glide.yaml"), project_util.GlideTemplate())
    makeFile(filepath.Join(configPath, "Makefile"), project_util.MakefileTemplate())

    makeFile(filepath.Join(absolute_path, ".gitignore"), project_util.GitIgnoreTemplate())
    makeFile(filepath.Join(absolute_path, "README.md"), project_util.ReadmeTemplate())
    fmt.Println(aurora.Cyan("[GoRestT] project files build done..."))

    fmt.Println(aurora.Cyan("[GoRestT] loading project dependencies ..."))
    cmd := exec.Command("/bin/bash", "-c", "cd " + absolute_path + "/src; export GOPATH="+ absolute_path +"; glide install;")
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    if err := cmd.Run(); err != nil {
        fmt.Println(aurora.Red("[GoRestT] project dependencies load error, check GOPATH or glide if set proper"))
        os.Exit(1)
    }

    fmt.Println(aurora.Green("[GoRestT] project dependencies loading done..."))
    fmt.Println(aurora.Green("[GoRestT] project init done"))
}

func makeFile(path string, content string) (err error){
    if file, err := os.Create(path); err == nil {
        defer file.Close()
        _, err = file.WriteString(content)
    }
    return
}

func makeFileWithTemplate(templateString string, valueTemplate interface{}, s []string) (er error){
    for idx, value := range s {
        tt := template.Must(template.New(value).Funcs(funcMap).Parse(templateString))
        if file, err := os.Create(value); err == nil {
            if r, ok := valueTemplate.([]string); ok {
                er = tt.Execute(file, r[idx])
            } else {
                er = tt.Execute(file, valueTemplate)
            }
            file.Close()
        }
    }

    return
}

