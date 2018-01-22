package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os/exec"
)

type Config struct {
	Input   string
	Output  string
	Command string
}

type Result struct {
	Host   string
	IP     string
	Output string
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	p := flag.String("config", "", "path for configuration file")
	flag.Parse()

	if *p == "" {
		log.Fatal("configuration file path not provided")
	}

	c, err := config(*p)
	check(err)

	df, err := ioutil.ReadFile(c.Input)
	check(err)
	d := make(map[string]string)
	err = json.Unmarshal([]byte(df), &d)
	check(err)

	results := []Result{}
	for h, ip := range d {
		out, err := exec.Command(c.Command, h).Output()
		check(err)
		res := Result{Host: h, IP: ip, Output: string(out)}
		results = append(results, res)
		log.Println(results)
		log.Println()
	}

	b, err := json.Marshal(results)
	check(err)
	err = ioutil.WriteFile(c.Output, b, 0644)
	check(err)

	log.Println(">> written to file " + c.Output)
}

func config(p string) (Config, error) {
	cf, err := ioutil.ReadFile(p)
	check(err)
	c := Config{}
	err = json.Unmarshal([]byte(cf), &c)
	check(err)

	return c, err
}
