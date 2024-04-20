package main
import(
	"fmt"
	"log"
	"net/http"
	"github.com/Anjana-Rose/GolangToDoApp/router"
)

func main(){
	r:=router.Router()
	fmt.Println("starting the server...")
	log.Fatal(http.ListenAndServe(":9000",r))
}