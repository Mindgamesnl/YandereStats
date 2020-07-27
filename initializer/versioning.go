package initializer

import (
	"github.com/Mindgamesnl/YandereFetch/changelog"
	"github.com/Mindgamesnl/YandereFetch/git"
	"github.com/PuerkitoBio/goquery"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func InitializeGameVersions() changelog.ChangeLog {
	logrus.Info("Loading and parsing game versions")
	start := time.Now()

	res, err := http.Get("https://yandere-simulator.fandom.com/wiki/Update_History")
	if err != nil {
		logrus.Error("Could not open fandom page with update history")
		log.Fatal(err)
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var changeLog []changelog.ChangelogRevision

	// Find the review items
	doc.Find("#mw-content-text > table").Each(func(i int, results *goquery.Selection) {
		// find all rows

		results.Find("tr").Each(func(i int, selection *goquery.Selection) {
			date := selection.Find("td:nth-child(1)").Text()
			date = strings.ReplaceAll(date, "\n", "")
			if date != "" {
				// convert date to git format
				parts := strings.Split(date, "/")
				gitDate := "ys" + parts[2] + "." + MakeTwoDecimal(parts[0]) + "." + MakeTwoDecimal(parts[1])

				parsedDate, _ := time.Parse("2006-01-02", parts[2] + "-" + MakeTwoDecimal(parts[0]) + "-" + MakeTwoDecimal(parts[1]))
				gitDate = strings.ReplaceAll(gitDate, "\n", "")
				revision := changelog.ChangelogRevision{Date: parsedDate, FormattedWebDate: date, GitFormattedDate: gitDate}

				selection.Find("td:nth-child(2) > ul").Find("li").Each(func(i int, selection *goquery.Selection) {
					message := selection.Text()
					message = strings.ReplaceAll(message, "\n", "")
					revision.Note = append(revision.Note, git.NoteMessage{
						Message: message,
					})
				})

				changeLog = append(changeLog, revision)
			}
		})
	})

	logrus.Info("Scraped " + strconv.FormatInt(int64(len(changeLog)), 10) + " updates from the wiki")
	elapsed := time.Since(start)
	logrus.Info("Scraping took ", elapsed)
	return changelog.ChangeLog{Revisions: changeLog}
}

func MakeTwoDecimal(input string) string {
	intValue, _ := strconv.ParseInt(input, 10, 64)
	if intValue < 10 {
		return "0" + strconv.FormatInt(intValue, 10)
	}
	return input
}
