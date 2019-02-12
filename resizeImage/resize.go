package main

import (
	"flag"

	"github.com/Sirupsen/logrus"
	"github.com/disintegration/imaging"
)

func Resize(fileName, destinationFile string, height, width int) {
	tempOpen, err := imaging.Open(fileName)
	if err != nil {
		logrus.Panic(err)
	} else {
		dstImage := imaging.Resize(tempOpen, height, width, imaging.Lanczos)
		err = imaging.Save(dstImage, destinationFile)
		if err != nil {
			logrus.Panic(err)
		}
	}
}

func main() {
	var filename string
	var destinationfile string
	var height, width int
	flag.StringVar(&filename, "s", "", "Source file")
	flag.StringVar(&destinationfile, "d", "", "Destination File")
	flag.IntVar(&height, "h", 800, "Height")
	flag.IntVar(&width, "w", 600, "Width")
	flag.Parse()
	if filename == "" && destinationfile == "" {
		logrus.Warningln("need source file and destination file (-help for more information)")
	} else {
		Resize(filename, destinationfile, height, width)
	}
}
