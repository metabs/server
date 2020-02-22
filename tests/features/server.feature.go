package features

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/gherkin"
)

type serverFeature struct {
	res *http.Response
}

// ServerIsUpAndRunning Feature to check that the server is up and running
func ServerIsUpAndRunning(s *godog.Suite, ) {
	f := &serverFeature{
		res: &http.Response{},
	}
	s.Step(`^an HTTP "([^"]*)" request with the URI "([^"]*)"$`, f.anHTTPRequestWithTheURI)
	s.Step(`^an HTTP "([^"]*)" request with the URI "([^"]*)"$`, f.anHTTPRequestWithTheURI)
	s.Step(`^the server must reply with a status code (\d+)$`, f.theServerMustReplyWithAStatusCode)
	s.Step(`^the server must reply with a body:$`, f.theServerMustReplyWithABody)
}

func (f *serverFeature) anHTTPRequestWithTheURI(method, uri string) error {
	req, err := http.NewRequest(method, uri, nil)
	if err != nil {
		return fmt.Errorf("request creation failed. Due to: %s", err)
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("client failed. Due to: %s", err)
	}

	f.res = res

	return nil
}

func (f *serverFeature) theServerMustReplyWithAStatusCode(expectedStatusCode int) error {
	if f.res.StatusCode != expectedStatusCode {
		return fmt.Errorf("response status cose is wrong. Expected: %d, Given: %d", http.StatusOK, f.res.StatusCode)
	}

	return nil
}

func (f *serverFeature) theServerMustReplyWithABody(expectedResBody *gherkin.DocString) error {
	resBody, err := ioutil.ReadAll(f.res.Body)
	if err != nil {
		return fmt.Errorf("response body can't be read. Due to: %s", err)
	}

	if string(resBody) != expectedResBody.Content {
		return fmt.Errorf("response body is wrong. Expected: %s, Given: %s", expectedResBody.Content, resBody)
	}

	return nil
}
