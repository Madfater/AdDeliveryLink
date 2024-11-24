# AdDeliveryLink

## 問題陳述

設計並實現一個簡化的廣告投放服務，這個服務將提供兩個主要的 API：

admin API：用於創建廣告並設定廣告的條件

public API：用於根據特定條件列出符合條件的廣告

## 解決方案概述

- 創建一個資料庫

  - 將 condition 中的可複數資料的資料使用多對多的模式進行儲存
  - 其他皆為一般的 attribute 儲存在同一個資料表

- admin API：將根據用戶提供的資料，生成廣告並將其存儲到資料庫中

  - 假設用戶並未輸入，使用預設值進行儲存
  - Gender：以 F 及 M 分別代表女性及男性。預設值將為 NULL 以代表皆可
  - Age：有區間故使用 AgeStart、AgeEnd 區分上下限。預設值分別為 1 以及 100
  - Platform、Country：使用多對多的模式進行儲存。預設將不會儲存於聯合資料表中

- public API：將接收用戶的查詢，根據查詢條件從資料庫中檢索符合條件的廣告並返回給用戶
  - 除了 offset 以及 limit 皆為 optional
  - 廣告為一個小時刷新過期時間，以免造成 Memory Cache 快速過期

## 資料庫設計

<div align=center>
	<img src="./img/schema.png">
	<img src="./img/ER_model.png">
</div> 
 
-	對 ID、AgeEnd、AgeStart、Gender 設置 index 提高搜尋速度

## 技術選擇

### Tech Stack

- [Go](https://go.dev)
- [MySQL](https://www.mysql.com/)
- [Redis](https://redis.io)

### Library

- [go-redis](https://github.com/redis/go-redis)
- [Gin](https://github.com/gin-gonic/gin)
- [Gorm](https://gorm.io/)
- [go-gorm/cache](https://github.com/go-gorm/caches)
- [Validator](https://github.com/go-playground/validator)
- [gin-swagger](https://github.com/swaggo/gin-swagger)

### 使用原因

- Validator：它是用來驗證輸入的 Library。僅需要在 struct 中設定好限制就能方便的驗出錯誤
- go-redis：redis 的客戶端
- go-gorm/cache：Gorm 的 Plugin，只需實作 Get、Store 及 Invalidate，就能讓 Gorm 自行處裡 Memory cache 的邏輯。能使用上者實作，讓其變成使用 redis 作為快取伺服器
- gin-swagger: API 文檔自動生成

## 流量測試

### 測試環境

<div align=center>
	<h3>硬體配置</h3>
	<img src="./img/env.png">
	<h3>Linux 環境</h3>
	<img src="./img/linux_env.png">
</div>

### 測試結果

- 以下測試結果為使用 Jmeter 進行測試
- 測試方式為 10 分鐘的負載測試
- [Jmeter 的詳細測試報告](https://madfater.github.io/AdDeliveryLink/)

<div align=center>
	<h3>測資</h3>
	<img src="./img/test.png">
	<p>除了 limit 和 offset 皆為隨機</p>
	<h3>結果圖表</h3>
	<img src="./img/TPS.png">
	<p>Transaction Per Second</p>
	<img src="./img/TRT.png">
	<p>Response Times Over Time</p>
</div>

## 運行專案

### 方法 1 : clone 專案

1. 確保你已經在本地環境中安裝了 Docker 和 Docker Compose
2. 確保你已經 clone 了專案並在終端中切換到專案目錄
3. 執行以下命令來啟動容器：

```bash
docker compose up --build
```

4. /swagger/index.html 觀看 API 文檔，確保正確啟動

### 方法 2 : 使用 docker image

1. 確保你已經在本地環境中安裝了 Docker 和 Docker Compose
2. 確保你已經複製下方的 Docker Compose 到本地
3. 執行以下命令來啟動容器：

```bash
docker compose up
```
4. /swagger/index.html 觀看 API 文檔，確保正確啟動

```yml
version: "3"
services:
  app:
    image: ghcr.io/madfater/app:main
    environment:
      GIN_MODE: release
      MYSQL_USERNAME: root
      MYSQL_IP: db
      MYSQL_PORT: 3306
      MYSQL_Database: DcardAssignment
    ports:
      - "8080:8080"
    depends_on:
      - db
      - redis
    restart: always
  db:
    image: ghcr.io/madfater/db:main
    volumes:
      - db-db:/var/lib/mysql
      - db-conf:/etc/mysql/conf.d
      - db-logs:/logs
    environment:
      MYSQL_DATABASE: DcardAssignment
      MYSQL_ALLOW_EMPTY_PASSWORD: "true"
    ports:
      - "3306:3306"
  redis:
    image: redis:5.0
    volumes:
      - redis-data:/data
    ports:
      - "6379:6379"
volumes:
  redis-data:
  db-db:
  db-conf:
  db-logs:
```
