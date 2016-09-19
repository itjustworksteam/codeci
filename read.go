package main 

import (
	"fmt"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"strings"
)

func check(e error){
	if e != nil {
		panic(e)
	}
}

type T struct {
	A string
	B string
	C []string
}

func main() {
	data, err := ioutil.ReadFile("./codeci.yml")
	check(err)
	fmt.Print(string(data))

	m := make(map[string]string)

	err = yaml.Unmarshal([]byte(string(data)), &m)
	check(err)

	fmt.Printf("\n%v\n", m)
	fmt.Printf("%v\n", m["os"]);
	fmt.Printf("%v\n", m["language"])
	fmt.Printf("%v\n", m["script"])

	// create the test.sh file
	s := []string{"#!/bin/bash", "\n", "\n", m["script"], "\n"}
	d1 := []byte(strings.Join(s, ""))
    err = ioutil.WriteFile("./test.sh", d1, 0644)
    check(err)

    // create the Dockerfile
    s = []string{"FROM therickys93/", m["os"], m["language"], "\n", "ADD . /app\nWORKDIR /app\nCMD [\"bash\", \"test.sh\"]", "\n"}
    d1 = []byte(strings.Join(s, ""))
    err = ioutil.WriteFile("./Dockerfile", d1, 0644) 
    check(err)

    // create the docker-compose.yml file
    s = []string{"sut:\n", "  build: .\n", "  dockerfile: Dockerfile", "\n"}
    d1 = []byte(strings.Join(s, ""))
    err = ioutil.WriteFile("./docker-compose.yml", d1, 0644)
    check(err)

    // create the onlytest.sh file
    // run the script onlytest.sh
    // remove all the files
}