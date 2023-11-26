package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"sync"
)

func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}
func (app *application) indexHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		http.Error(w, "Error parsing form", http.StatusInternalServerError)
		return
	}
	numInputs, err := strconv.Atoi(r.FormValue("numInputs"))
	// Loop through the dynamically created input fields
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusInternalServerError)
		return
	}
	data := PageVariables{
		Title:  "Book",
		Images: []ImagePair{},
	}
	var wg sync.WaitGroup
	for i := 1; i <= numInputs; i += 2 {
		inputName1 := fmt.Sprintf("input_%d", i)
		inputValue1 := r.FormValue(inputName1)
		//resultChan := make(chan []byte, numInputs)
		wg.Add(1)
		payload1 := make(map[string]interface{})
		payload1["inputs"] = inputValue1
		go app.queryAPI(payload1, &wg, i)
		inputName2 := fmt.Sprintf("input_%d", i+1)
		inputValue2 := r.FormValue(inputName2)
		//resultChan := make(chan []byte, numInputs)
		wg.Add(1)
		payload2 := make(map[string]interface{})
		payload2["inputs"] = inputValue2
		go app.queryAPI(payload2, &wg, i+1)
		newImg := ImagePair{
			FrontImagePath: "static/img/saved_image" + strconv.Itoa(i) + ".png",
			FrontTitle:     inputValue1,
			BackImagePath:  "static/img/saved_image" + strconv.Itoa(i+1) + ".png",
			BackTitle:      inputValue2,
		}
		data.Images = append(data.Images, newImg)
	}
	wg.Wait()
	files := []string{
		"ui/html/index.layout.tmpl",
		"ui/html/base.layout.tmpl",
		"ui/html/footer.partial.tmpl",
	}
	tmpl, err := template.ParseFiles(files[0], files[1], files[2])
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Error parsing template file", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		return
	}
}
func (app *application) getuserinput(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"ui/html/userinput.layout.tmpl",
		"ui/html/base.layout.tmpl",
		"ui/html/footer.partial.tmpl",
	}
	tmpl, err := template.ParseFiles(files[0], files[1], files[2])
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Error parsing template file", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, &PageVariables{
		Form: New(nil),
	})
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		return
	}
}
