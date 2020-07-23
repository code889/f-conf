package fconf

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func is_line_ignore(line string) bool {

	ings := []string{";", "#"}
	for _, ig := range ings {
		if strings.HasPrefix(line, ig) {
			return true
		}
	}
	return false
}

type icf struct {
	cf map[string]map[string]string
}

func (i *icf) Load(path string) {

	i.cf = make(map[string]map[string]string)

	cf := i.cf

	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	var sec string
	var key, val string

	reader := bufio.NewReader(f)
	for {
		line, err := reader.ReadString('\n')
		if err != nil && line == "" {
			break
		}
		line = strings.TrimSpace(line)
		if line == "" || is_line_ignore(line) {
			continue
		}
		if strings.HasPrefix(line, "[") {
			sec = line[1 : len(line)-1]
			if sec == "" {
				panic("empty section")
			}
			_, exists := cf[sec]
			if !exists {
				cf[sec] = make(map[string]string)
			} else {
				panic(fmt.Sprintf("duplicate section: %v", sec))
			}

		} else {
			pos := strings.Index(line, "=")
			if pos == -1 {
				panic(fmt.Sprintf("error line: %v", line))
			}
			key = line[:pos]
			val = line[pos+1:]
			_, exists := cf[sec][key]
			if exists {
				panic(fmt.Sprintf("duplicate option: %v-%v", sec, key))
			} else {
				cf[sec][key] = val
			}
		}
	}
}

func (i *icf) Opt(section, option string) string {
	return i.cf[section][option]
}

func Main() {

	i := &icf{}
	i.Load("t.ini")
	fmt.Println(i.Opt("nihao", "xx"))
}

func Test() {

}
