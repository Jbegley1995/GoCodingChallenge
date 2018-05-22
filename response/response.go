//Package response high level wrapper for writing a json response.
package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//Normally I would write a response package which doesn't repeat itself at all,
//but for the coding challenge it should be ok, also would normally have a struct for handling errors.

//OK write an OK status with the data passed.
func OK(w http.ResponseWriter, data interface{}) {
	jsonResp, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(jsonResp))
}

//InternalServerError Write an internal server error with the error passed.
func InternalServerError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	jsonResp, _ := json.Marshal(map[string]interface{}{"Error": err.Error()})
	fmt.Fprintf(w, string(jsonResp))
}

//UnprocessableEntity Write an error with the error passed.
func UnprocessableEntity(w http.ResponseWriter, data string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnprocessableEntity)
	jsonResp, _ := json.Marshal(map[string]interface{}{"Error": data})
	fmt.Fprintf(w, string(jsonResp))
}
