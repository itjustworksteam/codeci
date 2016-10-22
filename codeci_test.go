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
    assert.Equal(t, testscript, "#!/bin/bash\n\necho 'Job Node Info: '\necho \n\necho 'uname -a'\nuname -a\necho \n\necho 'df -h'\ndf -h\necho \n\necho 'free -m'\nfree -m\necho \n\necho 'bash --version'\nbash --version\necho \n\necho \n\necho 'running your commands: '\necho hello && echo CodeCi\n", "should be equals")
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

func TestCodeCITest(t *testing.T) {
    t.Log("Test codeci test");
    expected := "image: docker/whalesay\nscript:\n   - cowsay Hello CodeCI!"
    attual := codeCIWhalesay()
    assert.Equal(t, attual, expected, "should be equals");
}