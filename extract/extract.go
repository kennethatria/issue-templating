package extract

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

type extract struct {
	input  io.Reader
	output io.Writer
}

type option func(*extract) error

func WithInput(input io.Reader) option {
	return func(e *extract) error {
		if input == nil {
			return errors.New("nil input reader")
		}
		e.input = input
		return nil
	}
}

func WithInputFromArgs(args []string) option {
	return func(e *extract) error {
		if len(args) < 1 {
			return nil
		}

		f, err := os.Open(args[0])
		if err != nil {
			return err
		}
		e.input = f
		return nil
	}
}

func WithOutput(output io.Writer) option {
	return func(e *extract) error {
		if output == nil {
			return errors.New("nil output writer")
		}
		e.output = output
		return nil
	}
}

func NewExtract(opts ...option) (*extract, error) {

	e := &extract{
		output: os.Stdout,
		input:  os.Stdin,
	}

	for _, opt := range opts {
		err := opt(e)
		if err != nil {
			return nil, err
		}
	}
	return e, nil
}

func (e *extract) GetEnvironment() string {

	stormEnvironment := ""
	input := bufio.NewScanner(e.input)
	for input.Scan() {
		returnValue := extractEnvironment(input.Text())
		if len(returnValue) > 1 {
			stormEnvironment = returnValue
		}
	}
	return stormEnvironment
}

func (e *extract) GetVersion() string {

	stormVersion := ""
	input := bufio.NewScanner(e.input)
	for input.Scan() {
		//fmt.Println(input.Text())
		returnValue := extractVersion(input.Text())
		//fmt.Println(returnValue)
		if len(returnValue) > 1 {
			stormVersion = returnValue
		}
	}
	return stormVersion
}

func extractEnvironment(input string) string {

	environmentPattern := regexp.MustCompile(`(ite|ute|cae)`)

	environmentMatch := environmentPattern.FindStringSubmatch(input)
	environment := ""
	if len(environmentMatch) > 1 {
		environment = strings.TrimSpace(environmentMatch[1])
	}
	return environment
}

func extractVersion(input string) string {

	versionPattern := regexp.MustCompile(`\s*\s*(v[\d.]+)`)

	versionMatch := versionPattern.FindStringSubmatch(input)
	version := ""
	if len(versionMatch) > 1 {
		version = strings.TrimSpace(versionMatch[1])
	}
	return version
}

func Main() int {

	versionMode := flag.Bool("version", false, "return version , not environment")

	flag.Parse()

	e, err := NewExtract(
		WithInputFromArgs(flag.Args()),
	)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	if *versionMode {
		fmt.Println(e.GetVersion())
	} else {
		fmt.Println(e.GetEnvironment())
	}
	return 0
}
