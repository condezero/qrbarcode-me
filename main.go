package main

import (
	"bytes"
	"fmt"
	"image/png"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/oned"
	"github.com/makiuchi-d/gozxing/qrcode"
)

type CodeRequest struct {
	Content string `json:content`
	Width   int    `json:width`
	Height  int    `json:height`
}
type CodeType int16

const (
	BarCode CodeType = 0
	QrCode           = 1
)

func init() {
	file := fmt.Sprintf(".env%s", os.Args[1])
	if err := godotenv.Load(file); err != nil {
		log.Printf("No %s file found", file)
		godotenv.Load()
	}
	mode := os.Getenv("GIN_MODE")
	gin.SetMode(mode)
}

func main() {
	r := setupRouter()
	r.Run(os.Getenv("PORT"))
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/generate/barcode", func(c *gin.Context) {
		generateCode(c, BarCode)
	})
	r.POST("/generate/qrcode", func(c *gin.Context) {
		generateCode(c, QrCode)
	})
	return r
}
func generateCode(c *gin.Context, codeType CodeType) {
	payload, err := deserializePayload(c)

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if len(strings.TrimSpace(payload.Content)) == 0 {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	switch codeType {
	case BarCode:
		data, err := getBarCode(*payload)
		if err != nil {
			c.AbortWithError(500, err)
		}
		c.Data(http.StatusOK, "application/octet-stream", data)

	case QrCode:
		data, err := getQrCode(*payload)
		if err != nil {
			c.AbortWithError(500, err)
		}
		c.Data(http.StatusOK, "application/octet-stream", data)
	}
}

func getBarCode(payload CodeRequest) ([]byte, error) {

	writer := oned.NewCode128Writer()

	img, err := writer.Encode(payload.Content, gozxing.BarcodeFormat_CODE_128, payload.Width, payload.Height, nil)
	if err != nil {
		log.Fatalf("impossible to encode barcode: %s", err)
		return nil, err
	}

	var buf bytes.Buffer

	err = png.Encode(&buf, img)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func getQrCode(payload CodeRequest) ([]byte, error) {
	writer := qrcode.NewQRCodeWriter()

	img, err := writer.Encode(payload.Content, gozxing.BarcodeFormat_QR_CODE, payload.Width, payload.Height, nil)
	if err != nil {
		log.Fatalf("impossible to encode qrcode: %s", err)
		return nil, err
	}
	var buf bytes.Buffer

	err = png.Encode(&buf, img)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func deserializePayload(c *gin.Context) (*CodeRequest, error) {
	payload := new(CodeRequest)
	if err := c.Bind(payload); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return nil, err
	}
	return payload, nil
}
