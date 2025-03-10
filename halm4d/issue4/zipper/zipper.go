package zipper

import (
	"archive/zip"
	"csv-exporter/config"
	"fmt"
	"io"
	"os"
)

type Zipper struct {
	conf      *config.Config
	outputZip string
}

func NewZipper(conf *config.Config) *Zipper {
	return &Zipper{
		conf:      conf,
		outputZip: fmt.Sprintf("%s.zip", conf.Output()),
	}
}

func (z *Zipper) ZipFile(fileToZip string) error {
	if z.conf.Verbose() {
		fmt.Println("creating zip...")
	}
	archive, err := os.Create(z.outputZip)
	if err != nil {
		return err
	}
	defer archive.Close()
	zipWriter := zip.NewWriter(archive)

	if z.conf.Verbose() {
		fmt.Println("opening file to zip...")
	}
	f1, err := os.Open(fileToZip)
	if err != nil {
		return err
	}
	defer f1.Close()

	if z.conf.Verbose() {
		fmt.Println("writing file to zip...")
	}
	w1, err := zipWriter.Create(fileToZip)
	if err != nil {
		return err
	}
	if _, err := io.Copy(w1, f1); err != nil {
		return err
	}

	if z.conf.Verbose() {
		fmt.Println("closing zip file...")
	}
	return zipWriter.Close()
}
