package main

import (
	"fmt"
	"os/exec"
)

func test() {
	//c1 := exec.Command("echo", "aaa")
	//c2 := exec.Command("/mnt/c/Windows/System32/clip.exe")
	//c2.Stdin, _ = c1.StdoutPipe()
	//c2.Stdout = os.Stdout
	//_ = c2.Start()
	//_ = c1.Run()
	//_ = c2.Wait()

	//cmd := "cat /proc/cpuinfo | egrep '^model name' | uniq | awk '{print substr($0, index($0,$4))}'"
	cmd := "echo -n aaa | clip.exe"
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		fmt.Sprintf("Failed to execute command: %s", cmd)
	}
	fmt.Println(string(out))
}
