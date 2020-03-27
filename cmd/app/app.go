package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	utils "hotel-data-converter/lib"
)

var (
	dataFile      = flag.String("file", "", "the file to convert")
	sort          = flag.String("sort", "", "to sort hotels by Stars. It can either be ascend or descend")
	outputFolder  = "share"
	filesFolder   = fmt.Sprintf("%s/files", outputFolder)
	resultsFolder = fmt.Sprintf("%s/results", outputFolder)
)

const (
	ascending  = "ascend"
	descending = "descend"
)

// App ...
type App struct {
	hotelDataConverter *utils.HotelDataConverter
	outputPath         string
}

// Initialize App
func (a *App) Initialize() {
	flag.Parse()

	// make output folder
	utils.CreateDir(resultsFolder)
	a.hotelDataConverter = utils.NewHotelDataConverter()
}

// ConvertCSVFormat : handles the conversition from CSV Format to other formats
func (a *App) ConvertCSVFormat() {
	records, err := utils.ReadCSVFile(a.hotelDataConverter.File)
	if err != nil {
		log.Fatalln("[ERROR: READ CSV]", err)
	}

	validRecords, invalidRecords := a.hotelDataConverter.GetHotelsFromCSVRecords(records)

	// sort if required
	if strings.ToLower(*sort) == ascending {
		validRecords = a.hotelDataConverter.SortByStars(validRecords, true)
	} else if strings.ToLower(*sort) == descending {
		validRecords = a.hotelDataConverter.SortByStars(validRecords, false)
	}

	jsonData := a.hotelDataConverter.ConvertDataToJSON(&validRecords)
	xmlData := a.hotelDataConverter.ConvertDataToXML(&validRecords)

	fmt.Println("WRITING TO JSON FILE...")
	jsonFile := utils.CreateFile(a.outputPath + ".json")
	defer utils.CloseFile(jsonFile)
	jsonFile.Write(jsonData)

	fmt.Println("WRITING TO XML FILE...")
	xmlFile := utils.CreateFile(a.outputPath + ".xml")
	defer utils.CloseFile(xmlFile)
	xmlFile.Write(xmlData)

	if len(invalidRecords) > 0 {
		fmt.Println("WRITING INVALID RECORDS...")

		titleFields := records[0]
		invalidRecords = append([][]string{titleFields}, invalidRecords...)

		// open file to convert
		file, err := os.OpenFile(a.outputPath+"-invalid.csv", os.O_CREATE|os.O_RDWR, os.ModePerm)
		if err != nil {
			log.Fatalf("[ERROR: OPENING INVALID CSV] %v", err)
		}

		defer utils.CloseFile(file)

		utils.WriteCSVFile(file, invalidRecords)
	}
}

// setOutputpath creates and sets the outpath for the conversion result
func (a *App) setOutputpath(fi os.FileInfo) {
	name := fi.Name()

	s := strings.Split(name, ".")
	// remove the file extension e.g ".csv"
	if len(s) > 1 {
		s = s[:len(s)-1]
	}
	name = strings.Join(s, "")

	a.outputPath = fmt.Sprintf("%s/%d-%s", resultsFolder, a.hotelDataConverter.CreatedAt, name)
}

// Run runs the App
func (a *App) Run() {
	fmt.Println("STARTING DATA CONVERSION...")

	// validate sort flag
	if *sort != "" && strings.ToLower(*sort) != ascending && strings.ToLower(*sort) != descending {
		log.Fatalln("[ERROR: Run] sort value must either be ascend or descend")
	}

	// open file to convert
	dataFilePath := fmt.Sprintf("%s/%s", filesFolder, *dataFile)
	file, err := os.OpenFile(dataFilePath, os.O_RDONLY, os.ModePerm)

	if err != nil {
		log.Fatalf("[ERROR: Run] %v", err)
	}

	defer utils.CloseFile(file)
	a.hotelDataConverter.File = file

	fs, err := file.Stat()
	if err != nil {
		panic(err)
	}

	a.setOutputpath(fs)

	a.ConvertCSVFormat()

	fmt.Println("DONE!")
}
