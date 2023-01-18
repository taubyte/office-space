# Office-space

## Introduction
  This is a CLI for managing your VSCode workspace


## Installation

### From source
```bash
# Clone the repository
$ git clone https://github.com/taubyte/office-space.git

# Change directory into the repository
$ cd office-space

# Install the dependencies
$ go mod tidy

# Build the binary into your GOPATH as `asd`, this enables the command to be ran from anywhere.
$ go build -o $GOPATH/bin/asd
```

### From go install
```bash
# Run the go install command
$ go install github.com/taubyte/office-space@latest

# Rename the binary to `asd`
$ mv $GOPATH/bin/office-space $GOPATH/bin/asd

# Verify the binary is installed
$ which asd
```

## Usage

```bash
# Enable your workspace by creating a `main.code-workspace` in your env["OFFICE_WORKSPACE_DIRECTORY"] directory
$ asd init
```

### Basic Multi-Repo Operations
```bash
$ asd <command>
```

#### Examples:
```bash
# Open the go.mod of every repo in the workspace.
$ asd code go.mod
```

```bash
# Checkout the master branch of every repo in the workspace.
$ asd git checkout master
```

```bash
# Pull the latest changes from every repo in the workspace.
$ asd git pull
```

### Go Workspace Operations
```bash
# Replace all workspace opened repositories. This will add all opened repositories to `go.work`
$ asd work
```

```bash
# Cleans/empties the relative `go.work`
$ asd work clean, c
```

```bash
# Deletes the relative `go.work`
$ asd work delete, d
```

```bash
# Builds workspace `./main.code-workspace` with replaces from `go.work`
$ asd work build, b
```

```bash
# Adds arg[0] to workspace and replaces `./main.code-workspace` in `go.work`
$ asd work add, a
$ asd work add <directory>
```

```bash
# Removes arg[0] from workspace and removes replace for `./main.code-workspace` from `go.work`
$ asd work remove, rm
$ asd work remove <directory>
```

```bash
# Removes replaces of a given package, removes from go.work, removes from ./main.code-workspace, and updates versions throughout to latest
$ asd work update, u
```

### Issue
```bash
# Looks for branches containing the provided prefix, checks out the branches, and adds the given repositories to go.work and the main workspace
$ asd issue TP-43
```

```bash
# Creates a directory with the provided prefix and clones the given repositories based on your `OFFICE_GIT_PREFIX` environment key.  It will also check out the branch and set the new `OFFICE_WORKSPACE_DIRECTORY` as an environment variable of the new workspace.
$ asd issue-clone TP-43
```


### Air

Runs air and appends arguments to `go test`, see: [Air Config](commands/air/.air.toml)


- [ ] Add support for `--ignore-dirs` to ignore changes in certain directories

```bash
# Livereload a test using air, see: https://github.com/cosmtrek/air
# Note: be sure you have air installed `which air` => ~/go/bin/air

# Simple Live reload of all relative tests
$ asd air

# Matcher Live reload equivalent to `go test -v --run <some-test>
$ asd air <some-test>

# Live reload of all tests below a given directory equivalent to `go test -v ./...`
$ asd air ./...
```


### Update

When merging PRs it has become an issue updating each version, removing replaces, etc.
asd update x will do the following:

 - Remove replaces of the package
 - Checkout and pull master on the package
 - Remove the package from the main.code-workspace file
 - `$ go get packageName@latest` for each item in workspace and tidy
 - `$ asd work`


```bash
$ asd update node

$ asd update node --dry
```

### Environment:
- `OFFICE_WORKSPACE_DIRECTORY`  Required, Set it in your ~/.bash_profile, ~/.profile or ~/.bashrc or with $ export OFFICE_WORKSPACE_DIRECTORY="path/to/package/directory"
- `OFFICE_WORKSPACE_NAME` Defaults "main"
- `OFFICE_WORKSPACE_EXT`  Defaults ".code-workspace"
- `OFFICE_GIT_PREFIX` Prefix used for cloning repositories with issue-clone, Ex "git@bitbucket.org/taubyte"

### Autocomplete:

Either run the following or add to .bashrc

```bash
$ eval "`asd autocomplete`"
```

## Future Operations
- `asd revert`: 
Revert the last command made, this will take tracking
what to do to revert each time a command is made.  For example, if you
just run a asd [command]  that will not be possible to revert.  But if you
run `asd work add <repository>` this could easily be reverted with `asd work remove <repository>`

Setting current workspace, although this will need to open a new shell.
- `asd ws <name>`
- `asd ws main` Sets workspace to main
- `asd ws current` Sets workspace to current

# Maintainers
 - Sam Stoltenberg @skelouse






git request-pull -p https://github.com/AldanisVigo/office-space