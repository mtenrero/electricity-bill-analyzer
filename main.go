package electricity

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", FullMeanAnalysis)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
