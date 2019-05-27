# ツッコミロボ

実行するには：

```
go run main.go
```

[ツッコミロボ動画](https://www.youtube.com/watch?v=XTnjc0cWiko&feature=youtu.be)

## タスク：APIルーターを実装しよう (server/router.js)

|Step|ステップ|
|---|---|
Read the files to make sure you understand how the data flows through the service | サービス側でどうデータが動くかを理解するため全てのファイルを見てみよう
Add logic to the `main.go` file to handle a HTTP Get Request to `/tsukkomi?phrase=${phrase}` | クライアントが `/tukkomi?phrase=${phrase}` へ HTTP GET Request を送りますので、ちゃんとフレーズにつっこんでレスポンスを送ってあげましょう。
Hit the API endpoint with POSTMAN to verify your work | POSTMANでAPIから正しいレスポンスが戻ってくるか確認しとこう

**本日のキーワード**

* [HTTP](http://www.atmarkit.co.jp/ait/articles/1703/29/news045.html)
* [REQUEST RESPONSE](https://itsakura.com/network-http-get-post)
* [GET REQUEST](https://wa3.i-3-i.info/word11405.html)

**役に立つリンク：**

* [Go package: net/http](https://golang.org/pkg/net/http/)
* [Go package: math/rand](https://golang.org/pkg/math/rand/)
