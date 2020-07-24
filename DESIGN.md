# Proposal: Rooster, a File Extractor for Git Repositories

Author: Jorge Henriquez (contact@jorgehenriquez.dev)

Last updated: Mon Jun 29 2020

## Abstract

This proposal will add the ability to filter git repositories for certain file extensions.

## Background

Since testing the wordJS project will rely on testing various file formats, it is efficient to reuse test files that have already been written by other members of the open source community.

This tool will make it easier to do so.

## Proposal

Add a cross platform utility to extract files that match a certain extension from git repositories.

## Rationale

The alternate method of doing this is to do so manually, this would be inefficient and would require a tedious amount of human hours to do.

Using `rooster`, we can clone many repositories at a time and get all the desired content.

## Implementation

### Cross Platform Compatibility

Since this utility will aim to be cross platform this eliminates the use of UNIX programs like `bash` and `find`.

Instead we will use Golang as it is easy to develop with and can ship binaries that are directly runnable by the end user without any other dependencies.

### Command Line Interface

Rooster will be executed in the following format

`rooster [git repo URL] [-ext ext1,ext2,ext3] (-out dirname)`

Where the `[]` brackets represent required parameters and the `()` represent optional parameters.

##### Example

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

### Getting the Repository

It is necessary to clone the entire git repo as there is no way to know where the files may be located within the repository.

To clone the repo we have two options:

1. Make a child process that directly invokes the `git clone` command.
2. We can use [libgit2](https://libgit2.org) to directly clone the repo within the same process.

Going with option 2 is indeed the "appropriate" way to do this.
But since this utility is rather small, the overhead of making a child process isn't too taxing on the system.
This also has the benefit using whatever credentials the user has stored with git for free without extra implementation.

#### The Filtering Algorithm

1.  Initialize a set `S`
2.  Initialize a map `M` that maps a file extension to an array of file paths.
3.  Add all given file extensions to the set `S`.
4.  Walk the given repo recursively
    1. Assign `F` to the current file.
    2. If `F`'s file extension in in the set `S`, then append `F.path` to `M[F.ext]`.
    3. Continue to the next file.

This algorithm then leaves `M` complete; we can then write `M`'s representation back to the disk.
