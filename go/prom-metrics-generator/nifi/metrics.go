package nifi

import (
        "prom-metrics-generator/logger"
        "os"
        "io/ioutil"
		"errors"
)

func createDir(
	path string,
){
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			clog.Error.Println(err)
		}
	}
}

func createFile( 
	path string,
) {
	// check if file exists
	var _, err = os.Stat(path)

	// create file if not exists
	if os.IsNotExist(err) {
			var file, err = os.Create(path)
			if err != nil {
					clog.Error.Println(err.Error()+path)
			}
			defer file.Close()
	}
}

func writeMetrics(
	Name string,
	Help string,
	Type string,
	Field string,
	Match string,
	Value string,        
) {
	metricsDir := "/tmp/metrics"
	metricsPath := metricsDir+"/metrics"
	tempMetricsPath := metricsDir+"/temp.metrics"

	// create dir if not exists
	createDir(metricsDir)
	
	// create file if not exists
	createFile(metricsPath)

	metricsFile, err := os.OpenFile(metricsPath, os.O_RDWR, 0644)
	if err != nil {
			clog.Error.Println(err.Error()+metricsPath)
	}
	defer metricsFile.Close()
	
	// create temp metrics file
	createFile(tempMetricsPath)

	tempFile, err := os.OpenFile(tempMetricsPath, os.O_RDWR, 0644)
	if err != nil {
			clog.Error.Println(err.Error()+tempMetricsPath)
	}
	defer tempFile.Close()

	// Get matching line and dump metrics file data to slice
	matchedLineNumber,allLines := fileSliceLineMatch(metricsFile, Name+Match)

	if len(matchedLineNumber) > 0{
			// reading from file slice
			for n, line := range allLines {
					// looping over matched line list
					for _, ln := range matchedLineNumber{
							// if the current line is same as matchedLine than update the matric
							if ln == n+1 {
									// Write updated matric
									updatedMatric := Name+Field+" "+Value+"\n"
									_, err = tempFile.WriteString(updatedMatric)
									if err != nil {
											clog.Error.Println(err.Error())
									}
									// Save file changes.
									err = tempFile.Sync()
									if err != nil {
											clog.Error.Println(err.Error())
									}
							} else{
									_, err = tempFile.WriteString(line+"\n")
									if err != nil {
											clog.Error.Println(err.Error())
									}
									// Save file changes.
									err = tempFile.Sync()
									if err != nil {
											clog.Error.Println(err.Error())
									} 
							}
					}
			}

			input, err := ioutil.ReadFile(tempMetricsPath)
			if err != nil {
					clog.Error.Println(err.Error())
			}
	
			//tempMetricsPathInProgress := tempMetricsPath+".in-progress"
			err = ioutil.WriteFile(tempMetricsPath+".in-progress", input, 0644)
			if err != nil {
					clog.Error.Println(err.Error())
			}

			// delete metrics file
			err = os.Remove(metricsPath)
			if err != nil {
					clog.Error.Println(err.Error())
			}

			// Copy temp in-progress file to metrics file
			err = os.Rename(tempMetricsPath+".in-progress", metricsPath)
			if err != nil {
					clog.Error.Println(err.Error())
			}
	}else{
			// Add HELP for new metric
			addMatricHelp := "# HELP "+Name+" "+Help+"\n"
			_, err = metricsFile.WriteString(addMatricHelp)
			if err != nil {
					clog.Error.Println(err.Error())
			}
			// Add TYPE for new metric
			addMatricType := "# TYPE "+Name+" "+Type+"\n"
			_, err = metricsFile.WriteString(addMatricType)
			if err != nil {
					clog.Error.Println(err.Error())
			}
			// Add matric and value for new metric
			addMatricValue := Name+Field+" "+Value+"\n"
			_, err = metricsFile.WriteString(addMatricValue)
			if err != nil {
					clog.Error.Println(err.Error())
			}
			// Save file changes.
			err = metricsFile.Sync()
			if err != nil {
					clog.Error.Println(err.Error())
			}
	}
}
