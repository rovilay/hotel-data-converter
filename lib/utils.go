package utils

import (
	"encoding/csv"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"time"
	"unicode/utf8"
)

// NewHotelDataConverter instantiates and returns a new HotelDataConverter
func NewHotelDataConverter() *HotelDataConverter {
	now := time.Now()

	hc := &HotelDataConverter{
		CreatedAt: now.Unix(),
	}

	return hc
}

// GetHotelsFromCSVRecords converts CSV Records to hotel data using the Hotel struct
func (hc *HotelDataConverter) GetHotelsFromCSVRecords(csvRecords [][]string) (hs []Hotel, invalidRecords [][]string) {
	fmt.Println("MATCHING CSV DATA TO HOTEL STRUCT...")
	var h Hotel

	// convert each record to Hotel
	// we exclude the first record because of the assumption that they represent
	// the field names
	for _, record := range csvRecords[1:] {
		if hc.isCSVRecordValid(record) {
			h.Name = record[0]
			h.Address = record[1]
			h.Stars, _ = strconv.Atoi(record[2])
			h.Contact = record[3]
			h.Phone = record[4]
			h.URI = record[5]

			hs = append(hs, h)
		} else {
			invalidRecords = append(invalidRecords, record)
		}
	}

	return hs, invalidRecords
}

// ConvertDataToJSON converts hotels to JSON data
func (hc *HotelDataConverter) ConvertDataToJSON(hotels *[]Hotel) []byte {
	fmt.Println("CONVERTING DATA TO JSON...")

	jsonRecord, err := json.Marshal(hotels)

	if err != nil {
		log.Fatalln("[ERROR: ConvertDataToJSON]", err)
	}

	return jsonRecord
}

// ConvertDataToXML convets hotels to XML data
func (hc *HotelDataConverter) ConvertDataToXML(hotels *[]Hotel) []byte {
	fmt.Println("CONVERTING DATA TO XML...")

	xmlRecord, err := xml.Marshal(&hotels)

	if err != nil {
		log.Fatalln("[ERROR: ConvertDataToXML]", err)
	}

	return xmlRecord
}

// SortByStars sorts hotels by the hotel's stars
func (hc *HotelDataConverter) SortByStars(hotels []Hotel, ascending bool) []Hotel {
	sort.SliceStable(hotels, func(i, j int) bool {
		if ascending {
			return hotels[i].Stars < hotels[j].Stars
		}

		return hotels[i].Stars > hotels[j].Stars
	})

	return hotels
}

// ValidateRecord validates the name, stars and uri fields of the CSV record
func (hc *HotelDataConverter) isCSVRecordValid(record []string) bool {
	valid := false
	const (
		maxStars = 5 // maximum number of stars a hotel can have
		minStars = 0 // minimum number of stars a hotel can have
	)

	numOfHotelFields := reflect.TypeOf(Hotel{}).NumField()
	// the records must be greater than or equal to the fields in the Hotel struct
	valid = len(record) >= numOfHotelFields

	// based on the order of fields title...
	// first field is name
	// third field is stars
	// sixth field is URI

	// validate name
	if valid {
		valid = utf8.ValidString(record[0])
	}

	// validate stars
	if valid {
		i, err := strconv.Atoi(record[2])

		if err != nil || i < minStars || i > maxStars {
			valid = false
		}
	}

	// validate uri
	if valid {
		expression := `^(http:\/\/www\.|https:\/\/www\.|http:\/\/|https:\/\/)?[a-z0-9]+([\-\.]{1}[a-z0-9]+)*\.[a-z]{2,5}(:[0-9]{1,5})?(\/.*)?$`
		re := regexp.MustCompile(expression)

		if !re.MatchString(record[5]) {
			valid = false
		}
	}

	return valid
}

// ReadCSVFile reads from a CSV file
func ReadCSVFile(file io.Reader) ([][]string, error) {
	csvReader := csv.NewReader(file)
	csvReader.FieldsPerRecord = -1

	return csvReader.ReadAll()
}

// WriteCSVFile writes records to a CSV file
func WriteCSVFile(file io.ReadWriteSeeker, records [][]string) {
	csvWriter := csv.NewWriter(file)

	for _, r := range records {
		if err := csvWriter.Write(r); err != nil {
			panic(err)
		}
	}

	csvWriter.Flush()

	if err := csvWriter.Error(); err != nil {
		panic(err)
	}
}

// CreateDir creates a directory
func CreateDir(path string) {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		panic(err)
	}
}

// CreateFile helps create file according to the path provided
func CreateFile(path string) *os.File {
	f, err := os.Create(path)

	if err != nil {
		panic(err)
	}

	return f
}

// CloseFile closes the file
func CloseFile(f *os.File) {
	err := f.Close()

	if err != nil {
		panic(err)
	}
}
