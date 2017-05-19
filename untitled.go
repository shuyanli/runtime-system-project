package main



import "fmt"                           //import format package 
import "net/http"                      //for server implementation
import "github.com/stackimpact/stackimpact-go"


func handler(w http.ResponseWriter, r *http.Request) {
      fmt.Fprintf(w, "Hello world!")
}

func main() {
    agent := stackimpact.NewAgent()
    agent.Start(stackimpact.Options{
      AgentKey: "agent key here",
      AppName: "Basic Go Server",
      AppVersion: "1.0.0",
      AppEnvironment: "production",
    })


    //========= start a server ===========

    // use MeasureHandlerFunc or MeasureHandler to additionally measure HTTP request execution time.
    http.HandleFunc(agent.MeasureHandlerFunc("/", handler))   //instead of using http.HandleFunc("/", handler1) 
    http.ListenAndServe(":8080", nil)

}