package main


import(
	"io"
	"net/http"
)

func verificationHandler(w http.ResponseWriter, r *http.Request){
	verification := "{\"statusText\":\"OK\", \"name\":\"Moonshard\"}"
	io.WriteString(w, verification)
}












