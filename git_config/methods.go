package git_config

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/bigkevmcd/go-configparser"
)

/*
	Open method will take path string that's passed in and initialize and return a configuration parser instance or an error
*/
func(w *gitConfig) Open(path string) (*configparser.ConfigParser, error){
	if len(path) != 0 { //If a path was provided as expected
		//Attempt to create a new configuration file parser
		f,err := configparser.NewConfigParserFromFile(path)
		if err != nil { //If we encounter an error
			return nil, err //Return nil and the error
		}
		return f, nil //Otherwise return the configparser and nil for the error
	}else{ //Otherwise 
		//Return nil for the parser and an error explaining the issue with the path
		return nil, errors.New("You must include a valid path when calling Open() method on a Github Configuration instance.")
	}
}

//GitRepo Modify Methods

//Github Config File Core Section Fields
func (rc *gitConfig) ModifyRepoConfigFormatVersion(version int){
	rc.repositoryformatversion = version
}

func (rc *gitConfig) ModifyRepoConfigFilemode(filemode bool){
	rc.filemode = filemode
}

func (rc *gitConfig) ModifyRepoConfigBare(bare bool){
	rc.bare = bare
}

func (rc *gitConfig) ModifyRepoConfigLogAllRefUpdates(logallrefupdates bool){
	rc.logallrefupdates = logallrefupdates
}

func (rc *gitConfig) ModifyRepoConfigIgnoreCase(ignorecase bool){
	rc.ignorecase = ignorecase
}

func (rc *gitConfig) ModifyRepoConfigPrecomposeUnicode(precomposeunicode bool){
	rc.precomposeunicode = precomposeunicode
}

//Remotes sections
func (rc *gitConfig) ModifyRepoConfigRemotes(remotes []GithubRemote){
	rc.remotes = remotes
}

//Branch sections
func (rc *gitConfig) ModifyRepoConfigBranches(branches []GithubBranch){
	rc.branches = branches
}

//Github Remote Modify Methods
func (gr *GithubRemote) ModifyGithubRemoteUrl(url  string){
	gr.url = url
}

func (gr *GithubRemote) ModifyGithubRemoteFetch(fetch string){
	gr.fetch = fetch
}

//Github Branch Modify Methods
func (gb *GithubBranch) ModifyGithubBranchRemote(remote string){
	gb.remote = remote
}

func (gb *GithubBranch) ModifyGithubBranchMerge(merge string){
	gb.merge = merge
}


//Config File Data Extraction Methods

/*
	Extract the formatversion value from the core of the github configuration file
*/
func extractRepoFormatVersion(repoConfigParsed *gitConfig, p *configparser.ConfigParser) error {
	//Get the repositoryformatversion from the core section of the config file
	repoConfigFormatVersionStr, err := p.Get("core", "repositoryformatversion")
	if err != nil {
		return err
	}

	//Convert it to an integer
	repoConfigFormatVersionInt,err := strconv.Atoi(repoConfigFormatVersionStr)
	if err != nil {
		return err
	}

	//Add it to the parsed configuration struct
	repoConfigParsed.ModifyRepoConfigFormatVersion(repoConfigFormatVersionInt)

	return nil
}

/*
	Extract the filemode value from the core section of the github configuration file
*/
func extractRepoFilemode(repoConfigParsed *gitConfig, p *configparser.ConfigParser) error {
	//Get the filemode from the core section of the config file
	repoConfigFileModeStr, err := p.Get("core","filemode")
	if err != nil {
		return err
	}

	//Convert it to a boolean
	repoConfigFileModeBool := strings.Contains(repoConfigFileModeStr,"true")
	
	//Add it to the parsed configuration struct
	repoConfigParsed.ModifyRepoConfigFilemode(repoConfigFileModeBool)

	return nil
}

/*
	Extract the bare value from the core section of the github configuration file
*/
func extractRepoBare(repoConfigParsed *gitConfig, p *configparser.ConfigParser) error {
	//Get the bare value from the core section of the config file
	repoBareStr, err := p.Get("core","bare")
	if err != nil {
		return err
	}

	//Convert it to a boolean
	repoBareBool := strings.Contains(repoBareStr, "true")

	//Add it to the parsed configuration struct
	repoConfigParsed.ModifyRepoConfigBare(repoBareBool)

	return nil
}

/*
	Extract the logallrefupdates value from the core section of the github configuration file
*/
func extractLogAllRefUpdates(repoConfigParsed *gitConfig, p *configparser.ConfigParser) error {
	//Get the logallrefupdates value from the core section of the config file
	repoLogRefAllUpdatesStr, err := p.Get("core","logrefallupdates")
	if err != nil {
		return err
	}

	//Convert it to a boolean
	repoLogRefAllUpdatesBool := strings.Contains(repoLogRefAllUpdatesStr,"true")

	//Add it to the parsed configuration struct
	repoConfigParsed.ModifyRepoConfigLogAllRefUpdates(repoLogRefAllUpdatesBool)

	return nil
}


/*
	Extract the ignorecase value from the core section of the github configuration file
*/
func extractIgnoreCase(repoConfigParsed *gitConfig, p *configparser.ConfigParser) error {
	//Get the ignorecase value from the core section of the config file
	repoIgnoreCaseStr, err := p.Get("core","ignorecase")
	if err != nil {
		return err
	}

	//Convert it to a boolean
	repoIgnoreCaseBool := strings.Contains(repoIgnoreCaseStr,"true")

	//Add it to the parsed configuration struct
	repoConfigParsed.ModifyRepoConfigIgnoreCase(repoIgnoreCaseBool)

	return nil
}

/*
	Extract the precomposeunicode value from the core section of the github configuration file
*/
func extractPrecomposeUnicode(repoConfigParsed *gitConfig, p *configparser.ConfigParser) error {
	//Get the precomposeunicode value from the core section of the config file
	repoPrecomposeUnicodeStr, err := p.Get("core","precomposeunicode")
	if err != nil {
		return err
	}

	//Convert it to a boolean
	repoPrecomposeUnicodeBool := strings.Contains(repoPrecomposeUnicodeStr,"true")

	//Add it to the parsed configuration struct
	repoConfigParsed.ModifyRepoConfigPrecomposeUnicode(repoPrecomposeUnicodeBool)

	return nil
}

/*
	TODO: Extract all Remotes and their values from the github configuration file
*/


/*
	TODO: Extract all the Branches and their values from the github configuration file
*/


/*
	Serialize the entire file at the given path into the gitConfig instance
*/
func serializeGithubRepoConfigFile(path string) (*gitConfig, error){
	//Relative location of the config file
	
	//Create an empty GithubRepoConfig object
	repoConfigParsed := &gitConfig{}

	//Call open method on it and pass the path, this will return a parser
	p, err := repoConfigParsed.Open(path)
	if err != nil {
		return nil, err
	}

	//Read the contents of the config file
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	//Display on terminal
	fmt.Println("====== Parsing Repository Config File ======")
	fmt.Println()
	fmt.Println(string(data))
	fmt.Println()
	fmt.Println("============================================")

	//Extract the values out of the config file
	
	//Extract the format version
	err = extractRepoFormatVersion(repoConfigParsed,p)
	if err != nil {
		return nil, err
	}

	//Extract filemode value from core section of config
	err = extractRepoFilemode(repoConfigParsed,p)
	if err != nil {
		return nil, err
	}

	//Extract the bare value out of the core section
	err = extractRepoBare(repoConfigParsed,p)
	if err != nil {
		return nil, err
	}

	//Extract the logallrefupdates value from the core section
	err = extractLogAllRefUpdates(repoConfigParsed,p)
	if err != nil {
		return nil, err
	}

	//Extract the ignorecase value from the core section
	err = extractIgnoreCase(repoConfigParsed,p)
	if err != nil {
		return nil, err
	}

	//TODO: Extract all the remotes and their values from their respective sections

	//TODO: Extract all the branches and their values from their respective sections

	//Return the serialized configuration file and nil for the error
	return repoConfigParsed, nil
	
}