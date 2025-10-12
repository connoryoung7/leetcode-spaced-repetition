package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

const leetcodeProblemsDir = "leetcode-company-wise-problems"

type LeetcodeProblem struct {
	Name       string
	Difficulty QuestionDifficulty
	Slug       string
	Acceptance float64
	Frequency  float64
	TimeTag    string
}

type QuestionDifficulty int

type TimeTag string

const (
	LastThirtyDays    = "last_thirty_days"
	LastThreeMonths   = "last_three_months"
	LastSixMonths     = "last_six_months"
	MoreThanSixMonths = "more_than_six_months"
)

const (
	EasyDifficulty QuestionDifficulty = iota + 1
	MediumDifficulty
	HardDifficulty
)

const (
	DifficultyIdx int = iota
	NameIdx
	FrequencyIdx
	AcceptanceIdx
	LinkIdx
)

func convertFileNameToTimeTag(filename string) (TimeTag, error) {
	lCase := strings.ToLower(filename)

	if strings.Contains(lCase, "more than six months") {
		return MoreThanSixMonths, nil
	} else if strings.Contains(lCase, "six months") {
		return LastSixMonths, nil
	} else if strings.Contains(lCase, "three months") {
		return LastThreeMonths, nil
	} else if strings.Contains(lCase, "thirty days") {
		return LastThirtyDays, nil
	}
	return LastThirtyDays, fmt.Errorf("'%s' is not a valid time tag", filename)
}

func convertStringToDifficulty(diffStr string) (QuestionDifficulty, error) {
	lowerDiffStr := strings.ToLower(strings.Trim(diffStr, " `"))

	switch lowerDiffStr {
	case "easy":
		return EasyDifficulty, nil
	case "medium":
		return MediumDifficulty, nil
	case "hard":
		return HardDifficulty, nil
	default:
		return EasyDifficulty, fmt.Errorf("'%s' is not a valid difficulty", diffStr)
	}
}

func getSlugFromLink(link string) string {
	parts := strings.Split(link, "/")

	return parts[len(parts)-1]
}

func main() {
	fmt.Println("Downloading the leetcode problems...")

	cmd := exec.Command("git", "clone", "https://github.com/liquidslr/leetcode-company-wise-problems.git")
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Downloaded the leetcode company wise problems")

	companies, err := os.ReadDir(leetcodeProblemsDir)
	if err != nil {
		fmt.Println(err)
	}

	companyNames := []string{}

	for _, e := range companies {
		companyName := e.Name()
		hiddenFile := strings.Contains(companyName, ".")
		if !hiddenFile {
			companyNames = append(companyNames, companyName)

		}
		fmt.Println(e.Name())
	}

	allProblems := []LeetcodeProblem{}

	for _, company := range companyNames {
		problemsByFreq, err := os.ReadDir(filepath.Join(leetcodeProblemsDir, company))
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, freq := range problemsByFreq {
			freqName := freq.Name()
			fmt.Println(freqName)
			fPath := filepath.Join(leetcodeProblemsDir, company, freqName)
			file, err := os.Open(fPath)
			if err != nil {
				fmt.Println(err)
			}
			reader := csv.NewReader(file)
			records, _ := reader.ReadAll()

			for _, record := range records {
				freq, err := strconv.ParseFloat(record[FrequencyIdx], 5)
				if err != nil {

				}

				acceptance, err := strconv.ParseFloat(record[AcceptanceIdx], 5)
				if err != nil {

				}

				difficulty, err := convertStringToDifficulty(record[DifficultyIdx])
				if err != nil {

				}

				timeTag, err := convertFileNameToTimeTag(freqName)
				if err != nil {

				}

				allProblems = append(allProblems, LeetcodeProblem{
					Name:       record[NameIdx],
					Frequency:  freq,
					Difficulty: difficulty,
					Acceptance: acceptance,
					Slug:       getSlugFromLink(record[LinkIdx]),
					TimeTag:    string(timeTag),
				})
			}
		}
	}

	fmt.Printf("%+v\n", allProblems)

	defer func() {
		os.RemoveAll(leetcodeProblemsDir)
		fmt.Println("Removed")
	}()
}
