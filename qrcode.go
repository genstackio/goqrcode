package goqrcode

import (
	svg "github.com/ajstarks/svgo"
	"github.com/boombuler/barcode/qr"
	"image/color"
	"io"
	"math"
	"strings"
)

var DefaultFilledBlockStyle = "fill:black"
var DefaultNotFilledBlockStyle = "fill:white"

//goland:noinspection GoUnusedExportedFunction
func GenerateAndStreamQrCode(w io.Writer, config Config) {
	qrCode, _ := qr.Encode(config.Data, qr.H, qr.Auto)
	s := svg.New(w)
	blockSize := config.BlockSize
	if 0 == blockSize {
		blockSize = 4
	}
	offset := config.Offset
	if 0 == offset {
		offset = 4
	}
	if config.DisableOffset {
		offset = 0
	}
	if math.Abs(config.Scale) > 0.00001 {
		blockSize = int(float64(blockSize) * config.Scale)
	}
	width := qrCode.Bounds().Max.X
	firstX := blockSize * offset
	firstY := blockSize * offset

	ns := []string{}
	if len(config.Style) > 0 {
		ns = append(ns, "style=\""+strings.ReplaceAll(config.Style, "\"", "\\\"")+"\"")
	}
	s.Start(width*blockSize+(blockSize*(offset*2)), width*blockSize+(blockSize*(offset*2)), ns...)

	currY := firstY

	filledBlockStyle, notFilledBlockStyle := buildBlockStylesFromConfig(config)

	for x := 0; x < width; x++ {
		currX := firstX
		for y := 0; y < width; y++ {
			switch qrCode.At(x, y) {
			case color.Black:
				if len(filledBlockStyle) > 0 {
					s.Rect(currX, currY, blockSize, blockSize, filledBlockStyle)
				}
			case color.White:
				if len(notFilledBlockStyle) > 0 {
					s.Rect(currX, currY, blockSize, blockSize, notFilledBlockStyle)
				}
			}
			currX += blockSize
		}
		currY += blockSize
	}

	s.End()
}

func buildBlockStylesFromConfig(config Config) (string, string) {
	filledBlockStyle := config.BlockStyle
	if len(filledBlockStyle) == 0 {
		filledBlockStyle = DefaultFilledBlockStyle
	}
	if "-" == filledBlockStyle {
		filledBlockStyle = ""
	}
	notFilledBlockStyle := config.EmptyBlockStyle
	if len(notFilledBlockStyle) == 0 {
		notFilledBlockStyle = DefaultNotFilledBlockStyle
	}
	if "-" == notFilledBlockStyle {
		notFilledBlockStyle = ""
	}

	return filledBlockStyle, notFilledBlockStyle
}
