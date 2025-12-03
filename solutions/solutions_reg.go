package solutions

import (
	"github.com/aivarasbaranauskas/aoc/solutions/framework"
	"github.com/aivarasbaranauskas/aoc/solutions/year_2015"
	"github.com/aivarasbaranauskas/aoc/solutions/year_2018"
	"github.com/aivarasbaranauskas/aoc/solutions/year_2019"
	"github.com/aivarasbaranauskas/aoc/solutions/year_2021"
	"github.com/aivarasbaranauskas/aoc/solutions/year_2022"
	"github.com/aivarasbaranauskas/aoc/solutions/year_2023"
	"github.com/aivarasbaranauskas/aoc/solutions/year_2024"
	"github.com/aivarasbaranauskas/aoc/solutions/year_2025"
)

var solutions = map[int]map[int]framework.Solution{
	2015: year_2015.Solutions,
	2018: year_2018.Solutions,
	2019: year_2019.Solutions,
	2021: year_2021.Solutions,
	2022: year_2022.Solutions,
	2023: year_2023.Solutions,
	2024: year_2024.Solutions,
	2025: year_2025.Solutions,
}
