package score

import (
	"fmt"
	"os"
	"path/filepath"
)

type Rules struct {
	rules []Rule
}

func ParseRules(path string) (*Rules, error) {
	// check path is a dir
	if !checkIsDir(path) {
		return nil, fmt.Errorf("path is not a dir")
	}

	rules := Rules{
		rules: make([]Rule, 0),
	}

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 判断是否为文件，并且判断文件名是否以 ".rule" 结尾
		if !info.IsDir() && filepath.Ext(path) == ".rule" {
			rule, err := ParseRule(path)
			if err != nil {
				return err
			}
			rules.rules = append(rules.rules, *rule)
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to walk through the dir: %w", err)
	}

	return &rules, nil
}

func checkIsDir(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}

	// 判断是否为目录
	if fileInfo.IsDir() {
		return true
	}
	return false
}

type Score interface {
	Calculate()
	Match()
}
