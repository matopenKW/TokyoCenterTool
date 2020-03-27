### コマンド一覧

```
docker-compose run --rm node sh

```

```
一括停止
docker stop $(docker ps -q)

全コンテナ削除
docker rm $(docker ps -q -a)

全イメージ削除
docker rmi $(docker images -q)

```