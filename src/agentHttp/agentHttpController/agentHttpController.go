package agentHttpController

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/JieanYang/HelloWorldGoAgent/src/common/runCommand"
)

type Reponse struct {
	Results string
}

func HomeController(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Go - Hello World</h1>")
}

// func RunCommandWithFormData(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == http.MethodPost {
// 		err := r.ParseForm()
// 		if err != nil {
// 			http.Error(w, "Error parsing form data", http.StatusBadRequest)
// 			return
// 		}
// 		name := r.FormValue("name")

// 		fmt.Fprintf(w, "<h1>Paris Hello, %s</h1>", name)
// 	} else {
// 		fmt.Fprint(w, "<h1>Hellow Word</h1>")
// 	}

// }

// func RunCommandWithBody(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == http.MethodPost {
// 		var body struct{ Name string }
// 		err := json.NewDecoder(r.Body).Decode(&body)
// 		if err != nil {
// 			http.Error(w, "Error parsing JSON data", http.StatusBadRequest)
// 			return
// 		}

// 		output := runCommand.RunCommand()

// 		resultsObj := Reponse{Results: string(output)}
// 		data, err := json.Marshal(resultsObj)
// 		if err != nil {
// 			http.Error(w, "Error generate JSON reuslts", http.StatusBadRequest)
// 			return
// 		}

// 		w.Header().Set("Content-Type", "application/json")
// 		w.Write(data)

// 		// fmt.Fprintf(w, "<h1>Hello, %s</h1>", body.Name)
// 	} else {
// 		fmt.Fprint(w, "<h1>Hellow Word</h1>")
// 	}

// }

func RunCommandByScriptContent(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var body struct{ ScriptContent string }
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			http.Error(w, "Error parsing JSON data", http.StatusBadRequest)
			return
		}

		scriptOutput := runCommand.RunCommandByScriptContent(string(body.ScriptContent))

		resultsObj := Reponse{Results: string(scriptOutput)}
		data, err := json.Marshal(resultsObj)
		if err != nil {
			http.Error(w, "Error generate JSON reuslts", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	} else {
		fmt.Fprint(w, "<h1>Hellow Word</h1>")
	}

}

func RunCommandByUrl(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var bodyFromReq struct{ Url string }
		err := json.NewDecoder(r.Body).Decode(&bodyFromReq)
		if err != nil {
			http.Error(w, "Error parsing JSON data", http.StatusBadRequest)
			return
		}

		// Get content from url
		fmt.Println("bodyFromReq", bodyFromReq)
		responseFromScriptUrl, err := http.Get(bodyFromReq.Url)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer responseFromScriptUrl.Body.Close()
		scriptContent, err := ioutil.ReadAll(responseFromScriptUrl.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("string(scriptContent)", string(scriptContent))

		scriptOutput := runCommand.RunCommandByScriptContent(string(scriptContent))

		resultsObj := Reponse{Results: string(scriptOutput)}
		data, err := json.Marshal(resultsObj)
		if err != nil {
			http.Error(w, "Error generate JSON reuslts", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(data)

		// fmt.Fprintf(w, "<h1>Hello, %s</h1>", body.Name)
	} else {
		fmt.Fprint(w, "<h1>Not support</h1>")
	}
}
