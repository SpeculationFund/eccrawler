package crawlers



import (

    "github.com/nbio/st"
    "gopkg.in/h2non/gock.v1"
    "testing"
)


func TestGet_WebAPI(t *testing.T) {
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
