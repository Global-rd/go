/*

 */

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"sync"
)

const (
	defaultPufferSize         = 10
	jsonFile                  = "countries.json"
	countrtiesApi             = "https://restcountries.com/v3.1/all"
	simulatedPanicCountryCode = "RO"
)

// Country struktúra egy ország adatairól
type Country struct {
	Name struct {
		Common   string `json:"common"`   // Az ország általános neve
		Official string `json:"official"` // az ország hivatalos neve
	} `json:"name"`
	CCA2    string   `json:"cca2"`    // ISO 3166-1 alpha-2 országkód
	CCA3    string   `json:"cca3"`    // ISO 3166-1 alpha-3 országkód
	Capital []string `json:"capital"` // Főváros(ok)
	Region  string   `json:"region"`  // Régió
	Maps    struct {
		GoogleMaps     string `json:"googleMaps"`     // Google Maps URL
		OpenStreetMaps string `json:"openStreetmaps"` // OpenStreetMap URL
	} `json:"maps"`
	Flags struct {
		SVG string `json:"svg"` // SVG zászló URL
		PNG string `json:"png"` // PNG zászló URL
	} `json:"flags"`
}

// Writer interface az országadatok írásához, puffereléshez és lezárásához
type Writer interface {
	Write(Country) error
	Flush() error
	Close() error
}

// BatchWriter
type BatchWriter struct {
	puffer      []Country
	pufferSize  int
	file        *os.File
	mutex       sync.Mutex
	isFirst     bool // első írás-e
	puffercount int  // pufferelt sorok száma
}

var logger *slog.Logger

// NewBatchWriter létrehoz és inicializál egy új BatchWriter-t
func NewBatchWriter(filename string, pufferSize int) (*BatchWriter, error) {

	logger.Info("BatchWriter létrehozása...", "fájl", filename, "pufferméret", pufferSize)

	// Puffer méretének ellenőrzése
	if pufferSize < 1 {
		pufferSize = defaultPufferSize
	}

	// Fájl megnyitása írásra
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return nil, fmt.Errorf("fájl létrehozása sikertelen: %v", err)
	}

	// Kezdő zárójel kiírása mert egy tömböt írunk
	_, err = file.Write([]byte("[\n"))
	if err != nil {
		file.Close()
		return nil, fmt.Errorf("kezdő zárójel írása sikertelen: %v", err)
	}

	return &BatchWriter{
		pufferSize:  pufferSize,
		file:        file,
		isFirst:     true,
		puffercount: 1,
	}, nil
}

// Write hozzáad egy országot a pufferhez és kiírja ha a puffer megtelt.
func (bw *BatchWriter) Write(country Country) error {

	fmt.Printf("#%d/%d - Ország irása pufferbe: %s - %s\n", len(bw.puffer)+1, bw.puffercount, country.CCA2, country.Name.Common)
	bw.puffercount++

	bw.mutex.Lock()
	bw.puffer = append(bw.puffer, country)
	shouldFlush := len(bw.puffer) >= bw.pufferSize
	bw.mutex.Unlock()

	if shouldFlush {
		return bw.Flush()
	}
	return nil
}

// Flush kiírja az összes pufferelt országot a fájlba.
func (bw *BatchWriter) Flush() error {
	bw.mutex.Lock()
	defer bw.mutex.Unlock()

	if len(bw.puffer) == 0 {
		return nil
	}

	fmt.Println("\nPuffer kiírása...")
	fmt.Println()

	// Minden országot külön írunk ki
	for _, country := range bw.puffer {
		// Ország JSON formázása
		data, err := json.MarshalIndent(country, "  ", "  ")
		if err != nil {
			return fmt.Errorf("JSON konvertálás sikertelen: %v", err)
		}

		// Ha nem az első elem, vessző hozzáadása
		if !bw.isFirst {
			_, err = bw.file.Write([]byte(",\n"))
			if err != nil {
				return fmt.Errorf("vessző írása sikertelen: %v", err)
			}
		}

		// JSON adat kiírása
		_, err = bw.file.Write(data)
		if err != nil {
			return fmt.Errorf("adat írása sikertelen: %v", err)
		}

		bw.isFirst = false
	}

	// Puffer ürítése
	bw.puffer = nil

	return nil
}

// Close befejezi a JSON tömböt és lezárja a fájlt.
func (bw *BatchWriter) Close() error {
	// Maradék adatok kiírása
	if err := bw.Flush(); err != nil {
		return fmt.Errorf("végső flush sikertelen: %v", err)
	}

	// Záró zárójel kiírása
	_, err := bw.file.Write([]byte("\n]"))
	if err != nil {
		return fmt.Errorf("záró zárójel írása sikertelen: %v", err)
	}

	// Fájl lezárása
	if err := bw.file.Close(); err != nil {
		return fmt.Errorf("fájl lezárása sikertelen: %v", err)
	}

	println("Fájl írása és lezárása sikeres")

	return nil
}

// FetchCountries lekéri az országadatokat a REST Countries API-ról.
func FetchCountries() ([]Country, error) {
	resp, err := http.Get("https://restcountries.com/v3.1/all")
	if err != nil {
		return nil, fmt.Errorf("API hívás sikertelen: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("API válasz olvasása sikertelen: %v", err)
	}

	var countries []Country
	if err := json.Unmarshal(body, &countries); err != nil {
		return nil, fmt.Errorf("JSON feldolgozás sikertelen: %v", err)
	}

	return countries, nil
}

func main() {
	logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	bw, err := NewBatchWriter(jsonFile, defaultPufferSize)
	if err != nil {
		logger.Error("Hiba a batch writer létrehozásakor", "error", err)
		os.Exit(1)
	}
	defer bw.Close()
	logger.Info("Batch writer létrehozva", "fájl", jsonFile, "pufferméret", defaultPufferSize)

	countries, err := FetchCountries()
	if err != nil {
		logger.Error("Hiba a batch writer létrehozásakor", "error", err)
		os.Exit(1)
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("\nHelyreállitás pánikból...", r)
			bw.Flush()
		}
	}()

	for _, country := range countries {
		if err := bw.Write(country); err != nil {
			fmt.Println("Hiba az ország kiírásakor:", err)
		}

		if country.CCA2 == simulatedPanicCountryCode {
			msg := fmt.Sprintf("\n - országkód: %s - %s", simulatedPanicCountryCode, country.Name.Common)
			panic(msg)
		}
	}
}
