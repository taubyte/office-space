package issue

/*
Retrieves the remote url from the git repo config
*/
// func getConfigRepoUrl() (string, error) {
// 	gitrepoconfiglocation := "../.git/config"

// 	data, err := ioutil.ReadFile(gitrepoconfiglocation)
// 	if err != nil {
// 		return "", err
// 	}

// 	regexpression := `(^?url).*`
// 	urlLineInFile, err := regexp.Compile(regexpression)
// 	if err != nil {
// 		return "", err
// 	}

// 	urlmatch := urlLineInFile.FindStringSubmatch(string(data))
// 	if urlmatch != nil {
// 		url := strings.Split(urlmatch[0], "url =")[1]
// 		return url, nil
// 	} else {
// 		return "", errors.New("No github URL found in the git config. Most likely not a repository.")
// 	}
// }
