package main

import (
    "testing"
)

func TestCreateTestShFile(t *testing.T) {
    t.Log("TestCreateTestShFile")
    var codeci CodeCi
    codeci.Script = []string{"echo hello", "echo CodeCi"}
    testscript := createTestScript(codeci)
    t.Log(testscript)
}

func TestCreateDockerFile(t *testing.T) {
    t.Log("TestCreateDockerFile")
    var codeci CodeCi
    codeci.Language = "none"
    codeci.Os = "ubuntu14"
    t.Log("language == none")
    dockerfile := createDockerFile(codeci)
    t.Log(dockerfile)
    codeci.Language = "java"
    t.Log("language == java")
    dockerfile = createDockerFile(codeci)
    t.Log(dockerfile)
    t.Log("image kylef/swiftenv")
    codeci.Image = "kylef/swiftenv"
    dockerfile = createDockerFile(codeci)
    t.Log(dockerfile)
}