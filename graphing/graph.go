package graphing

import (
	"github.com/Mindgamesnl/YandereFetch/changelog"
	"github.com/go-echarts/go-echarts/charts"
	"github.com/sirupsen/logrus"
	"log"
	"math/rand"
	"os"
	"path"
	"time"
)

const maxNum = 50

var seed = rand.NewSource(time.Now().UnixNano())

func GenerateGraph(cl changelog.ChangeLog) {
	logrus.Info("Making main graphs, fun!")


	AddRemoveGraph := charts.NewLine()
	AddRemoveGraph.SetGlobalOptions(
		charts.TitleOpts{Title: "Added and removed lines", Subtitle: "Lines that got added/removed per update."},
		charts.TooltipOpts{Show: true},
		charts.DataZoomOpts{Type: "slider"},
	)

	var addedPerUpdate = []int{0}
	var removedPerUpdate = []int{0}
	var updateDates = []string{""}
	for i := range cl.Revisions {
		revision := cl.Revisions[i]
		updateDates = append(updateDates, revision.GitFormattedDate)

		addedLines := len(revision.CommitData.AddedLines)

		removedLines := len(revision.CommitData.RemovedLines)

		addedPerUpdate = append(addedPerUpdate, addedLines)
		removedPerUpdate = append(removedPerUpdate, removedLines)
	}

	AddRemoveGraph.AddXAxis(updateDates).AddYAxis("Added lines", addedPerUpdate)
	AddRemoveGraph.AddXAxis(updateDates).AddYAxis("Removed lines", removedPerUpdate)

	page := charts.NewPage(charts.RouterOpts{URL: "url", Text: "text"})
	page.Add(
		AddRemoveGraph,
	)
	f, err := os.Create(getRenderPath("index.html"))
	if err != nil {
		log.Println(err)
	}
	page.Render(f)
}

func getRenderPath(f string) string {
	return path.Join("docs", f)
}
