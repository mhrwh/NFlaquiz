## 初期設定 (Initialization)
1. ファイルを追加 (add file, directry.)
```
NFlauiz/
　└ api/
　│  └ .air.toml
　└ ui/
     └ src/
        └ assets/
           └ img/
```

2. 自分の環境でnode_modules を再インストールして以下に配置 (Reinstall node_modules in your environment and place it below.)
    * [node_modulesの再インストール方法](https://zenn.dev/mo_ri_regen/articles/node-modules-article)
    * node version -> 16.14
```
NFlauiz/
　└ ui/
　 　└ node_modules/
```


## 使い方
起動 (Start)
```
docker-compose up -d
```

起動確認 (Check it's activated.)
```
docker-compose ps
```

終了 (Stop)
```
docker-compose stop
```

コンテナ内に入る (Entering the Container)
```
docker-compose exec コンテナ名 sh
```

docker-compose.ymlの内容を変更後，dbのボリュームが変更後のバージョンとコンフリクトするのを防ぐためにボリュームを削除する


ボリューム一覧 (voloume list)
```
docker volume ls
```

ボリュームの削除 (Delete volume)
```
docker volume rm "project_db-data"
```

volumeの全削除 (Delete all volumes)
```
docker volume rm $(docker volume ls -qf dangling=true)
```


ブラウザでアクセス (Access with a browser)
```
lcalhost:ポート番号
```
|  コンテナ  |  ポート番号  |
| ---- | ---- |
|  ui(Vue3)  |  8080  |
| api(Gin) | 8888 |
|  db  |  3306  |
|  phpMyAdmin  |  4040  |
