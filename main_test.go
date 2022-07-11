package ribes

import (
    "io/ioutil"
    "net/http"
    "testing"
)

func TestRibes(t *testing.T) {
    t.Log("start to test ribes")
    res, err := http.Get("https://azusachino.cn")
    if err != nil {
        t.Fail()
    }
    defer res.Body.Close()
    body, _ := ioutil.ReadAll(res.Body)
    t.Log(string(body))
}
