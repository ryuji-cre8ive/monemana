# LINEで金銭管理を簡単に

## 環境開発

### .env

- LINE_BOT_CHANNEL_SECRET
- LINE_BOT_CHANNEL_TOKEN
- DATABASE_URL(postgres)

### migrate

dbconfig.ymlのdatasourceにdbのURLをはる

## LINE BOT 使用

想定人数: 2人
（今後のアップデートで使いやすくするつもりではいます）

アイテムを追加する際
`タイトル お金`
集計結果を確認する
`集計`

## ER図
![Monemana v2](https://github.com/ryuji-cre8ive/monemana/assets/49904836/f55922ed-8eeb-44e9-aae1-bb3422bf138e)
