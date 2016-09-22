package main 

import (
	"fmt"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"strings"
	"os/exec"
    "os"
    "bufio"
)

const version = "0.1.2"

func check(e error){
	if e != nil {
		panic(e)
	}
}

type CodeCi struct {
	Os string `yaml:"os"`
	Language string `yaml:"language"`
    Image string `yaml:"image"`
	Script []string `yaml:"script"`
}

func officialImages() string{
    s := []string{"therickys93/ubuntu14", "therickys93/ubuntu14node", "therickys93/ubuntu14java", "therickys93/ubuntu14swiftenv", "therickys93/ubuntu14python", "therickys93/ubuntu14php", "therickys93/ubuntu14go", "therickys93/ubuntu14cpp", ""}
    return strings.Join(s, "\n")
}

func createTestScript(codeci CodeCi) string {
    s := []string{"#!/bin/bash", "\n", "\n", strings.Join(codeci.Script, " && "), "\n"}
    return strings.Join(s, "")
}

func createDockerFile(codeci CodeCi) string{
    if codeci.Image != "" {
        s := []string{"FROM ", codeci.Image, "\n", "ADD . /app\nWORKDIR /app\nCMD [\"bash\", \"test.sh\"]", "\n"}
        return strings.Join(s, "")
    }
    if codeci.Language == "none" {
        s := []string{"FROM therickys93/", codeci.Os, "\n", "ADD . /app\nWORKDIR /app\nCMD [\"bash\", \"test.sh\"]", "\n"}
        return strings.Join(s, "")
    } else {
        s := []string{"FROM therickys93/", codeci.Os, codeci.Language, "\n", "ADD . /app\nWORKDIR /app\nCMD [\"bash\", \"test.sh\"]", "\n"}
        return strings.Join(s, "")
    }
} 



func main() {
    filename := "codeci.yml"
    if len(os.Args) > 1 {
        if os.Args[1] == "--version" {
            fmt.Printf("%s version: %s\n", os.Args[0], version)
            os.Exit(0)
        } else if os.Args[1] == "images" {
            fmt.Print(officialImages())
            os.Exit(0)
        } else if os.Args[1] == "--help" {
            fmt.Printf("usage: %s --> runs the build and search for the codeci.yml\n", os.Args[0])
            fmt.Printf("usage: %s --version --> show the current version\n", os.Args[0])
            fmt.Printf("usage: %s images --> show default images\n", os.Args[0])
            os.Exit(0)
        } else if os.Args[1] == "-f" {
            if strings.HasPrefix(os.Args[2], "codeci") && strings.HasSuffix(os.Args[2], ".yml") {
                filename = os.Args[2]
            } else {
                os.Exit(1)
            }
        } else {
            os.Exit(1)
        }
    }
    filenames := []string{"./", filename}
	data, err := ioutil.ReadFile(strings.Join(filenames, ""))
	check(err)
    fmt.Printf("reading the provided codeci.yml file...\n\n")
	fmt.Print(string(data))
    fmt.Printf("\n\n")
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
    s = []string{"#!/bin/bash", "\n", "\n", "docker-compose -f docker-compose.yml -p ci build", "\n", "echo running the script...", "\n", "echo -e '\n'", "\n", "docker-compose -f docker-compose.yml -p ci up -d", "\n", "docker logs -f ci_sut_1", "\n", "echo -e '\n'", "\n","echo 'BUILD EXIT CODE:'",  "\n", "docker wait ci_sut_1", "\n", "if [ $(docker wait ci_sut_1) == 0 ]; then echo -e '\nBUILD SUCCESS\n'; else echo -e '\nBUILD FAILED\n'; fi", "\n", "docker-compose -f docker-compose.yml -p ci kill", "\n", "docker rm ci_sut_1", "\n", "docker rmi ci_sut"}
    d1 = []byte(strings.Join(s, ""))
    err = ioutil.WriteFile("./onlytest.sh", d1, 0644)
    check(err)

    // run the script onlytest.sh
    fmt.Print("run the build...\n")
    cmd := exec.Command("/bin/bash", "./onlytest.sh")
    cmdReader, err := cmd.StdoutPipe()
    if err != nil {
        fmt.Fprintln(os.Stderr, "Error creating pipe", err)
        return
    }
    scanner := bufio.NewScanner(cmdReader)
    go func() {
        for scanner.Scan() {
            fmt.Printf(scanner.Text() + "\n")
        }
    }()
    err = cmd.Start()
    if err != nil {
        fmt.Fprintln(os.Stderr, "error starting command", err)
        return
    }
    err = cmd.Wait()
    if err != nil {
        fmt.Fprintln(os.Stderr, "error waiting command", err)
        return
    }

    // remove all the files
    fmt.Print("removing the temp files...\n")
    os.Remove("./test.sh")
    os.Remove("./Dockerfile")
    os.Remove("./onlytest.sh")
    os.Remove("./docker-compose.yml")

    fmt.Print("done!\n")
}