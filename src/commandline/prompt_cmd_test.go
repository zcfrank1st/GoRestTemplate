package main

import "testing"

func TestParseCommands(t *testing.T) {
    command := "dir /usr/local/goTest"

    _, err := parseCommands(command)

    if err != nil {
        t.Fatal(err)
    }
}