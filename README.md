# Yandere Stats
Yandere Stats is a utility that scans code differences between every release and makes a relational database of each update with its provided change log and code deference. This data then gets used to render graphs about the project progress and changelog claims get validated (eg, if the changelog mentions a keyword without a key file being changed, then mark it as **false** or **plausible**)

This project does not contain any of YS' code, nor is it 100% accurate (file changes in files that haven't been tracked prior to that release count as a 100% net gain in total lines, since the whole file is being added to the system in that release).

Main output gets rendered on [this page. Most data is in the database that gets created by this project, which I use to do research on the codebase and it's evolution over time. Questions? contact me on Discord. `Mats - Mindgamesnl#6985`

## Collected data (so far, more on that later)
 - [Code changes per release](docs/index.html)
 - [File changes and introductions](docs/file_change_graph.md)
 - [Changelog analysis](docs/changelog_keyword_assurances.md)
 - [File length breakdown](docs/file_length_breakdown.md)
 - [C# keywrods/function analysis](docs/code_keyword_occurrences.md)
### (Small notes)
 - This is the result of a sleepdeprived yet curious allgnither and is still a work in progress, want to write more systems to automatically analyze data than just that one graph.
 - Want to know more? I can't publish the dataset itself for ovbius reasons but feel free to make a Issue on github or message me on Discord if you have any ideas/suggestions.
