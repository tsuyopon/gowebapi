package util

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// 実際にはリクエストをせずに必要なオブジェクトを生成した後に、DisplayPingActionを直接呼び出すテスト
func TestDisplayPingAction_1(t *testing.T) {

	// リクエスト(実際にリクエストはしない)
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)                              // ここがポイントです。responseを引き渡してginに必要なcontextを生成します。
	c.Request, _ = http.NewRequest(http.MethodGet, "/api/v1/ping", nil)  // 後で直接DisplayPingActionを呼び出しているので実際にこの引数は意味がない

	// テスト実行
	DisplayPingAction(c)

	/* テスト結果確認 */
	// ステータスコード確認(以下の2つは同じ)
	assert.EqualValues(t, http.StatusOK, response.Code)
	assert.EqualValues(t, http.StatusOK, c.Writer.Status())

	// レスポンスボディ確認
	assert.EqualValues(t, "{\"message\":\"pong\"}", response.Body.String())

}
