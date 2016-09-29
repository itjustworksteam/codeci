package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestCreateTestShFile(t *testing.T) {
    t.Log("TestCreateTestShFile")
    var codeci CodeCi
    codeci.Script = []string{"echo hello", "echo CodeCi"}
    testscript := createTestScript(codeci)
    assert.Equal(t, testscript, "#!/bin/bash\n\necho hello && echo CodeCi\n", "should be equals")
}

func TestOfficialImages(t *testing.T) {
    t.Log("TestOfficialImages")
    t.Log(officialImages())
    assert.Equal(t, officialImages(), "therickys93/ubuntu14\ntherickys93/ubuntu14node\ntherickys93/ubuntu14java\ntherickys93/ubuntu14swiftenv\ntherickys93/ubuntu14python\ntherickys93/ubuntu14php\ntherickys93/ubuntu14go\ntherickys93/ubuntu14cpp\ntherickys93/ubuntu14scala\ntherickys93/ubuntu14ruby\n", "should be equals")
}

func TestCreateDockerFileWithAssert(t *testing.T) {
    t.Log("TestCreateDockerFileWithAssert")
    var codeci CodeCi
    codeci.Os = "ubuntu14"
    codeci.Language = "java"
    dockerfile := createDockerFile(codeci)
    t.Log("language == java")
    assert.Equal(t, dockerfile, "FROM therickys93/ubuntu14java\nADD . /app\nWORKDIR /app\nCMD [\"bash\", \"test.sh\"]\n", "should be equals")
    codeci.Language = "none"
    dockerfile = createDockerFile(codeci)
    t.Log("language == none")
    assert.Equal(t, dockerfile, "FROM therickys93/ubuntu14\nADD . /app\nWORKDIR /app\nCMD [\"bash\", \"test.sh\"]\n", "should be equals")
    codeci.Image = "kylef/swiftenv"
    dockerfile = createDockerFile(codeci)
    t.Log("image == kylef/swiftenv")
    assert.Equal(t, dockerfile, "FROM kylef/swiftenv\nADD . /app\nWORKDIR /app\nCMD [\"bash\", \"test.sh\"]\n", "should be equals")
}