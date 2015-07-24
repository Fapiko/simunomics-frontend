package main
import (
	"net/http"
	"golang.org/x/net/html"
	"log"
	"net/url"
	"io/ioutil"
	"bytes"
	"launchpad.net/xmlpath"
	"errors"
)

func main() {
//	http.Handle("/authenticate", authenticateHandler)
	http.ListenAndServe(":7890", nil)
}

//func authenticateHandler(w http.ResponseWriter, r *http.Request) {
//
//}

func authenticate(username, password string) {
	data := url.Values{}
	data.Set("UserName", username)
	data.Set("Password", password)

	response, err := http.PostForm("http://www.simunomics.com/Login.php", data)
	panicError(err)

	found, err := loadXpath(response, "//*[@id=\"BadPassInner\"]/text()[1]")
	panicError(err)

	found = bytes.TrimSpace(found)

	log.Print(string(found))
}

func loadXpath(response *http.Response, xpath string) ([]byte, error) {
	body, err :=  ioutil.ReadAll(response.Body)
	panicError(err)

	// Parse body to see if login worked
	//	reader := strings.NewReader(body)
	root, err := html.Parse(bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	var b bytes.Buffer
	html.Render(&b, root)
	fixedHtml := b

	//	body = bytes.NewReader(fixedHtml)
	xmlroot, xmlerr := xmlpath.ParseHTML(bytes.NewReader(fixedHtml.Bytes()))

	if xmlerr != nil {
		return nil, xmlerr
	}

	path := xmlpath.MustCompile(xpath)
	if value, ok := path.Bytes(xmlroot); ok {
		return value, nil
	}

	return nil, errors.New("Could not find xpath")
}

func panicError(err error) {
	if (err != nil) {
		log.Panic(err)
	}
}