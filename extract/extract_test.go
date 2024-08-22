package extract_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/kennethatria/extract"
	"github.com/rogpeppe/go-internal/testscript"
)

func TestInputAsStdin(t *testing.T) {
	t.Parallel()

	type testCase struct {
		a, want string
	}

	testCases := []testCase{
		{a: "ute\nv1.3.4\nother text", want: "ute"},
		{a: "ite\nv1.3.4\nother text", want: "ite"},
	}

	for _, tc := range testCases {

		inputBuf := bytes.NewBufferString(tc.a)
		e, err := extract.NewExtract(
			extract.WithInput(inputBuf),
		)
		if err != nil {
			t.Fatal(err)
		}
		want := tc.want
		got := e.GetEnvironment()
		if want != got {
			t.Errorf("want %s, got %s", want, got)
		}

	}

}

func TestInputForIteValue(t *testing.T) {
	t.Parallel()
	args := []string{"testdata/gh_issue_response_two.txt"}

	e, err := extract.NewExtract(
		extract.WithInputFromArgs(args),
	)

	if err != nil {
		t.Fatal(err)
	}
	want := "ite"
	got := e.GetEnvironment()
	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestInputForCaeValue(t *testing.T) {
	t.Parallel()
	args := []string{"testdata/gh_issue_response_one.txt"}

	e, err := extract.NewExtract(
		extract.WithInputFromArgs(args),
	)

	if err != nil {
		t.Fatal(err)
	}
	want := "cae"
	got := e.GetEnvironment()
	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestInputForValidVersion(t *testing.T) {
	t.Parallel()
	args := []string{"testdata/gh_issue_response_one.txt"}

	e, err := extract.NewExtract(
		extract.WithInputFromArgs(args),
	)

	if err != nil {
		t.Fatal(err)
	}
	want := "v4.15.6"
	got := e.GetVersion()
	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestInputForInvalidVersion(t *testing.T) {
	t.Parallel()
	args := []string{"testdata/gh_issue_response_invalid.txt"}

	e, err := extract.NewExtract(
		extract.WithInputFromArgs(args),
	)

	if err != nil {
		t.Fatal(err)
	}
	want := ""
	got := e.GetVersion()
	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}

func InputForEmptyArgument(t *testing.T) {
	t.Parallel()
	e, err := extract.NewExtract(
		extract.WithInputFromArgs([]string{}),
	)
	if err != nil {
		t.Fatal(err)
	}
	want := ""
	got := e.GetVersion()
	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestMain(m *testing.M) {
	os.Exit(testscript.RunMain(m, map[string]func() int{
		"extract": extract.Main,
	}))
}

func Test(t *testing.T) {
	t.Parallel()
	testscript.Run(t, testscript.Params{
		Dir: "testdata/script",
	})
}
