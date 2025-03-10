package utils

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"time"
)

type Zip struct {
	CSVFileName string
}

// A NewZip függvény egy Zip struktúra új példányát adja vissza.
func NewZip() *Zip {
	return &Zip{}
}

// Az Execute létrehoz egy ZIP fájlt a megadott CSV fájlból.
// A ZIP fájl a CSV fájl nevével megegyező néven, de „.zip” kiterjesztéssel jön létre
// A függvény megnyitja a CSV fájlt, létrehoz egy új ZIP fájlt, és a CSV fájlt beírja a ZIP fájlba
//
// Parameters:
//   - z: a Zip struktúra mutatója
//
// Returns:
//   - error: a tömörítés során keletkező hibát tartalmazza.
func (z *Zip) Execute() error {
	// csv fájl megnyitása
	sourceFile, err := os.Open(z.CSVFileName)
	if err != nil {
		return fmt.Errorf("error opening source file: %v", err)
	}
	defer sourceFile.Close()

	// zip fájl létrehozása
	zipFileName := fmt.Sprintf("%s.zip", z.CSVFileName)
	zipFile, err := os.Create(zipFileName)
	if err != nil {
		return fmt.Errorf("error creating ZIP file: %v", err)
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	header := zip.FileHeader{
		Name:     z.CSVFileName,
		Method:   zip.Deflate,
		Modified: time.Now(),
	}

	writer, err := zipWriter.CreateHeader(&header)
	if err != nil {
		return fmt.Errorf("error creating ZIP header: %v", err)
	}

	_, err = io.Copy(writer, sourceFile)
	if err != nil {
		return fmt.Errorf("error copying file to ZIP: %v", err)
	}

	return nil
}
