# Yandere Fetch
Yandere fetch is a utility that scans code differences between every release and makes a relational database of each update with its provided change log and code deference. This data then gets used to render graphs about the project progress and changelog claims get validated (eg, if the changelog mentions a keyword without a key file being changed, then mark it as **false**)

This project does not contain any of YS' code, nor is it 100% accurate (file changes in files that haven't been tracked prior to that release count as a 100% net gain in total lines, since the whole file is being added to the system in that release).

It's just nice to visualize drama, and that's exactly what this project does, lmao.