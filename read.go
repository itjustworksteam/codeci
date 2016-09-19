package main 

import (
	"fmt"
	"io/ioutil"
	"gopkg.in/yaml.v2"
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
	fmt.Printf("%v\n", m["script"]);
}