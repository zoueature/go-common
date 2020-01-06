/* +----------------------------+
 * | Author: Zoueature          |
 * +----------------------------+
 * | Email: zoueature@gmail.com |
 * +----------------------------+
 */
package httprequest

import (
	"io/ioutil"
	"log"
	"net/http"
	"testing"
	"time"
)

type handle struct {

}

func (h *handle)ServeHTTP(w http.ResponseWriter, r *http.Request)()  {
	if r.Method == http.MethodGet {
		queries := r.RequestURI
		_, _ = w.Write([]byte(queries))
	} else if r.Method == http.MethodPost {
		body, _ := ioutil.ReadAll(r.Body)
		_, _  = w.Write(body)
	}
}

func startHttpServe() {
	handle := &handle{}
	err := http.ListenAndServe("127.0.0.1:8080", handle)
	if err != nil {
		log.Fatal("start http server error: " + err.Error())
	}
}

func TestClient_Get(t *testing.T) {
	go startHttpServe()
	time.Sleep(1*time.Second)
	client := NewClient()
	response := client.Get("http://127.0.0.1:8080", map[string]string{"name":"hello", "age":"world"}, nil)
	if response.StatusCode != http.StatusOK {
		t.Error("response http code is not 200")
	}
	if response.Err != nil {
		t.Error("response error is not nil : " + response.Err.Error())
	}
	if response.Body != "/?age=world&name=hello" {
		t.Error("response body is not match, is : " +response.Body)
	}
}

func TestClient_Post(t *testing.T) {
	go startHttpServe()
	time.Sleep(1*time.Second)
	client := NewClient()
	response := client.Post("http://127.0.0.1:8080", map[string]string{"name":"hello", "age":"world"}, nil)
	if response.StatusCode != http.StatusOK {
		t.Error("response http code is not 200")
	}
	if response.Err != nil {
		t.Error("response error is not nil : " + response.Err.Error())
	}
	if response.Body != "name=hello&age=world" {
		t.Error("response body is not match, is : " +response.Body)
	}
}