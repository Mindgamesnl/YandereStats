package analytics

import (
	"github.com/Mindgamesnl/YandereStats/changelog"
)

var AnalyticalTasks = []func(changelog.ChangeLog){

	// generates a graph page based on the changes per release
	GenerateGraphTask,

}
