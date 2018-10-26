package design

import (
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("goa simple sample", func() {
	// APIのタイトル
	Title("multi-pattern-api")
	// APIの説明
	Description("様々なリソースアクセスを想定したAPI")
	// 対応しているプロトコル定義、httpかhttpsまたはその両方
	Scheme("http", "https")
	// 全てのエンドポイントのベースパス
	// /usersというエンドポイントがあったら、/api/v1/usersとなる
	BasePath("/multi/pattern")

	//　CORSポリシーの定義
	Origin("http://localhost:8080/swagger", func() {
		// 1つまたは複数の許可されたHTTPメソッド
		Methods("GET", "POST", "PUT", "DELETE")
		// Access-Control-Allow-Credentialsヘッダーを設定する
		Credentials()
	})
})
