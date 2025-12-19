package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"router-gostashlg-template/entities/common/logger"

	"strings"

	"github.com/gin-gonic/gin"
)

func RequestLogger(ctx *gin.Context) {
	var bodyBytes []byte

	// buffer info sebagai builder yang menampung string yang ditulis ke buffer untuk nantinya dibuild menjadi satu string dan akan digunakan sebagai log
	var bufferInfo strings.Builder
	bufferInfo.WriteString(fmt.Sprintf("[REQUEST] %s, Path: %s, Request ID: %s, Request Info: \n", ctx.Request.Method, ctx.Request.URL.String(), ctx.Request.RemoteAddr))

	// Tulis semua header dari request ke buffer
	for name, values := range ctx.Request.Header {
		for _, value := range values {
			bufferInfo.WriteString(fmt.Sprintf("%s: %s\n", name, value))
		}
	}

	// Cek dulu ada query params atau tidak. Karena disini tidak ada pengecekan method, panjang lagi urusan kondisi nya jika ada
	queryParams := ctx.Request.URL.Query()
	if len(queryParams) > 0 {
		bufferInfo.WriteString("Raw Query: ")
		bufferInfo.WriteString(ctx.Request.URL.RawQuery)
	}

	// Cek juga, apakah ada body atau tidak.. kalo ada read dulu kemudian masukan lagi ke request body untuk dapat digunakan untuk handle berikutnya
	// Jika tidak dimasukan ke request body, read all ini akan mengosongkon request body, sehingga handler berikutnya yang membutuhkan request body
	// akan menerima kosong
	bodyBytes, er := io.ReadAll(ctx.Request.Body)
	if er != nil {
		logger.PrintError("Could not read request body")
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if len(bodyBytes) > 0 {
		// Tulis ke buffer raw dari body request, untuk informasi raw yang diterima sebelum di olah.
		bufferInfo.WriteString("Body: ")
		//! Kembalikan ke context agar handler selanjutnya masih bisa nerima isi dari body
		ctx.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	}
	bufferInfo.Write(bodyBytes)

	// jika ada query params, lakukan beautify printing
	if len(queryParams) > 0 {
		bufferInfo.WriteString("\n>> Parse RawQuery:\n")
		for key, values := range queryParams {
			for _, value := range values {
				bufferInfo.WriteString(fmt.Sprintf("%s: %s\n", key, value))
			}
		}
	}

	if ctx.ContentType() == ContentTypeAppJson {
		// Karena body berupa json, agar lebih efisien jangan menggunakan methode unmarshal, yang kemudian nantinya di marshal indent,
		// Cukup menggunakan json.Indent(), untuk itu butuh json source yang berupa byte, yang nantinya di append ke bytes.Buffer.
		// Untuk itu siapkan juga variable bytes.Buffer nya
		var prettyJson bytes.Buffer
		if er = json.Indent(&prettyJson, bodyBytes, "", "    "); er == nil {
			bufferInfo.WriteString(fmt.Sprintf("\n>> Beautify:\n%s\n", prettyJson.String()))
		}
	} else if ctx.ContentType() == ContentTypeAppForm {
		// Jika yang dikirim form, parse form kemudian lakukan beautify printing
		if er = ctx.Request.ParseForm(); er == nil {
			urlValues := ctx.Request.Form
			bufferInfo.WriteString("\n>> Parse Form:\n")
			for key, values := range urlValues {
				for _, value := range values {
					bufferInfo.WriteString(fmt.Sprintf("%s: %s\n", key, value))
				}
			}
		}
	}

	logger.PrintLog(bufferInfo.String())
}
