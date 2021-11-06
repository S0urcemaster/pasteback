package utils

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/atotto/clipboard"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
)

func Clip(s string) {
	_, err := exec.Command("bash", "-c", "echo -n \""+s+"\" | clip.exe").Output()
	if err != nil {
		fmt.Println(err)
	}
}

func ClipWrite(s string) {
	err := clipboard.WriteAll(s)
	if err != nil {
		panic(err)
	}
}

func ClipRead() string {
	s, err := clipboard.ReadAll()
	if err != nil {
		panic(err)
	}
	return s
}

func CharMap(char byte) (int, error) {
	if char > 47 && char < 58 {
		return int(char - 48), nil
	} else if char > 96 && char < 123 {
		return int(char - 87), nil
	}
	return 0, errors.New("Char not supported: " + strconv.FormatInt(int64(char), 10))
}

func CharMapLetters(char byte) (int, error) {
	if char > 96 && char < 123 {
		return int(char - 97), nil
	}
	return 0, errors.New("Char not supported: " + strconv.FormatInt(int64(char), 10))
}

func Base36(i int) (string, error) {
	if i < 0 {
		return "", errors.New("Not supported")
	}
	if i < 10 {
		return strconv.Itoa(i), nil
	} else if i < 36 {
		return string(byte(i) + 87), nil
	}
	return "", errors.New("Not supported")
}

func Base26(i int) (string, error) {
	if i < 0 {
		return "", errors.New("Not supported")
	}
	if i < 26 {
		return string(byte(i) + 97), nil
	}
	return "", errors.New("Not supported")
}

func Contains(haystack []string, needle string) bool {
	for _, straw := range haystack {
		if straw == needle {
			return true
		}
	}
	return false
}

func TtyOn() {
	//disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// do not display entered characters on the screen
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
}

func TtyOff() {
	exec.Command("stty", "-F", "/dev/tty", "-cbreak", "min", "1").Run()
	exec.Command("stty", "-F", "/dev/tty", "echo").Run()
}

func GetChar() byte {
	var b []byte = make([]byte, 1)
	_, err := os.Stdin.Read(b)
	if err != nil {
		return 0
	}
	return b[0]
}

func getDigit() (int, error) {
	var b = make([]byte, 1)
	for {
		_, err := os.Stdin.Read(b)
		if err != nil {
			return 0, err
		}
		i, err := strconv.Atoi(string(b))
		if err == nil {
			return i, nil
		}
	}
}

func GetNumber(digits int) (int, error) {
	var input1, err = getDigit()
	if err != nil {
		return 0, err
	}
	input2, err := getDigit()
	if err != nil {
		return 0, err
	}
	return input1*10 + input2, nil
}

func GetWorkDir() string {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	return path
}

func GetExeDir() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return filepath.Dir(ex)
}

func Keys(m map[string][]string) []string {
	keys := make([]string, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}

func GetNthItem(m map[string][]string, n int) []string {
	i := 0
	for _, v := range m {
		if i == n {
			return v
		}
		i++
	}
	return nil
}

func OsRun(command string) {
	cmd := exec.Command(command)
	cmd.Dir = "/mnt/c/Windows/System32"
	fmt.Println(cmd.String())
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return
	}
	fmt.Println("Result: " + out.String())
}
