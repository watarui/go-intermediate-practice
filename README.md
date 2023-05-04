# go-intermediate-practice

https://techbookfest.org/product/jXDAEU1dR53kbZkgtDm9zx?productVariantID=dvjtgpjw8VDTXNqKaanTVi
の勉強記録

# 初期設定

```sh
docker volume create go-intermediate-practice-volume
cp .env.example .env
```

# コマンド

VSCode のコマンドパレットを開いて Run Task してください

# DB

```sh
mysql -udocker -hlocalhost -P3306 -Dsampledb -pdocker --protocol=tcp
```

# OpenID Connect

https://developers.google.com/identity/openid-connect/openid-connect?hl=ja#sendauthrequest

```
https://accounts.google.com/o/oauth2/v2/auth?client_id=hoge&response_type=id_token&scope=openid%20profile&redirect_uri=http://localhost:8081/callback&nonce=0394852-3190485-2490358
```
