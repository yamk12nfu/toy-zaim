## 認証トークンの取得

- `python auth/auth.py` を実行する
- `consumer key`と`consumer secret`を入力する(zaim にアプリ登録すると確認できる)
- `Please go here and authorize : hoge` hoge の部分をクリックして認証する
- 認証ページのソースを確認。`oauth verifier`の値を確認して入力
- `access token`と`access secret`が取得できるので`config.yml`に書いておく
