package main



import "fmt"                           //import format package 
import "net/http"                      //for server implementation
import "github.com/stackimpact/stackimpact-go"




func main() {
   agent := stackimpact.NewAgent();
agent.Start(stackimpact.Options{
  AgentKey: "d7400ebf4d95f2b0f49c25f829336044c4ff0cb7",
  AppName: "testServer",
})

    //========= start a server ===========

    // use MeasureHandlerFunc or MeasureHandler to additionally measure HTTP request execution time.
    http.HandleFunc(agent.MeasureHandlerFunc("/", handler))   //instead of using http.HandleFunc("/", handler1) 
    
    http.HandleFunc("/hello", handler2)  // in /hello directory

    http.ListenAndServe(":8080", nil)

}


func handler(w http.ResponseWriter, r *http.Request) {
      fmt.Fprintf(w, "This is the main hello")
}
func handler2(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "hello world")
}