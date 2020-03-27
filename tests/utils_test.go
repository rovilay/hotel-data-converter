package tests

import (
	"io/ioutil"
	"os"
	"testing"

	utils "hotel-data-converter/lib"
)

var (
	hotelCSVHeaderFields = []string{"name", "address", "stars", "contact", "phone", "uri"}
	validHotelCSVRecord  = []string{
		"The test hotel",
		"01, test street",
		"5",
		"test test",
		"+33 (0)2 24 56 78 90",
		"https://test.com/",
	}
	validHotelCSVRecord2 = []string{
		"The test hotel",
		"01, test street",
		"2",
		"test test",
		"+33 (0)2 24 56 78 90",
		"https://test.com/",
	}
	validHotelCSVRecord3 = []string{
		"The test hotel",
		"01, test street",
		"4",
		"test test",
		"+33 (0)2 24 56 78 90",
		"https://test.com/",
	}
	invalidHotelCSVRecord = []string{
		"",
		"3, invalid test street",
		"6",
		"Mr invalid test",
		"+33 (0)2 24 56 78 90",
		"schneider.fr/index/",
	}
)

type test struct {
	csvRecords                  [][]string
	expectedHotelsCount         int
	expectedInvalidRecordsCount int
	err                         error
}

// TestHotelDataConverter ...
func TestHotelDataConverter(t *testing.T) {
	testCases := []test{
		{
			csvRecords:                  [][]string{hotelCSVHeaderFields, validHotelCSVRecord},
			expectedHotelsCount:         1,
			expectedInvalidRecordsCount: 0,
		},
		{
			csvRecords:                  [][]string{hotelCSVHeaderFields, validHotelCSVRecord, invalidHotelCSVRecord},
			expectedHotelsCount:         1,
			expectedInvalidRecordsCount: 1,
		},
	}

	for _, tc := range testCases {
		hdc := utils.NewHotelDataConverter()
		hs, invalidRecords := hdc.GetHotelsFromCSVRecords(tc.csvRecords)

		if len(hs) != tc.expectedHotelsCount {
			t.Fatalf("expected: %v hotels, got: %v", tc.expectedHotelsCount, len(hs))
		}

		if len(invalidRecords) != tc.expectedInvalidRecordsCount {
			t.Fatalf("expected: %v invalidRecords, got: %v", tc.expectedHotelsCount, len(invalidRecords))
		}
	}
}

func TestHotelDataConverterSortByStars(t *testing.T) {
	hdc := utils.NewHotelDataConverter()
	csvRecords := [][]string{
		hotelCSVHeaderFields,
		validHotelCSVRecord,
		validHotelCSVRecord2,
		validHotelCSVRecord3,
		invalidHotelCSVRecord,
	}
	hs, _ := hdc.GetHotelsFromCSVRecords(csvRecords)

	ascendingHotels := hdc.SortByStars(hs, true)

	if ascendingHotels[0].Stars != 2 {
		t.Fatalf("expected: 2, got: %v", ascendingHotels[0].Stars)
	}

	if ascendingHotels[len(ascendingHotels)-1].Stars != 5 {
		t.Fatalf("expected: 5, got: %v", ascendingHotels[len(ascendingHotels)-1].Stars)
	}

	descendingHotels := hdc.SortByStars(hs, false)

	if descendingHotels[0].Stars != 5 {
		t.Fatalf("expected: 5, got: %v", descendingHotels[0].Stars)
	}

	if descendingHotels[len(descendingHotels)-1].Stars != 2 {
		t.Fatalf("expected: 2, got: %v", descendingHotels[len(descendingHotels)-1].Stars)
	}
}

func setup(t *testing.T) (*os.File, func()) {
	t.Parallel()

	f, err := ioutil.TempFile("", "")
	if err != nil {
		t.Errorf("temp file not created: %v", err)
	}

	tearDown := func() {
		utils.CloseFile(f)
		os.Remove(f.Name())
	}

	return f, tearDown
}

func TestFileUtils(t *testing.T) {
	f, teardown := setup(t)

	defer teardown()

	t.Run("Test ReadCSVFile", func(t *testing.T) {
		if _, err := utils.ReadCSVFile(f); err != nil {
			t.Fatalf("error Reading file: %v", err)
		}
	})

	t.Run("Test ReadCSVFile", func(t *testing.T) {
		csvRecords := [][]string{hotelCSVHeaderFields, validHotelCSVRecord}
		utils.WriteCSVFile(f, csvRecords)
	})
}
