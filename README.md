# Rooster

![LogoMakr_7ZIRlg](https://user-images.githubusercontent.com/13544676/86405652-91fa2f80-bc66-11ea-8543-f56ab909bb9d.png)

Rooster is a file filter for version control systems.

## Installation

1. Clone the repo

`git clone https://github.com/SheetJS/rooster.git rooster`

2. Run go build

```
cd rooster
go build
```

## Usage

In order to use `rooster` a config file must be provided:

```
Usage of ./rooster:
  -config string
        Rooster's config file (default ".rooster.yaml")
```

The layout of the .rooster.yaml is as follows:

```yaml
- repo: https://github.com/foo/bar
  extensions:
    - docx
    - doc
    - xlsx
  out: foo_bar

- repo: https://github.com/baz/qux
  type: svn
  extensions:
    - .docx
    - .doc
    - xlsx
  out: baz_qux

- repo: https://github.com/quux/corge
  type: hg
  extensions:
    - md
```

Each object must have the following keys:

- `repo`: The repository's URL
- `extension`: An array of extensions to filter for

Note the following:

> 1. The leading `.` is optional
> 2. Rooster will be strict and will NOT infer other file extension given another. (i.e. using .txt will only grab .txt files and NOT .text files)

The following keys are optional:

- `type`: The version control system to use (supported options: git,svn,hg) [default: git]
- `out`: The output directory to save the filtered files to (default: rooster_output_xxxx)

### Example Output

Once `rooster` is done each filtered repository will have a output directory similar to the one below:

```
rooster_output_xxx
├── doc
├── docx
└── txt
```

Where all docx files will be in the `docx` directory and all the `doc` files will be in the `doc` folder, etc, etc.

### Dependencies

You'll need to have your vcs provider installed on your system.
An exception to this is `git` as `rooster` uses a go-based implementation of git.

The currently supported vcs'es are:

- Git
- Subversion
- Mercurial

## Credits

### Logo

Created at [LogoMakr.com](https://www.LogoMakr.com)

### go-git

For the awesome git [library](https://github.com/go-git/go-git) written in pure go.

### go-yaml

For the awesome [YAML parser](https://github.com/go-yaml/yaml).
