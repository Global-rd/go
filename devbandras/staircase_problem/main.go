package main

import (
	"fmt"
	"os"
	"strconv"
)

// környezet változó neve
const C_STAIRS_COUNT = "STAIRS_COUNT"

// sorok száma printeléshez
var lineCount int = 0

// -----------------------------------------------------------------------------
// A clearScreen függvény a ""\033[H\033[2J" escape szekvenciával törli a konzol képernyőjét.
func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

// -----------------------------------------------------------------------------
// Főmenü printelése a konzolra, menüpont bekérése
func mainMenu() {
	clearScreen()

	fmt.Println(`A "lépcsőfok probléma" megoldása`)
	fmt.Println("-----------------------------------------------------------------")
	fmt.Println("Válasszon az alábbi menüből:")
	fmt.Println("-----------------------------------------------------------------")
	fmt.Println("1: Lépcsőfokok bekérése")
	fmt.Println("2: Lépcsőfokok beolvasása környezeti változóból (" + C_STAIRS_COUNT + ")")
	fmt.Println("X: Kilépés")
	fmt.Println("-----------------------------------------------------------------")

	// végtelen ciklus amíg x-et adunk meg
	for {
		var choice string
		fmt.Print("Menü: ")
		fmt.Scanln(&choice)

		if choice == "X" || choice == "x" {
			clearScreen()
			break
		}
		readMenu(choice)
	}
}

// -----------------------------------------------------------------------------
// Menüpontok kiértékelése és a feldolgozó függvények meghívása
func readMenu(menu string) {
	switch menu {
	case "1":
		readFromInput()
	case "2":
		readFromEnvironment()
	default:
		fmt.Println("Érvénytelen menüpont!")
	}
}

// -----------------------------------------------------------------------------
// Lépcsofokok számának típusellenőezése. A függvény ellenőrzi, hogy a megadott paraméter érvényes pozitív egész szám-e.
// Paraméterek:
// - value: az ellenőruzni kívánt érték
// Result:
// Ha a paraméter érték egy érvényes pozitív egész szám, akkor visszaadja az egész szám értékét
// Ha a paraméter nem érvényes pozitív egész szám, hibaüzenetet ír ki, és -1-et ad vissza
func stairTypeCheck(value string) int {
	stairs, err := strconv.Atoi(value)
	if err != nil || stairs <= 0 {
		fmt.Println("Hibás érték! A lépcsőfokok száma csak pozitív egész szám lehet")
		return -1
	}
	return stairs
}

// -----------------------------------------------------------------------------
// A függvény felhasználói oldalró olvassa be a lépcsők számát.
// Ha a bemenet megfelelő akkor meghívja a printSolution() függvényt
// Ha a bemenet nem megfelelő akkor a rekurzívan meghívja önmagát, amíg érvényes bemenetet nem adunk meg
func readFromInput() {
	// billentyűzetről kérjük be az adatokat...
	var s string
	fmt.Print("\nAdja meg a lépcsőfokok számát: ")
	fmt.Scanln(&s)

	// típusellenőrzés (csak pozitév egész számokat fogadunk el)...
	i := stairTypeCheck(s)
	if i <= 0 {
		readFromInput()
	}

	// megoldás kiírása
	printSolution(i)
}

// -----------------------------------------------------------------------------
// A függvény a lépcsők számát a C_STAIRS_COUNT-ban deklarált környezeti változóból olvassa ki
// Ellenőrzi, hogy a környezeti változó deklarálva van-e. Ha nincs hibát írunk ki.
// Meglévő könyezeti változó esetén a stairTypeCheck függvény segítségével ellenőrzi a környezeti változó értékét
// Ha az érték érvényes, meghívja a printSolution függvényt a megoldás megjelenítéséhez
func readFromEnvironment() {
	// a környezeti változót sztringként kezeljük, majd ellenőrizzük a tipusát..
	var s string
	s = os.Getenv(C_STAIRS_COUNT)

	// ha nincs deklarálva a megadott környezeti változó akkor hibát írunk...
	if s == "" {
		fmt.Printf("A %s környezeti változó nincs deklarálva!\n\n", C_STAIRS_COUNT)
		return
	}

	// típusellenőrzés...
	i := stairTypeCheck(s)
	if i <= 0 {
		fmt.Printf("A környezeti változó értéke: %s\n", s)
		return
	}

	// megoldás kiírása
	printSolution(i)
}

// -----------------------------------------------------------------------------
// A függvény kiszámítja, hogy hány különböző módon lehet feljutni a lépcső tetejére
// A függvény bemenetként egy egész számot vesz fel stairCount, amely a lépcsőházban lévő lépcsők számát jelenti
// Paraméterek:
// - stairCount: egy egész szám, amely a lépcsők számát adja meg.
// Result:
// - visszaad egy egész számot, amely a lépcső tetejére vezető különböző utak számát jelenti
func countWays(stairCount int) int {

	// rekurziv meghívás --> lasabb megoldás ???
	if stairCount == 0 || stairCount == 1 {
		return 1
	}
	return countWays(stairCount-1) + countWays(stairCount-2)

	/*if stairCount == 1 {
		return 1
	}
	if stairCount == 2 {
		return 2
	}

	// Tároljuk az előző két értéket (Fibonacci alapú)
	prev1, prev2 := 2, 1
	var result int
	for i := 3; i <= stairCount; i++ {
		result = prev1 + prev2
		prev2 = prev1
		prev1 = result
	}

	return result*/
}

// -----------------------------------------------------------------------------
// A függvény kiszámítja és kiírja az összes lehetséges utat a lépcső tetejére
// Paraméterek:
// - stairCount: egy egész szám, amely a lépcsők számát adja meg.
// - path: egész számok szelete, amely a lépcsőfokok aktuális útvonalát jelöli.
// - wayCount: az eddig megtett lépések számát jelképező egész szám.
func findWays(stairCount int, path []int, wayCount int) {
	if stairCount == 0 {
		// Ha elértük a célt, kiírjuk az aktuális utat
		lineCount++
		fmt.Printf("%d: %d (összesen %d lépés)\n", lineCount, path, wayCount)
		return
	}
	if stairCount >= 1 {
		// 1 lépés megtétele
		findWays(stairCount-1, append(path, 1), wayCount+1)
	}
	if stairCount >= 2 {
		// 2 lépés megtétele
		findWays(stairCount-2, append(path, 2), wayCount+1)
	}
}

// -----------------------------------------------------------------------------
// A függvény kiírja a "lépcsőfok probléma" megoldását
// Kiszámítja és kiírja a lépcső tetejére vezető utak számát, valamint az összes lehetséges lépéskombinációt
// Paraméterek:
// - lépcsőszám: a lépcsőházban lévő lépcsők számát adja meg
func printSolution(stairCount int) {
	lineCount = 0
	fmt.Printf("\nÖsszesen %d módon lehet felmenni %d lépcsőfokon\n", countWays(stairCount), stairCount)
	fmt.Println("Lehetséges lépéskombinációk:")
	findWays(stairCount, []int{}, 0)
	fmt.Println()
}

// -----------------------------------------------------------------------------
// A program belépési pontja
func main() {
	mainMenu()
}
