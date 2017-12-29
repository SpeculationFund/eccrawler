package crawlers



import (

    "github.com/nbio/st"
    "gopkg.in/h2non/gock.v1"
    "testing"
    "encoding/json"
//    "fmt"
)


func TestGet_Simple(t *testing.T) {
    defer gock.Off()

    gock.New("http://coincheck.com").
        Get("/api/ticker").
        Reply(200).
        JSON(map[string]string{"foo": "bar"})


    webapi := &C_WebAPI{}
    result := make(chan interface{})
    go webapi.Get("http://coincheck.com/api/ticker", result)

    st.Expect(t, string((<-result).([]byte))[:13], `{"foo":"bar"}`)

    st.Expect(t, gock.IsDone(), true)

}


func TestGet_WebAPI(t *testing.T) {
    defer gock.Off()

    webapi := &C_WebAPI{}
    result := make(chan interface{})

    cases := []struct{
        Host string
        Url string
        StatusCode int
        ResultMsg map[string]string
    }{
        {
            Host: "http://coincheck.com",
            Url: "/api/ticker",
            StatusCode: 200,
            ResultMsg: map[string]string{"foo": "bar"},
        },
        /*
        {
            Host: "xxx",
            Url: "/api/ticker",
            StatusCode: 200,
            ResultMsg: map[string]string{},
        },
        {
            Host: "http://coincheck.com",
            Url: "/api/ticker",
            StatusCode: 200,
            ResultMsg: nil,
        },
        */


    }

    for _, c := range cases {
        gock.New(c.Host).
            Get(c.Url).
            Reply(c.StatusCode).
            JSON(c.ResultMsg)
        ioresult, _ := json.Marshal(c.ResultMsg)
        go webapi.Get(c.Host+c.Url, result)
        st.Expect(t, (<-result).([]byte)[:len(ioresult)], ioresult)
    }

    st.Expect(t, gock.IsDone(), true)

}



