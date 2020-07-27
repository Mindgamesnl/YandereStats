package analytics

import (
	"github.com/Mindgamesnl/YandereStats/changelog"
	"github.com/olekukonko/tablewriter"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func CountFileRevisions(cl changelog.ChangeLog)  {

	var values = make(map[string]int)
	var firstChanges = make(map[string]string)

	for i := range cl.Revisions {
		revision := cl.Revisions[i]
		for ci := range revision.CommitData.Changes {
			change := revision.CommitData.Changes[ci]
			values[change.FileName]++

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
	table.SetHeader([]string{"File", "Revision", "First released"})
	for _, v := range rows {
		table.Append(v)
	}

	table.Render()

	content := "**This file tracks how many times certain files got changed in the game release changelog, and when they were first created/changed.**\n\n\n\n" + tableString.String()

	_ = ioutil.WriteFile("docs/file_change_graph.md", []byte(content), 0644)
}

func rankByCount(fileFrequencies map[string]int) PairList{
	pl := make(PairList, len(fileFrequencies))
	i := 0
	for k, v := range fileFrequencies {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
}

type Pair struct {
	Key string
	Value int
}

type PairList []Pair

func (p PairList) Len() int { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int){ p[i], p[j] = p[j], p[i] }