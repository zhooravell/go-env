package env

// A functions to load environment variables from `.env`.

import (
	"bufio"
	"errors"
	"log"
	"os"
	"regexp"
	"strings"
)

const (
	defaultFileName = ".env"
	separatorChar   = "="
	commentChar     = "#"
	exportPrefix    = "export"
)

var (
	ErrNoSeparator = errors.New("no separator")
	ErrSeparate    = errors.New("can't separate key from value")
)

//Load environment variables from files.
//If the variable already exists, its value will not change.
func Load(fileNames ...string) error {
	for _, filename := range fileNamesOrDefault(fileNames) {
		if err := loadVariablesFromFile(filename, false); err != nil {
			return err
		}
	}

	return nil
}

//LoadWithOverriding environment variables from files.
//If the variable already exists, its value will change.
func LoadWithOverriding(fileNames ...string) error {
	for _, filename := range fileNamesOrDefault(fileNames) {
		if err := loadVariablesFromFile(filename, true); err != nil {
			return err
		}
	}

	return nil
}

//loadVariablesFromFile parse & set environment variables from file.
//If overload is TRUE value of exists variable will change
func loadVariablesFromFile(fileName string, overload bool) error {
	file, err := os.Open(fileName)

	if err != nil {
		return err
	}

	defer func() {
		if err := file.Close(); err != nil {
			log.Println(err)
		}
	}()

	scanner := bufio.NewScanner(file)

	currentEnv := map[string]bool{}
	for _, rawEnvLine := range os.Environ() {
		currentEnv[strings.Split(rawEnvLine, "=")[0]] = true
	}

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		// ignore line if empty or commented
		if strings.HasPrefix(line, commentChar) || len(line) == 0 {
			continue
		}

		k, v, err := parseLine(scanner.Text())

		if err != nil {
			return err
		}

		if _, ok := currentEnv[k]; ok && !overload {
			continue
		}

		if err := os.Setenv(k, v); err != nil {
			return err
		}
	}

	if err = scanner.Err(); err != nil {
		return err
	}

	return nil
}

//parseLine parse line to key value
func parseLine(line string) (string, string, error) {
	if !strings.Contains(line, separatorChar) {
		return "", "", ErrNoSeparator
	}

	splitString := strings.SplitN(line, separatorChar, 2)

	if len(splitString) != 2 {
		return "", "", ErrSeparate
	}

	k := splitString[0]

	if strings.HasPrefix(k, exportPrefix) {
		k = strings.TrimPrefix(k, exportPrefix)
	}

	return strings.TrimSpace(k), parseValue(splitString[1]), nil
}

//parseValue parse variable value
func parseValue(value string) string {
	singleQuotes := regexp.MustCompile(`\A'(.*)'\z`).FindStringSubmatch(value)
	doubleQuotes := regexp.MustCompile(`\A"(.*)"\z`).FindStringSubmatch(value)

	if singleQuotes != nil || doubleQuotes != nil {
		value = value[1 : len(value)-1] // pull the quotes off the edges
	}

	if doubleQuotes != nil {
		value = regexp.MustCompile(`\\.`).ReplaceAllStringFunc(value, func(match string) string {
			switch strings.TrimPrefix(match, `\`) {
			case "n":
				return "\n"
			case "r":
				return "\r"
			default:
				return match
			}
		})

		value = regexp.MustCompile(`\\([^$])`).ReplaceAllString(value, "$1") // unescape characters
	}

	return value
}

//fileNamesOrDefault if slice is empty return slice with default filename
func fileNamesOrDefault(fileNames []string) []string {
	if len(fileNames) == 0 {
		fileNames = append(fileNames, defaultFileName)
	}

	return fileNames
}
