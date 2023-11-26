package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"image"
	"image/png"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"
)

const (
	apiURL        = "https://xdwvg9no7pefghrn.us-east-1.aws.endpoints.huggingface.cloud"
	authorization = "Bearer VknySbLLTUjbxXAXCjyfaFIPwUTCeRXbFSOjwRiCxsxFyhbnGjSFalPKrpvvDAaPVzWEevPljilLVDBiTzfIbWFdxOkYJxnOPoHhkkVGzAknaOulWggusSFewzpqsNWM"
)

func (app *application) queryAPI(payload map[string]interface{}, wg *sync.WaitGroup, i int) {
	defer wg.Done()

	startTime := time.Now()

	reqBody := bytes.NewBuffer([]byte{})
	err := json.NewEncoder(reqBody).Encode(payload)
	if err != nil {
		log.Println("Error encoding payload:", err)
		return
	}
	log.Println("Helll")
	req, err := http.NewRequest("POST", apiURL, reqBody)
	if err != nil {
		log.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Accept", "image/png")
	req.Header.Set("Authorization", authorization)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error making API request:", err)
		return
	}
	defer resp.Body.Close()
	imageBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return
	}
	elapsedTime := time.Since(startTime)
	log.Printf("API response time: %s", elapsedTime)
	img, _, err := image.Decode(bytes.NewReader(imageBytes))
	if err != nil {
		log.Fatal(err)
		return
	}
	var buf bytes.Buffer
	err = png.Encode(&buf, img)
	if err != nil {
		log.Fatal(err)
		return
	}
	imageBase64 := base64.StdEncoding.EncodeToString(buf.Bytes())
	imagename := "saved_image" + strconv.Itoa(i) + ".png"
	filename := filepath.Join("ui", "static", "img", imagename)

	imageData, err := base64.StdEncoding.DecodeString(imageBase64)
	if err != nil {
		log.Println("Hi1:", err)
		return
	}

	file, err := os.Create(filename)
	if err != nil {
		log.Println("Hi2:", err)
		return
	}
	defer file.Close()

	_, err = file.Write(imageData)
	if err != nil {
		log.Println("Hi3:", err)
		return
	}

	log.Printf("Image saved to: %s", imagename)

}

//func main() {
//	payloads := []map[string]interface{}{
//		{"inputs": "Astronaut riding a horse"},
//		{"inputs": "boy helps a girl"},
//		{"inputs": "Astronaut riding a dog"},
//		{"inputs": "Astronaut riding a buffalo"},
//		{"inputs": "boy riding a bullet"},
//		{"inputs": "Astronaut playing football"},
//		{"inputs": "Astronaut playing cricket"},
//		{"inputs": "boy batting"},
//		{"inputs": "girls riding a dog"},
//		{"inputs": "Astronaut beating a horse"},
//		{"inputs": "boy enjoying a girl"},
//		{"inputs": "Astronaut riding a ferrari"},
//	}
//	var wg sync.WaitGroup
//	var i = 0
//	for _, payload := range payloads {
//		wg.Add(1)
//		go queryAPI(payload, &wg, i)
//		i = i + 1
//	}
//	wg.Wait()
//}
