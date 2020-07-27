package analytics

import (
	"github.com/Mindgamesnl/YandereStats/changelog"
	"github.com/olekukonko/tablewriter"
	"io/ioutil"
	"strconv"
	"strings"
)

var Keywords = []string{
	"Osana",
	"Fix",
	"Added",
	"Implemented",
	"Removed",
	"Disabled",
	"Hate",
	"Kill",
	"Crash",
	"Update",
	"I",
	"Reddit",
	"Leak",
	"Hacks",
	"Mod",
	"Fuck",
	"Dev",
	"Stop",
	"No",
}

func CountChangeLogEntries(cl changelog.ChangeLog)  {

	var values = make(map[string]int)
	var firstChanges = make(map[string]string)

	for i := range cl.Revisions {
		revision := cl.Revisions[i]
		for ci := range revision.Note {
			note := revision.Note[ci]
			message := strings.ToLower(note.Message)

			for i2 := range Keywords {
				keyword := strings.ToLower(Keywords[i2])

				if strings.Contains(message, keyword) {
					values[keyword]++

					if firstChanges[keyword] == "" {
						firstChanges[keyword] = revision.GitFormattedDate
					}
				}
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
	table.SetHeader([]string{"Changelog keyword", "occurrences", "First mentioned"})
	for _, v := range rows {
		table.Append(v)
	}

	table.Render()

	content := "**This file tracks how many times certain words got mentioned in the game release changelog, and when they were first mentioned.**\n\n\n\n" + tableString.String()
	_ = ioutil.WriteFile("docs/changelog_keyword_occurrences.md", []byte(content), 0644)
}