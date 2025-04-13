# 実行

go run main.go

# 検証

- curl http://localhost:8080/no-recover
→ Goアプリごと プロセスがクラッシュ（サーバー落ちる）
- curl http://localhost:8080/with-recover
→ panic は発生するが recover でキャッチされ、サーバーは落ちない
