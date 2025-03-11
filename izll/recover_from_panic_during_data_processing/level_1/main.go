package level_1

import (
	"math/rand"
	"os"
	"recover_from_panic_during_data_processing/writer"
	"strings"
	"time"
)

func Run() {
	bw := writer.NewBatchWriter(os.Stdout, 100)
	defer bw.Close()

	randomString := generateRandomString(16384)
	err := bw.Write(randomString)
	if err != nil {
		println("Error writing data")
		return
	}
}

func generateRandomString(length int) string {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	rand.New(rand.NewSource(time.Now().UnixNano()))

	var sb strings.Builder
	sb.Grow(length)

	for i := 0; i < length; i++ {
		randomIndex := rand.Intn(len(chars))
		sb.WriteByte(chars[randomIndex])
	}

	return sb.String()
}
