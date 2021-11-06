package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"os/signal"
	"path"
	"sort"
	"strings"
	"terminal-tool/menu"
	"terminal-tool/utils"
)

var dev = false
var env = "work"

var menus []menu.Menu

func main() {
	readFile()

	catchCtrlC()
	utils.TtyOn()
	defer utils.TtyOff()

	for {
		fmt.Println()
		printMainMenu()
		fmt.Print("[")
		var inp byte
		var input int
		inp = utils.GetChar()
		input, err := utils.CharMapLetters(inp)
		if err == nil {
			if input >= len(menus) {
				fmt.Println("~]")
				return
			}
			fmt.Printf("%c]\n", inp)
			fmt.Println()
			menuu := &menus[input]
			printMenu(*menuu)
			fmt.Print("[")

			inp = utils.GetChar()
			if input, err = utils.CharMapLetters(inp); err == nil {
				if input >= len(menuu.Items) {
					fmt.Println("~]")
					return
				}
				fmt.Printf("%c]\n", inp)
				utils.ClipWrite(menuu.Items[input].Content)
				menuu.Items[input].Usage += 1
				writeFile()
				return
			} else {
				if inp == '*' {
					fmt.Print("*] ")
					name := readName()
					if name != "" {
						menuItem := menu.MenuItem{
							Title:   name,
							Content: utils.ClipRead(),
							Usage:   0,
						}
						menuu.Items = append(menuu.Items, menuItem)
						writeFile()
					} else {
						fmt.Println()
						return
					}
				} else {
					fmt.Println("~]")
					return
				}
			}
			fmt.Println()
		} else {
			fmt.Println("~]")
			return
		}
	}
}

func readFile() {
	items, err := os.ReadFile(path.Join(dataPath(), "items.yml"))
	if err != nil {
		fmt.Println(err)
	}

	err = yaml.Unmarshal(items, &menus)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}

func writeFile() {
	for _, mennu := range menus {
		sort.Slice(mennu.Items, func(i, j int) bool {
			ui := mennu.Items[i].Usage
			uj := mennu.Items[j].Usage
			if ui == uj {
				return mennu.Items[i].Title < mennu.Items[j].Title
			} else {
				return ui > uj
			}
		})
	}
	yml, err := yaml.Marshal(menus)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(path.Join(dataPath(), "items.yml"), yml, 0644)
	if err != nil {
		panic(err)
	}
}

func readName() string {
	var name string
	for {
		inp := utils.GetChar()
		if inp == '\n' {
			break
		} else if inp == 27 {
			return ""
		}
		name += string(inp)
		fmt.Print(string(inp))
	}
	return name
}

func printMainMenu() {
	for i, menu := range menus {
		base, err := utils.Base26(i)
		if err != nil {
			panic("Not supported")
		}
		title := fmt.Sprintf("%-16.16s", menu.Title)
		title = utils.SetFgColor(title, 60, 170, 250)
		fmt.Printf("[%s] %s\n", base, title)
	}
	quit := utils.SetFgColor("Quit", 250, 170, 60)
	fmt.Printf("[~] %s\n", quit)
}

func printMenu(menu menu.Menu) {
	maxWidth := len(menu.Items[0].Content)
	for _, item := range menu.Items[1:] {
		contentLength := len(item.Content)
		if maxWidth < contentLength {
			maxWidth = contentLength
		}
	}
	if maxWidth > 90 {
		maxWidth = 90
	}
	for i, item := range menu.Items {
		base, err := utils.Base26(i)
		content := strings.Replace(item.Content, "\n", "\\n", -1)
		if err != nil {
			panic("Not supported")
		}
		title := fmt.Sprintf("%-16.16s", item.Title)
		title = utils.SetFgColor(title, 60, 170, 250)
		content = fmt.Sprintf("%-*.*s", maxWidth, maxWidth, content)
		content = utils.SetFgColor(content, 250, 230, 15)
		fmt.Printf("[%s] %s | %s | %d\n", base, title, content, item.Usage)
	}
	nuu := utils.SetFgColor("New", 250, 170, 60)
	quit := utils.SetFgColor("Quit", 250, 170, 60)
	fmt.Printf("[*] %s [~] %s\n", nuu, quit)
}

func dataPath() string {
	if dev {
		return path.Join(utils.GetWorkDir(), "env", env)
	} else {
		return utils.GetExeDir()
	}
}

// catchCtrlC tty settings should be reversed
func catchCtrlC() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		utils.TtyOff()
		fmt.Println()
		os.Exit(0)
	}()
}
