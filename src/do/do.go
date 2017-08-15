package main

import (
    "flag"
    "os"
    "templ/bootstrap"
    "templ/define"
    "templ/service"
    "templ/project_util"
    "path/filepath"
    "strings"
    "text/template"
    "os/exec"
    "log"
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
    InitValueTemplate struct {
        FlagVars []string
        IniKeys  []string
    }

    BootstrapValueTemplate struct {
        Project string
        Services []string
    }
)

func init() {
    flag.StringVar(&absolute_path, "absPath", "", "set project init path")

    flag.StringVar(&project, "project", "demo", "set project name")
    flag.StringVar(&services, "service", "demo", "set init services")
    flag.StringVar(&flags, "flag", "", "set cmd flags")
    flag.StringVar(&iniKeys, "ini", "", "set ini configs")
    flag.Parse()
}

func main() {
    if absolute_path == "" {
        file, _ := exec.LookPath(os.Args[0])
        dir,_ := filepath.Abs(filepath.Dir(file))
        absolute_path = filepath.Join(dir, project)
    }

    log.Printf("{{absPath}} :%s \n", absolute_path)
    log.Printf("{{project}} :%s \n", project)
    log.Printf("{{service}} :%s \n", services)
    log.Printf("{{flag}} :%s \n", flags)
    log.Printf("{{ini}} :%s \n", iniKeys)

    bootstrapPath := filepath.Join(absolute_path, "src", "bootstrap")
    definePath := filepath.Join(absolute_path, "src", "define")
    servicePath := filepath.Join(absolute_path, "src", "service")
    configPath := filepath.Join(absolute_path, "src")

    // create dirs
    log.Println("building project dirs ...")
    if err:= os.MkdirAll(bootstrapPath, os.ModePerm); err != nil {
        log.Fatal("create bootstrap dir failed")
    }
    if err:= os.MkdirAll(definePath, os.ModePerm); err != nil {
        log.Fatal("create define dir failed")
    }
    if err:= os.MkdirAll(servicePath, os.ModePerm); err != nil {
        log.Fatal("create service dir failed")
    }
    log.Println("dirs build done ...")

    // create files
    log.Println("building project files ...")
    iniKeys := strings.Fields(iniKeys)
    flags := strings.Fields(flags)
    initValue := InitValueTemplate{flags, iniKeys}

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
    log.Println("project files build done...")

    log.Println("loading project dependencies ...")
    cmd := exec.Command("/bin/bash", "-c", "cd " + absolute_path + "/src; export GOPATH="+ absolute_path +"; glide install;")
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    if err := cmd.Run(); err != nil {
        log.Fatal("project dependencies load error, check GOPATH or glide if set proper")
    }

    log.Println("project dependencies loading done...")
    log.Println("project init done")
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
