package filter

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"svn-pre-commit/log"

	"github.com/axgle/mahonia"
)

type MetaFilter struct {
	regexpList       []*regexp.Regexp
	ignoreRegexpList []*regexp.Regexp
	suffix           string
}

func (m *MetaFilter) Exec() {
	var notExistFiles []string
	for i := 0; i < len(CommitFiles); i++ {
		f := CommitFiles[i]
		if strings.HasSuffix(f, m.suffix) {
			continue
		}
		valid := m.IsMatch(f, m.ignoreRegexpList)
		if valid {
			//fmt.Println("忽略:", f)
			continue
		}
		valid = m.IsMatch(f, m.regexpList)
		if !valid {
			continue
		}
		valid = m.IsExist(f)
		if !valid {
			notExistFiles = append(notExistFiles, f)
		}
	}
	if len(notExistFiles) != 0 {
		//d := mahonia.NewDecoder("utf8")
		encoder := mahonia.NewEncoder("gbk")
		errorStr := fmt.Sprint("缺少 '.meta文件' ", notExistFiles)
		errorStr = encoder.ConvertString(errorStr)
		//fmt.Println(errorStr)
		log.Error(errorStr)
		os.Stderr.WriteString(errorStr)
		os.Exit(3)
	}
}

func (m *MetaFilter) IsMatch(text string, regexpList []*regexp.Regexp) bool {
	for i := 0; i < len(regexpList); i++ {
		regexpStr := regexpList[i]
		valid := regexpStr.MatchString(text)
		if valid {
			return true
		}
	}
	return false
}

func (m *MetaFilter) IsExist(text string) bool {
	newText := fmt.Sprintf("%s%s", text, m.suffix)
	for _, e := range CommitFiles {
		if e == newText {
			return true
		}
	}
	return false
}

func MetaHandler(config StringMap) {
	regexpList, ignoreRegexpList := ConfigToRegexp(config)
	mf := &MetaFilter{regexpList, ignoreRegexpList, ".meta"}
	mf.Exec()
}
