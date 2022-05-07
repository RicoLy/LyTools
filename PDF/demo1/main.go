package main

import "github.com/jung-kurt/gofpdf"

func GeneratePdf(filename string) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.CellFormat(190, 7, "Welcome to topgoer.com", "0", 0, "CM", false, 0, "")
	pdf.ImageOptions(
		"./PDF/demo1/topgoer.jpg",
		80, 20,
		0, 0,
		false,
		gofpdf.ImageOptions{ImageType: "JPG", ReadDpi: true},
		0,
		"",
	)
	return pdf.OutputFileAndClose(filename)
}

func main() {
	err := GeneratePdf("hello.pdf")
	if err != nil {
		panic(err)
	}
}
