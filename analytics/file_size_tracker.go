package analytics

import (
	"github.com/Mindgamesnl/YandereStats/changelog"
	"github.com/olekukonko/tablewriter"
	"io/ioutil"
	"strconv"
	"strings"
)

func CountFileSize(cl changelog.ChangeLog)  {

	var values = make(map[string]int)
	var firstChanges = make(map[string]string)

	for i := range cl.Revisions {
		revision := cl.Revisions[i]
		for ci := range revision.CommitData.Changes {

			change := revision.CommitData.Changes[ci]
			values[change.FileName] = values[change.FileName] + change.AddedLines
			values[change.FileName] = values[change.FileName] - change.RemovedLines

			if firstChanges[change.FileName] == "" {
				firstChanges[change.FileName] = revision.GitFormattedDate
			}
		}
	}

	list := rankByCount(values)
	var rows [][]string

	for i := range list {
		pair := list[i]

		rows = append(rows, []string{"`" + pair.Key + "`", strconv.FormatInt(int64(pair.Value), 10), "`" + firstChanges[pair.Key] + "`"})
	}

	tableString := &strings.Builder{}
	table := tablewriter.NewWriter(tableString)
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.SetHeader([]string{"File", "Lines", "First released"})
	for _, v := range rows {
		table.Append(v)
	}

	table.Render()

	content := "**This file tracks the length of every file over time. It does this by counting the added/removed lines over time, so now deleted files may also be listed.**\n\n\n\n" + tableString.String()

	_ = ioutil.WriteFile("docs/file_length_breakdown.md", []byte(content), 0644)
}
