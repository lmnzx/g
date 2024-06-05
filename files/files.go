package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /photo.jpg", handlePhotoJPG)
	mux.HandleFunc("GET /hello.pdf", handleHelloPDF)

	http.ListenAndServe(":4567", mux)
}

func handlePhotoJPG(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("derek.jpg")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	w.Header().Set("Content-Disposition", "attachment; filename=custom-name.jpg")
	w.Header().Set("Content-Type", "image/jpeg")

	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleHelloPDF(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/pdf")

	pdfContent := []byte(`%PDF-1.0
%µ¶
1 0 obj
<</Type/Catalog/Pages 2 0 R>>
endobj
2 0 obj
<</Kids[3 0 R]/Count 1/Type/Pages/MediaBox[0 0 595 792]>>
endobj
3 0 obj
<</Type/Page/Parent 2 0 R/Contents 4 0 R/Resources<<>>>>
endobj
4 0 obj
<</Length 58>>
stream
q
BT
/F1 96 Tf
1 0 0 1 36 684 Tm
(Hello!) Tj
ET
Q
endstream
endobj
xref
0 5
0000000000 65536 f 
0000000016 00000 n 
0000000062 00000 n 
0000000136 00000 n 
0000000209 00000 n 
trailer
<</Size 5/Root 1 0 R>>
startxref
316
%%EOF
`)

	_, err := w.Write(pdfContent)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
