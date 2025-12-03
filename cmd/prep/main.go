package main

import (
	_ "embed"
	"errors"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
	"text/template"
)

//go:embed day.tmpl
var dayTemplateS string
var DayTemplate = template.Must(template.New("day").Parse(dayTemplateS))

//go:embed solutions.tmpl
var solutionsTemplateS string
var SolutionsTemplate = template.Must(template.New("solutions").Parse(solutionsTemplateS))

//go:embed solutions_reg.tmpl
var solutionsRegistryTemplateS string
var SolutionsRegistryTemplate = template.Must(template.New("solutions_reg").Parse(solutionsRegistryTemplateS))

func main() {
	rootDir, year, days := parseArgs()

	yearDir := path.Join(rootDir, "solutions", fmt.Sprintf("year_%d", year))
	if err := os.MkdirAll(yearDir, 0755); err != nil {
		fmt.Printf("Failed to create year directory '%s'\nError: %v", yearDir, err)
		os.Exit(1)
	}

	outputTemplateToFile(
		SolutionsTemplate,
		path.Join(yearDir, "solutions.go"),
		struct{ Year int }{year},
		false,
	)

	for _, day := range days {
		outputTemplateToFile(
			DayTemplate,
			path.Join(yearDir, fmt.Sprintf("%d.go", day)),
			struct{ Year, Day int }{year, day},
			false,
		)
	}

	updateSolutionsRegistry(rootDir)

	fmt.Println("Done!")
}

func parseArgs() (rootDir string, year int, days []int) {
	if len(os.Args) != 4 {
		fmt.Printf("Usage: %s <root_dir> <year> <day range>\n", os.Args[0])
		os.Exit(1)
	}

	var err error

	rootDir = os.Args[1]

	if year, err = strconv.Atoi(os.Args[2]); err != nil {
		fmt.Printf("Invalid year (must be an integer): '%s'\n", os.Args[2])
		os.Exit(1)
	}

	dayRangeSplit := strings.Split(os.Args[3], "-")
	if len(dayRangeSplit) > 2 {
		fmt.Printf("Invalid day range: '%s'\n", os.Args[4])
		os.Exit(1)
	}

	if len(dayRangeSplit) == 2 {
		dayFrom, err := strconv.Atoi(dayRangeSplit[0])
		if err != nil {
			fmt.Printf("Invalid day from (must be an integer): '%s'\n", os.Args[2])
			os.Exit(1)
		}

		dayTo, err := strconv.Atoi(dayRangeSplit[1])
		if err != nil {
			fmt.Printf("Invalid day from (must be an integer): '%s'\n", os.Args[2])
			os.Exit(1)
		}

		for day := dayFrom; day <= dayTo; day++ {
			days = append(days, day)
		}
	} else {
		day, err := strconv.Atoi(dayRangeSplit[0])
		if err != nil {
			fmt.Printf("Invalid day from (must be an integer): '%s'\n", os.Args[2])
			os.Exit(1)
		}
		days = append(days, day)
	}

	return
}

func outputTemplateToFile(tpl *template.Template, path string, data any, override bool) {
	fileFlags := os.O_WRONLY | os.O_CREATE

	if override {
		fileFlags |= os.O_TRUNC
	} else {
		fileFlags |= os.O_EXCL
	}

	f, err := os.OpenFile(path, fileFlags, 0644)
	if err != nil {
		if errors.Is(err, os.ErrExist) {
			fmt.Printf("Skipping '%s' - already exists\n", path)
			return
		}

		fmt.Printf("Failed to create a file '%s' %d\nError: %v", path, err)
		os.Exit(1)
	}

	err = tpl.Execute(f, data)
	if err != nil {
		fmt.Printf("Failed to generate template for file '%s'\nError: %v", path, err)
		_ = f.Close()
		_ = os.Remove(path)
		os.Exit(1)
	}

	if err = f.Sync(); err != nil {
		fmt.Printf("Failed to sync generated file '%s'\nError: %v", path, err)
		os.Exit(1)
	}
	if err = f.Close(); err != nil {
		fmt.Printf("Failed to close generated file '%s'\nError: %v", path, err)
		os.Exit(1)
	}
}

func updateSolutionsRegistry(rootDir string) {
	solutionsDirPath := path.Join(rootDir, "solutions")
	solutionsRegPath := path.Join(solutionsDirPath, "solutions_reg.go")

	var years []string

	dirEntries, err := os.ReadDir(solutionsDirPath)
	if err != nil {
		fmt.Printf("Failed to read dir '%s'\nError: %v", solutionsDirPath, err)
		os.Exit(1)
	}

	for _, dirEntry := range dirEntries {
		if !dirEntry.IsDir() {
			continue
		}

		if !strings.HasPrefix(dirEntry.Name(), "year_") {
			continue
		}

		years = append(years, strings.TrimPrefix(dirEntry.Name(), "year_"))
	}

	outputTemplateToFile(
		SolutionsRegistryTemplate,
		solutionsRegPath,
		struct{ Years []string }{years},
		true,
	)
}
