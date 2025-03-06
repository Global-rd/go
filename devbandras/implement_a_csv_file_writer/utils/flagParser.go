package utils

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type Flags struct {
	Source      string
	Destination string
	Zip         bool
}

// A függvény inicializál egy új Flags struktúrát
func NewFlags() *Flags {
	return &Flags{}
}

// A Parse elemzi a parancssori kapcsokókat és feltölti a Flags struktúrát a megadott értékekkel.
// -s, -source: A JSON forrásfájl elérési útja (kötelező).
// -d, -destination: A cél CSV fájl elérési útvonala (kötelező)
// -z, -zip: A CSV archiválása ZIP fájlba (opcionális, alapértelmezett false).
//
// A függvény egy mutatót ad vissza a kitöltött Flags struktúrára.
func (f *Flags) Parse() *Flags {
	// flag names
	sourceFlagNames := []string{"s", "source"}
	destFlagNames := []string{"d", "destination"}
	zipFlagNames := []string{"z", "zip"}

	var source, destination string
	var zip bool

	for _, name := range sourceFlagNames {
		flag.StringVar(&source, name, "", "Path to the source JSON file")
	}

	for _, name := range destFlagNames {
		flag.StringVar(&destination, name, "", "Path to the destination CSV file")
	}

	for _, name := range zipFlagNames {
		flag.BoolVar(&zip, name, false, "Archive the CSV into a ZIP file")
	}

	flag.Parse()

	f.Source = source
	f.Destination = destination
	f.Zip = zip

	return f
}

// A Validate függvény ellenőrzi, hogy a parancssori kapcsolók érvényesek-e.
//
// Parameters:
// -
//
// Returns:
//   - error: Az ellenőrzés során felmerült hibák listája.
func (f *Flags) Validate() error {
	var errors []string

	if f.Source == "" {
		errors = append(errors, "source path is required")
	}

	if f.Destination == "" {
		errors = append(errors, "destination path is required")
	}

	if len(errors) > 0 {
		return fmt.Errorf("validation errors: %s", strings.Join(errors, ", "))
	}

	return nil
}

// A PrintUsage kiírja a parancssori kapcsolók elérhető opcióit, leírását és példákat a használatukra.
func PrintUsage() {
	fmt.Printf("Usage: %s [options]\n", os.Args[0])
	fmt.Println("Options:")

	flag.PrintDefaults()

	fmt.Println("\nExample usages:")
	fmt.Println("csv_writer -s <source> -d <destination> [-z]")
	fmt.Println("csv_writer --source <source> --destination <destination> [--zip]")
}
