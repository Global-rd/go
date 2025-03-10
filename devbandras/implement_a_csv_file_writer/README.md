# CSV Writer Eszköz

Egyszerű parancssori segédprogram JSON fájlok CSV formátumba konvertálásához, opcionális ZIP tömörítéssel.

## Funkciók

- JSON adatok konvertálása CSV formátumba
- Konfigurálható elválasztó karakter (alapértelmezetten pontosvessző `;`)
- Opcionális ZIP tömörítés a kimeneti CSV fájlhoz
- A köztes CSV fájlok automatikus törlése tömörítés után

## Használat

```bash
./csv_writer --source input.json --destination output.csv [--zip]
```

### Parancssori argumentumok

- `-s` vagy `--source`: A bemeneti JSON fájl elérési útja (kötelező)
- `-d` vagy `--destination`: A kimeneti CSV fájl elérési útja (kötelező)
- `-z` vagy `--zip`: A kimeneti CSV fájl ZIP tömörítésének engedélyezése (opcionális)

## Példa

JSON fájl konvertálása CSV-be:

```bash
./csv_writer --source adatok.json --destination adatok.csv
```

JSON fájl konvertálása CSV-be és tömörítése ZIP fájlba:

```bash
./csv_writer --source adatok.json --destination adatok.csv --zip
```

## Bemeneti JSON formátum

Az eszköz a következő formátumú JSON adatokat várja:

```json
[
  {
    "mező1": "érték1",
    "mező2": "érték2",
    ...
  },
  {
    "mező1": "érték3",
    "mező2": "érték4",
    ...
  }
]
```

A tömbben lévő minden objektum egy sorrá konvertálódik a CSV fájlban, az objektum kulcsai pedig oszlopfejlécekké válnak.

