package analytics

import (
	"github.com/Mindgamesnl/YandereStats/changelog"
	"github.com/olekukonko/tablewriter"
	"io/ioutil"
	"strconv"
	"strings"
)

var CodeKeywords = []string{
	"void",
	"IF",
	"ELSE",
	"class",
	"static",
	"case",
	"switch",
	"return",
	"throw",
	"catch",
	"onUpdate",
	"FindObjectOfType",
	"GetComponentsInChildren",
	"FindObjectsOfTypeAll",
	"GetComponents",
	"SendMessageUpwards",
	"BroadcastMessage",
	"OnGUI",
	"final",
	"[",
	"null",
	"red",
	"green",
	"yes",
	"no",
	"on",
	"off",
}

func CSOperationBreakdown(cl changelog.ChangeLog)  {

	var values = make(map[string]int)
	var firstChanges = make(map[string]string)

	for i := range cl.Revisions {
		revision := cl.Revisions[i]
		for i2 := range revision.CommitData.AddedLines {
			lineUpdate := revision.CommitData.AddedLines[i2]
			line := lineUpdate.Code
			for i2 := range CodeKeywords {
				keyword := strings.ToLower(CodeKeywords[i2])
				if strings.Contains(line, keyword) {
					if lineUpdate.Action == "ADD" {
						values[keyword]++
					} else {
						values[keyword]--
					}

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

		rows = append(rows, []string{"```c# " + pair.Key + "```", strconv.FormatInt(int64(pair.Value), 10), firstChanges[pair.Key]})
	}

	tableString := &strings.Builder{}
	table := tablewriter.NewWriter(tableString)
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.SetHeader([]string{"Keyword", "occurrences", "First used"})
	for _, v := range rows {
		table.Append(v)
	}

	table.Render()

	content := "**This class tracks how many times certain C# keywords are used. It does this by processing the added/removed lines over time, so now deleted files may also be listed.**\n\n\n\n" + tableString.String()
	_ = ioutil.WriteFile("docs/code_keyword_occurrences.md", []byte(content), 0644)
}