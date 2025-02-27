# Recover from panic during data processing

A program a https://restcountries.com/v3.1/all REST API-ból lekért országadatokat dolgozza fel majd azt egy JSON fájl-ban menti el pufferelve.

---

### A Program működése

#### 1. Fájl inicializálása
A program létrehoz egy BatchWriter-t, amely létrehozza a `jsonFile` konstansban megadott fájlt, és a kezdő zárójelet ([) kiírja a fájlba ezzel megkezdve a tömböt.

#### 2. Országadatok lekérése
A FetchCountries függvény segítségével a program lekéri az összes ország adatait a REST API-ból majd belerakja egy `Country` típusú tömbbe.

#### 3. Adatok feldolgozása és írása:
Az országadatok egyenként bekerülnek a BatchWriter pufferébe. A puffer méretét a `defaultPufferSize` konstans értékkel tudjuk beállítani. Ha a puffer megtelik, a tartalma kiírásra kerül a fájlba.
Ha az országkód megegyezik a `simulatedPanicCountryCode` értékével, a program szándékosan pánikba esik, és a pufferben lévő adatokat kiírja a fájlba, majd lezárja azt a leállás előtt. A `simulatedPanicCountryCode` értéknek valós országkódnk kell lennie (pl.: RO, HU, IT, FR, stb).

#### 3. Fájl lezárása
Miután az összes országadat feldolgozásra került, a program lezárja a fájlt, és kiírja a záró zárójelet (]), ezzel befejezve a JSON tömböt.

---

### Program kimenete
A program futása során a konzolon megjelenik a feldolgozott sorok száma a pufferelt sorok száma és az országok kódja, neve. Ha a puffer megtelik, akkor a program kiírja a puffert a fájlba. Ha a program pánikba esik (pl. `simulatedPanicCountryCode = "FR"`), akkor a program helyreáll, és a pufferben lévő adatokat kiírja a fájlba.






