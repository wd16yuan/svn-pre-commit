package filter

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/axgle/mahonia"
	"gopkg.in/yaml.v2"
)

type conf struct {
	Meta StringMap `yaml:"meta"`
}

func GetConfigInfo() (c *conf) {
	yamlFile, err := ioutil.ReadFile("check.yaml")
	if err != nil {
		fmt.Printf("读取文件失败: %v", err)
		os.Exit(0)
	}

	c = &conf{}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		fmt.Printf("解析文件失败: %v", err)
		os.Exit(0)
	}
	return
}

func GetCommitFiles() {
	fileInfo, _ := os.Stdin.Stat()
	//fmt.Println(fileInfo.Mode() & os.ModeNamedPipe)
	if (fileInfo.Mode() & os.ModeNamedPipe) != os.ModeNamedPipe {
		fmt.Printf("该命令用于处理管道输入")
		os.Exit(1)
	}
	// 管道输入读取
	//f, _ := os.Open("abc.log")
	decoder := mahonia.NewDecoder("gbk")
	bytes, _ := ioutil.ReadAll(decoder.NewReader(os.Stdin))

	//bytes, _ := ioutil.ReadAll(os.Stdin)

	//bytes, _ := ioutil.ReadAll(decoder.NewReader(f))
	//fmt.Println(bytes)
	text := string(bytes[:])

	//fmt.Println(text)
	lines := strings.Split(text, "\n")
	//fmt.Println(len(lines))
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		//只验证新增与删除的文件
		if !strings.HasPrefix(line, "D") && !strings.HasPrefix(line, "A") {
			continue
		}
		info := strings.Fields(strings.TrimSpace(line))
		if len(info) < 1 {
			continue
		}
		file_name := info[len(info)-1]
		CommitFiles = append(CommitFiles, file_name)
	}

}
