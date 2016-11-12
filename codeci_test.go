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
    assert.Equal(t, testscript, "#!/bin/bash\n\necho 'Job Node Info: '\necho \n\necho 'uname -a'\nuname -a\necho \n\necho 'df -h'\ndf -h\necho \n\necho 'free -m'\nfree -m\necho \n\necho 'bash --version'\nbash --version\necho \n\necho 'lscpu'\nlscpu\necho \n\necho 'lsb_release -a'\nlsb_release -a\necho \n\necho 'service --status-all'\nservice --status-all\necho \n\necho 'dpkg -l'\ndpkg -l\necho \n\necho \n\necho 'running your commands: '\necho hello && echo CodeCi\n", "should be equals")
}

func TestCreateDockerFileWithAssert(t *testing.T) {
    t.Log("TestCreateDockerFileWithAssert")
    var codeci CodeCi
    codeci.Os = "ubuntu14"
    codeci.Language = "java"
    dockerfile := createDockerFile(codeci)
    t.Log("language == java")
    assert.Equal(t, dockerfile, "FROM therickys93/ubuntu14java\nADD . /app\nWORKDIR /app\nCMD [\"bash\", \"test.codeci.sh\"]\n", "should be equals")
    codeci.Language = "none"
    dockerfile = createDockerFile(codeci)
    t.Log("language == none")
    assert.Equal(t, dockerfile, "FROM therickys93/ubuntu14\nADD . /app\nWORKDIR /app\nCMD [\"bash\", \"test.codeci.sh\"]\n", "should be equals")
    codeci.Image = "kylef/swiftenv"
    dockerfile = createDockerFile(codeci)
    t.Log("image == kylef/swiftenv")
    assert.Equal(t, dockerfile, "FROM kylef/swiftenv\nADD . /app\nWORKDIR /app\nCMD [\"bash\", \"test.codeci.sh\"]\n", "should be equals")
}

func TestCodeCITest(t *testing.T) {
    t.Log("Test codeci test");
    expected := "image: docker/whalesay\nscript:\n   - cowsay Hello CodeCI!"
    attual := codeCIWhalesay()
    assert.Equal(t, attual, expected, "should be equals");
}

func TestDockerfileName(t *testing.T) {
    t.Log("Test Dockerfile name")
    expected := "Dockerfile.codeci"
    attual := dockerfileName()
    assert.Equal(t, attual, expected, "should be equals")
}

func TestDockerComposeName(t *testing.T) {
    t.Log("Test docker compose name")
    expected := "docker-compose.codeci.yml"
    attual := dockercomposeName()
    assert.Equal(t, attual, expected, "should be equals")
}

func TestOnlyTestName(t *testing.T) {
    t.Log("Test onlytest name")
    expected := "onlytest.codeci.sh"
    attual := onlytestName()
    assert.Equal(t, attual, expected, "should be equals")
}

func TestTestName(t *testing.T) {
    t.Log("Test test name")
    expected := "test.codeci.sh"
    attual := testName()
    assert.Equal(t, attual, expected, "should be equals")
}