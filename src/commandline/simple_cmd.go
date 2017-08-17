package main

import (
    simple_flag "flag"
    "os"
    "path/filepath"
    "os/exec"
    "log"
    "produce"
)

var (
    absolute_path string
    project string

    services string
    flags string
    iniKeys string
)


func init() {
    simple_flag.StringVar(&absolute_path, "absPath", "", "set project init path")

    simple_flag.StringVar(&project, "project", "demo", "set project name")
    simple_flag.StringVar(&services, "service", "demo", "set init services")
    simple_flag.StringVar(&flags, "flag", "", "set cmd flags")
    simple_flag.StringVar(&iniKeys, "ini", "", "set ini configs")
    simple_flag.Parse()
}

func main1() {
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

    build.GenerateSkeleton(absolute_path, project, services, flags, iniKeys)
}