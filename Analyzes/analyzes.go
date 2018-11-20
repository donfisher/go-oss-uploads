package Analyzes

import (
	"os"
	"regexp"
)

type Analyzes struct {
	
}

var pathRe = regexp.MustCompile(`.apk|.zip`)


func (a *Analyzes) Check(path string) (bool bool,err error){
	_, err = os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func (a *Analyzes) IsFile(path string) (bool bool) {
	compile := pathRe
	s := compile.FindString(path)
	if s == ""{
		return false
	}
	return true
}


