package roosterconf

import (
	"reflect"
	"strings"
	"testing"
)

func TestNew(t *testing.T) {

	t.Run("valid YAML", func(t *testing.T) {
		exampleYAML := `- repo: github.com/foo/bar
  type: git
  extensions:
    - docx
    - doc
    - xlsx
  out: foo_bar
- repo: github.com/baz/qux
  type: svn
  extensions:
    - .docx
    - .doc
    - xlsx
  out: baz_qux
- repo: github.com/quux/corge
  type: hg
  extensions:
    - md
`

		want := []Set{
			{
				RepoURL:    "github.com/foo/bar",
				VCS:        "git",
				Extensions: []string{".docx", ".doc", ".xlsx"},
				OutputDir:  "foo_bar",
			},
			{
				RepoURL:    "github.com/baz/qux",
				VCS:        "svn",
				Extensions: []string{".docx", ".doc", ".xlsx"},
				OutputDir:  "baz_qux",
			},
			{
				RepoURL:    "github.com/quux/corge",
				VCS:        "hg",
				Extensions: []string{".md"},
				OutputDir:  "rooster_output_0001",
			},
		}

		tRepos, err := New(strings.NewReader(exampleYAML))
		if err != nil {
			t.Fatalf("err != nil, New() == %s", err)
		}
		if len(tRepos) != len(want) {
			t.Fatalf("len(want) != len(got): %d != %d", len(want), len(tRepos))
		}
		for i, repo := range tRepos {
			if !reflect.DeepEqual(repo, want[i]) {
				t.Fatalf("repo got != repo want: %+v != %+v", repo, want[i])
			}
		}
	})

	t.Run("valid YAML: assume git", func(t *testing.T) {
		exampleYAML := `- repo: github.com/foo/bar
  extensions:
    - docx
    - doc
    - xlsx
  out: foo_bar
- repo: github.com/baz/qux
  type: svn
  extensions:
    - .docx
    - .doc
    - xlsx
  out: baz_qux
- repo: github.com/quux/corge
  type: hg
  extensions:
    - md
`

		want := []Set{
			{
				RepoURL:    "github.com/foo/bar",
				VCS:        "git",
				Extensions: []string{".docx", ".doc", ".xlsx"},
				OutputDir:  "foo_bar",
			},
			{
				RepoURL:    "github.com/baz/qux",
				VCS:        "svn",
				Extensions: []string{".docx", ".doc", ".xlsx"},
				OutputDir:  "baz_qux",
			},
			{
				RepoURL:    "github.com/quux/corge",
				VCS:        "hg",
				Extensions: []string{".md"},
				OutputDir:  "rooster_output_0001",
			},
		}

		tRepos, err := New(strings.NewReader(exampleYAML))
		if err != nil {
			t.Fatalf("err != nil, New() == %s", err)
		}
		if len(tRepos) != len(want) {
			t.Fatalf("len(want) != len(got): %d != %d", len(want), len(tRepos))
		}
		for i, repo := range tRepos {
			if !reflect.DeepEqual(repo, want[i]) {
				t.Fatalf("repo got != repo want: %+v != %+v", repo, want[i])
			}
		}
	})

	t.Run("invalid YAML: missing extensions", func(t *testing.T) {
		exampleYAML := `- repo: github.com/foo/bar
  type: git
  out: foo_bar
`
		tRepos, err := New(strings.NewReader(exampleYAML))
		if err == nil {
			t.Fatalf("err == nil, want err != nil")
		}
		if tRepos != nil {
			t.Fatalf("repos != nil, got: %+v", tRepos)
		}
	})

	t.Run("invalid YAML: missing repo", func(t *testing.T) {
		exampleYAML := `- type: git
  extensions:
    - docx
    - doc
    - xlsx
  out: foo_bar`
		tRepos, err := New(strings.NewReader(exampleYAML))
		if err == nil {
			t.Fatalf("err == nil, want err != nil")
		}
		if tRepos != nil {
			t.Fatalf("repos != nil, got: %+v", tRepos)
		}
	})

	t.Run("invalid YAML: unsupported vcs", func(t *testing.T) {
		exampleYAML := `- type: unsupported
  repo: github.com/foo/bar
  extensions:
    - docx
    - doc
    - xlsx
  out: foo_bar`
		tRepos, err := New(strings.NewReader(exampleYAML))
		if err == nil {
			t.Fatalf("err == nil, want err != nil")
		}
		if tRepos != nil {
			t.Fatalf("repos != nil, got: %+v", tRepos)
		}
	})

	t.Run("invalid YAML: bad indentation", func(t *testing.T) {
		exampleYAML := `- type: unsupported
	repo: github.com/foo/bar
  extensions:
    - docx
    - doc
    - xlsx
  out: foo_bar`
		tRepos, err := New(strings.NewReader(exampleYAML))
		if err == nil {
			t.Fatalf("err == nil, want err != nil")
		}
		if tRepos != nil {
			t.Fatalf("repos != nil, got: %+v", tRepos)
		}
	})

}
