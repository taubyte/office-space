package issue

import (
	"errors"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"

	"github.com/taubyte/office-space/runtime"
)

func getBranchPrefix(ctx *runtime.Context) (branchPrefix string, err error) {
	branchPrefix = ctx.Args().First()
	if len(branchPrefix) == 0 {
		err = fmt.Errorf("Must provide a branch prefix, ex: `asd issue <prefix>`")
	}

	return
}

/*
	Temporary Fix:
	Gets the git repo config file relative to this path, reads the file, and extracts the repo url
*/
func getConfigRepoUrl() (string, error){
    gitrepoconfiglocation := "../.git/config"
    
    data, err := ioutil.ReadFile(gitrepoconfiglocation)
    if err != nil {
      return "", err
    }

    regexpression := `(^?url).*` 
    urlLineInFile,err := regexp.Compile(regexpression) 
    if err != nil {
      return "", err
    }
    
    urlmatch := urlLineInFile.FindStringSubmatch(string(data))
    if urlmatch != nil {   
      url := strings.Split(urlmatch[0],"url =")[1]
      return url, nil;
    }else{   
      return "", errors.New("No github URL found in the git config. Most likely not a repository.")
    }
}