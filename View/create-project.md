# 備忘録

## プロジェクトを新規作成する
1. コンテナを立ち上げコンテナ内に入る
    ```
    docker-compose run --rm node sh
    ```
    | コマンド<br>オプション | 説明 |
    |:-:|:--|
    | run | コンテナ起動 |
    | --rm [コンテナ名] | コンテナ実行後に削除。デタッチド・モードの場合は無視 |
    | [コマンド] | 実行コマンド |

2. nuxtプロジェクトを作成

コンテナ内でcreate-nuxt-appコマンドを実行する
```
/var/www # yarn global add create-nuxt-app

yarn global v1.15.2
[1/4] Resolving packages...
[2/4] Fetching packages...
[3/4] Linking dependencies...
[4/4] Building fresh packages...
warning Your current version of Yarn is out of date. The latest version is "1.22.0", while you're on "1.15.2".
info To upgrade, run the following command:
$ curl --compressed -o- -L https://yarnpkg.com/install.sh | bash
success Installed "create-nuxt-app@2.14.0" with binaries:
      - create-nuxt-app
Done in 12.15s.


```
エラーがでなければOK


create-nuxt-appを実行して、Nuxtプロジェクトを作成
色々、質問されるけど基本的にはプロジェクト名だけでOK
```
/var/www # create-nuxt-app ./

? Project name ⇦ここにはプロジェクト名（今回はPortfolioView）を入力
? Project description　⇦後はひたすらEnter
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
Warning: name can no longer contain capital letters
⠹ Installing packages with yarn


🎉  Successfully created project PortfolioView

  To get started:

        yarn dev

  To build & start for production:

        yarn build
        yarn start

// コンテナを抜ける
/var/www # exit
```

コンテナ起動
```
docker-compose up workspace
```

http://localhost:3000　にアクセスして確認

## Dockerが無い場合
1. node.jsをインストール
2. npx をインストール
3. 上記の***2. nuxtプロジェクトを作成***から行う

## Gitから落とした場合
1. src直下にnode_modulesがあれば削除する
2. docker 起動
```
$ docker-compose run --rm node sh
/var/www # 

```
3. yarn nuxt を取り込む
```
/var/www # yarn global add create-nuxt-app
yarn global v1.15.2
[1/4] Resolving packages...
[2/4] Fetching packages...
[3/4] Linking dependencies...
[4/4] Building fresh packages...
warning Your current version of Yarn is out of date. The latest version is "1.22.4", while you're on "1.15.2".
info To upgrade, run the following command:
$ curl --compressed -o- -L https://yarnpkg.com/install.sh | bash
success Installed "create-nuxt-app@2.14.0" with binaries:
      - create-nuxt-app
Done in 11.96s.
```

4. npm install
```
/var/www # npm install

```

5. exit 

6. docker-compose up workspace