// テストをルート直下に配布しているのは以下の理由による
//   - cycle importedが表示される
//   - router.LoadHTMLGlobで指定しているパスがルートからの相対的な部分が解釈されない
package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"gowebapi/server"
	"github.com/stretchr/testify/assert"
)

func TestGetRouter_v1_ping(t *testing.T) {

    // リクエスト(実際にリクエストはしない)
    response := httptest.NewRecorder()
    request, _ := http.NewRequest(http.MethodGet, "/api/v1/ping", nil)

    // リクエスト実行
    router := server.GetRouter()
    router.ServeHTTP(response, request)

    /* テスト結果確認 */
    assert.EqualValues(t, http.StatusOK, response.Code)
    assert.EqualValues(t, "{\"message\":\"pong\"}", response.Body.String())

}

