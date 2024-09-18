package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
)

// Convert PDF to images using pdfcpu
func convertPDFToImages(pdfPath string, outputDir string) error {
	conf := pdfcpu.NewDefaultConfiguration()
	conf.Cmd = pdfcpu.IMAGE
	conf.FileName = pdfPath
	conf.OutDir = outputDir

	if _, err := api.Process(conf); err != nil {
		return err
	}
	return nil
}

// Extract text from image using tesseract
func extractTextWithOCR(imagePath string) (string, error) {
	cmd := exec.Command("tesseract", imagePath, "stdout")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func main() {
	pdfPath := "path/to/your/file.pdf"
	outputDir := "output/images"

	// Create output directory if it doesn't exist
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		log.Fatalf("Failed to create output directory: %v", err)
	}

	// Convert PDF to images
	if err := convertPDFToImages(pdfPath, outputDir); err != nil {
		log.Fatalf("Failed to convert PDF to images: %v", err)
	}

	// Extract text from each image
	images, err := filepath.Glob(filepath.Join(outputDir, "*.png"))
	if err != nil {
		log.Fatalf("Failed to list images: %v", err)
	}

	var extractedText string
	for _, imagePath := range images {
		text, err := extractTextWithOCR(imagePath)
		if err != nil {
			log.Printf("Failed to extract text from image %s: %v", imagePath, err)
			continue
		}
		extractedText += text + "\n"
	}

	fmt.Println("Extracted text:")
	fmt.Println(extractedText)
}
