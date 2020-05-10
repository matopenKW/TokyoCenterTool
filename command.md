# command reference

## Nuxt.js をビルドする
1. 環境に入る
```
docker-compose run --rm node sh
```
2. generateしてdistファイルの作成
```
npm run generate
```
3. exit

## firebase hosting deploy
1. projectの確認
    ```
    firebase use
    ```
2. projectが違う場合は下記で変更
    ```
    firebase use \ tokyocentertool
    ```
3. firebaseのローカルでテストを行う
    ```
    firebase serve --only hosting
    ```
4. deploy
    ```
    firebase deploy 
    ```