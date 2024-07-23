package main

import (
	"bufio"
	"fmt"
	"image"
	"image/png"
	"os"
	"path/filepath"
	"regexp"
	"time"

	"github.com/fatih/color"
	"github.com/google/uuid"
	"github.com/skip2/go-qrcode"
)

const (
	qrSize = 512  // n*n pixels
	qridSizeLimit = 8 // uid max len
	saveDirPermCode = 0755 // Permission code for `saved` dir (and subdirs)
)

func cleanData(name string) string {
	reg := regexp.MustCompile(`[^\w\-]`)
	return reg.ReplaceAllString(name, "")
}

func saveQRCode(img image.Image, filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	return png.Encode(f, img)
}

func generateQRCode() error {
	// Colorful logging configuration
	errorColor := color.New(color.FgRed).Add(color.Bold)

	// Current date (yyyy/mm/dd)
	year := time.Now().Year()
	month := fmt.Sprintf("%02d", int(time.Now().Month()))
	day := fmt.Sprintf("%02d", time.Now().Day())

	// Input params for QR generation
	fmt.Print("Enter data to encode in QR code: ")
	reader := bufio.NewReader(os.Stdin)
	data, err := reader.ReadString('\n')
	if err != nil {
		errorColor.Println("Error reading input:", err)
		return err
	}
	data = data[:len(data)-1]

	// QR code generation
	qr, err := qrcode.New(data, qrcode.Medium)
	if err != nil {
		errorColor.Println("Error while generating:", err)
		return err
	}
	img := qr.Image(qrSize)

	// Create directory to save QR`s
	saveDir := fmt.Sprintf("saved/%d-%s-%s/", year, month, day)
	if _, err := os.Stat(saveDir); os.IsNotExist(err) {
		err = os.MkdirAll(saveDir, saveDirPermCode)
		if err != nil {
			errorColor.Println("Error creating directory:", err)
			return err
		}
	}

	// Generate unique id for QR
	uniqueQRid := uuid.New().String()[:qridSizeLimit]
	safeData := cleanData(data)
	currentDate := fmt.Sprintf("%d%s%s", year, month, day)
	filename := filepath.Join(saveDir, fmt.Sprintf("QR-%s-%s-%s.png", safeData, uniqueQRid, currentDate))

	// Saving QR
	err = saveQRCode(img, filename)
	if err != nil {
		errorColor.Println("Error during saving QR code:", err)
	}

	return nil
}

func main() {
	if err := generateQRCode(); err != nil {
		color.New(color.FgRed).Add(color.Bold).Println("QR code generation failed.")
	} else {
		color.New(color.FgGreen).Add(color.Bold).Println("QR code generation succeeded.")
	}
}

