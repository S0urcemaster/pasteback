package menu

type Menu struct {
	Title string     `yaml:"title"`
	Items []MenuItem `yaml:"items"`
}

type MenuItem struct {
	Title   string `yaml:"title"`
	Content string `yaml:"content"`
	Usage   int    `yaml:"usage"`
}
