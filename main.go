
package main


import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/intel/fastgo/compress/gzip"
)


func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [option] <file>\n", os.Args[0])
		flag.PrintDefaults()
	}

	decompress := flag.Bool("d", false, "Decompress the file")
	level := flag.Int("l", gzip.DefaultCompression, "Compression level (1-9)")
	flag.Parse()

	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}

	inputFile := flag.Arg(0)

	if *decompress {
		decompressFile(inputFile)	
	} else {
		compressFile(inputFile, *level)
	}
}


func compressFile(inputFile string, level int) {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("Failed to open input file: %v", err)
	}
	defer file.Close()

	outputFile := inputFile + ".gz"
	out, err := os.Create(outputFile)
	if err != nil {
		log.Fatalf("Failed to create output file: %v", err)
	}
	defer out.Close()

	gw, err := gzip.NewWriterLevel(out, level)
	if err != nil {
		log.Fatalf("Failed to create gzip writer: %v", err)
	}
	defer gw.Close()

	if _, err := io.Copy(gw, file); err != nil {
		log.Fatalf("Failed to compress file: %v", err)
	}

	fmt.Printf("Compressed %s to %s\n", inputFile, outputFile)
	
}

func decompressFile(inputFile string) {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("Failed to open input file: %v", err)
	}
	defer file.Close()

	gr, err := gzip.NewReader(file)
	if err != nil {
		log.Fatalf("Failed to create gzip reader: %v", err)
	}
	defer gr.Close()

	outputFile := inputFile[:len(inputFile)-3]  //Remove .gz extension
	out, err := os.Create(outputFile)
	if err != nil {
		log.Fatalf("Failed to create output file: %v", err)
	}
	defer out.Close()

	if _, err := io.Copy(out, gr); err != nil {
		log.Fatalf("Failed to decompress file: %v", err)
	}

	fmt.Printf("Decompressed %s to %s\n", inputFile, outputFile)
}
