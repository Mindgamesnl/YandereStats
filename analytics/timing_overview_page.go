package analytics

import (
	"github.com/Mindgamesnl/YandereStats/utils"
	"github.com/olekukonko/tablewriter"
	"io/ioutil"
	"strings"
)

func GenerateStopwatchBreakdown() {

	var rows [][]string

	for s := range utils.TimingLogs {
		duration := utils.TimingLogs[s]

		rows = append(rows, []string{"`" + s + "`" , "`" + duration.String() + "`"})
	}

	tableString := &strings.Builder{}
	table := tablewriter.NewWriter(tableString)
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.SetHeader([]string{"Task", "Elapsed time"})
	for _, v := range rows {
		table.Append(v)
	}

	table.Render()

	content := "**Debug timing output of how long it took to generate these statistics. Still faster than the game boot times, lmao.**\n\n\n\n" + tableString.String()

	_ = ioutil.WriteFile("docs/time_breakdown.md", []byte(content), 0644)
}