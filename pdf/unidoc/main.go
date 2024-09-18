package main

import (
	"fmt"
	"github.com/unidoc/unipdf/v3/common"
	"github.com/unidoc/unipdf/v3/extractor"
	"github.com/unidoc/unipdf/v3/model"
	"log"
	"os"
)

func extractTextFromPDF(filePath string) (string, error) {
	common.SetLogger(common.NewConsoleLogger(common.LogLevelDebug))

	// Open the PDF file.
	f, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	// Create a new PDF reader.
	pdfReader, err := model.NewPdfReader(f)
	if err != nil {
		return "", err
	}

	var text string
	numPages, err := pdfReader.GetNumPages()
	if err != nil {
		return "", err
	}

	for i := 1; i <= numPages; i++ {
		page, err := pdfReader.GetPage(i)
		if err != nil {
			return "", err
		}

		extractor, err := extractor.New(page)
		if err != nil {
			return "", err
		}

		pageText, err := extractor.ExtractText()
		if err != nil {
			return "", err
		}

		text += pageText
	}

	return text, nil
}

func main() {
	filePath := "./colly/files/2.pdf"
	text, err := extractTextFromPDF(filePath)
	if err != nil {
		log.Fatalf("Failed to extract text from PDF: %v", err)
	}
	fmt.Println("Extracted text:")
	fmt.Println(text)
}
