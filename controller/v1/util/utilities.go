package util

import (
	"io/ioutil"
	"fmt"
	"time"
    "strconv"
    "github.com/gin-gonic/gin"
)

func DisplaySleepAction(c *gin.Context){

	// デフォルトは文字列取得なのでstrconv.AtoiでINT型に変換。戻りは2つの値を返すので、２番目は不要なので_を指定
	second, _ := strconv.Atoi(c.Query("sec"))
	nanosecond, _ := strconv.Atoi(c.Query("nsec"))

	// secとnsecが指定されないといずれも0になるので、その場合は3秒スリープに固定する
	if second == 0 && nanosecond == 0 {
		second = 3
	}

    // cf. https://pkg.go.dev/time
    time.Sleep(time.Duration(nanosecond))            // ns は、ナノ秒
    time.Sleep(time.Duration(second) * time.Second)  // 秒スリープ
    s := fmt.Sprintf("Sleep sec=%d, nsec=%d", second, nanosecond)
    c.String(200, s)
}

func PrintParamsAction(c *gin.Context){
	// cf. https://stackoverflow.com/questions/54390031/how-to-read-variables-from-post-payload-in-gin-gonic
	data, _ := ioutil.ReadAll(c.Request.Body)
	c.String(200, string(data))
}
