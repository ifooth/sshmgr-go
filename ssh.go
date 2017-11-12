package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strconv"
)

const (
	defaultStor = ".ssh_servers"
	defaultCmd  = "ssh"
)

func readData() ([][]string, error) {
	var path string
	path, _ = filepath.Abs(defaultStor)
	csvfile, err := os.Open(path)
	if err != nil {
		usr, err := user.Current()
		if err != nil {
			return nil, err
		}
		path = filepath.Join(usr.HomeDir, defaultStor)
		csvfile, err = os.Open(path)
		if err != nil {
			return nil, err
		}
	}

	defer csvfile.Close()

	reader := csv.NewReader(csvfile)
	fields, err := reader.ReadAll()

	return fields, err
}

func run(cmdName string, arg ...string) {
	cmd := exec.Command(cmdName, arg...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Failed to start ssh, %s\n", err.Error())
		os.Exit(1)
	}
}

func main() {
	stor, err := readData()
	if err != nil {
		fmt.Printf("Failed to read data, %s\n", err.Error())
		os.Exit(1)
	}

	label := make(map[string][]string)
	for _, line := range stor {
		host := line[0]
		for _, j := range line[1:] {
			_, ok := label[j]
			if ok == true {
				label[j] = append(label[j], host)
			} else {
				label[j] = []string{host}

			}
		}

	}
	index := 1
	label_map := make(map[int]string)
	for k := range label {
		label_map[index] = k
		fmt.Printf("[%d] -> %s\n", index, k)
		index++
	}
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("choice label:")
	var choice_label int
	for scanner.Scan() {
		choice, _ := strconv.ParseUint(scanner.Text(), 10, 0)
		choice_label = int(choice)
		break
	}

	index = 1
	host_map := make(map[int]string)
	for _, k := range label[label_map[choice_label]] {
		host_map[index] = k
		fmt.Printf("[%d] -> %s\n", index, k)
		index++
	}

	fmt.Print("choice host:")
	var choice_host int
	for scanner.Scan() {
		choice, _ := strconv.ParseUint(scanner.Text(), 10, 0)
		choice_host = int(choice)
		break
	}

	run(defaultCmd, host_map[choice_host])
}
