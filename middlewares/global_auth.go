package middlewares

import (
	"fmt"
    "reflect"
    "io/ioutil"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
    //sjs "github.com/bitly/go-simplejson"
)

func GlobalAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		u := uuid.Must(uuid.NewV4())
		ctx.Set("_self_context__request_id", fmt.Sprintf("%v", u))

        var args map[string]interface{}
        args = make(map[string]interface{})

        //store form data params
        ctx.Request.ParseMultipartForm(32<<20)
        for k, v := range ctx.Request.PostForm {
            args[k] = v[0]
            fmt.Printf("=============> k:%v, v:%v, t:%v", k, args[k], reflect.TypeOf(v))
        }

        body := ctx.Request.Body
        x, _ := ioutil.ReadAll(body)

        //js, _ := sjs.NewJson(x)

        //fmt.Printf("=============#######: %s t:%v\n", string(x), reflect.TypeOf(x), js.Get("nickname").MustString(), reflect.TypeOf(js))
        fmt.Printf("=============#######: %s t:%v\n", string(x), reflect.TypeOf(x))

		ctx.Next()
	}
}
