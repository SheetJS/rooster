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

From `./rooster --help`:

```
Usage of ./rooster:
  -ext string
        The extension(s) to filter in a comma seperated list
  -out string
        The directory to store the output (default "rooster_output")
  -repo string
        The repo to grab
  -type string
        The VCS to use [supported options: git,svn,hg] (default "git")
```

### Example

Running `rooster -repo https://github.com/foo/bar -ext ".doc,docx,.txt"` will clone `https://github.com/foo/bar` and extract all files ending the `doc, docx, and txt` extensions. Two things of note:

1. The leading `.` is optional
2. Rooster will be strict and will NOT infer other file extension given another. (i.e. using .txt will only grab .txt files and NOT .text files)

Since the `-out` was omitted `rooster` will create a directory named `rooster_output` and once the program is done it will result in the following file structure:

```
rooster_output
├── doc
├── docx
└── txt
```

Where all docx files will be in the `docx` directory and all the `doc` files will be in the `doc` folder, etc, etc.

If the `-out` flag was used then `rooster_output` would be substituted for the argument passed to `-out`.

### Dependencies

You'll need to have your vcs provider installed on your system.

The currently supported vcs'es are:

- Git
- Subversion
- Mercurial

## Credits

### Logo

Created at [LogoMakr.com](https://www.LogoMakr.com)
