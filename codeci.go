package main 

import (
	"fmt"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"strings"
	"os/exec"
    "os"
)

const version = "0.0.2"

func check(e error){
	if e != nil {
		panic(e)
	}
}

type CodeCi struct {
	Os string
	Language string
	Script []string
}

func createTestScript(codeci CodeCi) string {
    s := []string{"#!/bin/bash", "\n", "\n", strings.Join(codeci.Script, " && "), "\n"}
    return strings.Join(s, "")
}

func createDockerFile(codeci CodeCi) string{
    if codeci.Language == "none" {
        s := []string{"FROM therickys93/", codeci.Os, "\n", "ADD . /app\nWORKDIR /app\nCMD [\"bash\", \"test.sh\"]", "\n"}
        return strings.Join(s, "")
    } else {
        s := []string{"FROM therickys93/", codeci.Os, codeci.Language, "\n", "ADD . /app\nWORKDIR /app\nCMD [\"bash\", \"test.sh\"]", "\n"}
        return strings.Join(s, "")
    }
} 



func main() {
    if len(os.Args) > 1 {
        if os.Args[1] == "--version" {
            fmt.Printf("%s version: %s\n", os.Args[0], version)
            os.Exit(0)
        }
    }
	data, err := ioutil.ReadFile("./codeci.yml")
	check(err)
    fmt.Printf("reading the codeci.yml file...\n")
	fmt.Print(string(data))
    fmt.Printf("\n")
	var codeci CodeCi

	err = yaml.Unmarshal([]byte(string(data)), &codeci)
	check(err)

    fmt.Printf("Creating temp files...\n")
	// create the test.sh file
	d1 := []byte(createTestScript(codeci))
    err = ioutil.WriteFile("./test.sh", d1, 0644)
    check(err)

    // create the Dockerfile
    d1 = []byte(createDockerFile(codeci))
    err = ioutil.WriteFile("./Dockerfile", d1, 0644) 
    check(err)

    // create the docker-compose.yml file
    s := []string{"sut:\n", "  build: .\n", "  dockerfile: Dockerfile", "\n"}
    d1 = []byte(strings.Join(s, ""))
    err = ioutil.WriteFile("./docker-compose.yml", d1, 0644)
    check(err)

    // create the onlytest.sh file
    s = []string{"#!/bin/bash", "\n", "\n", "docker-compose -f docker-compose.yml -p ci build", "\n", "echo running the script...", "\n", "docker-compose -f docker-compose.yml -p ci up -d", "\n", "docker logs -f ci_sut_1", "\n", "echo check if the number is 0 for all good...",  "\n", "docker wait ci_sut_1", "\n", "docker-compose -f docker-compose.yml -p ci kill", "\n", "docker rm ci_sut_1", "\n", "docker rmi ci_sut"}
    d1 = []byte(strings.Join(s, ""))
    err = ioutil.WriteFile("./onlytest.sh", d1, 0644)
    check(err)

    // run the script onlytest.sh
    fmt.Print("run the build...\n")
    out, err := exec.Command("/bin/bash", "./onlytest.sh").Output()
    check(err)
    fmt.Print(string(out))

    // remove all the files
    fmt.Print("removing the temp files...\n")
    os.Remove("./test.sh")
    os.Remove("./Dockerfile")
    os.Remove("./onlytest.sh")
    os.Remove("./docker-compose.yml")

    fmt.Print("done!\n")
}