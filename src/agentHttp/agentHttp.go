package agentHttp

import (
	"net/http"
	"sync"

	agentHttpController "github.com/JieanYang/HelloWorldGoAgent/src/agentHttp/agentHttpController"
)

func StartHttp() {
	wg := new(sync.WaitGroup)

	wg.Add(1)
	go func() {
		http.HandleFunc("/", agentHttpController.HomeController)
		// http.HandleFunc("/RunCommandWithFormData", agentHttpController.RunCommandWithFormData)
		// http.HandleFunc("/RunCommandWithBody", agentHttpController.RunCommandWithBody)
		http.HandleFunc("/RunCommandByScriptContent", agentHttpController.RunCommandByScriptContent)
		http.HandleFunc("/RunCommandWithUrl", agentHttpController.RunCommandByUrl)

		err := http.ListenAndServe(":9001", nil) // Block code
		if err != nil {
			panic(err)
		}
		wg.Done()
	}()

	wg.Wait()

}
