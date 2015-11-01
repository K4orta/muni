package muni

import "testing"

func TestRequest(t *testing.T) {
	fakeServer := makeFakeServer()

	resp, err := transitRequest(fakeServer.URL + "/")
	if err != nil {
		t.Error(err)
	}

	defer resp.Body.Close()

}
