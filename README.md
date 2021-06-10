# EnvMonitoring

## 開發環境
- Ubuntu 20.04.2 LTS
- go1.16.3 linux/amd64

## 下載
點選右方Releases下載[最新版本](https://github.com/vincent87720/EnvMonitoring/releases)

## 設定
使用`settings.yaml`檔案設定要監聽的serial port和上傳資料的目標資料庫資訊
```yaml
---
port:
  name: /dev/ttyUSB0 #填入serial port名稱 ex:/dev/ttyUSB0、COM3
  baudRate: 9600 #設定序列傳輸速率
  dataBits: 8 #設定資料位元組的長度
  parity: ParityNone #設定同位位元，可以設定為ParityNone、ParityOdd、ParityEven、ParityMark、ParitySpace
  stopBits: Stop1 #設定停止位元的數目，可以設定為Stop1、Stop1Half、Stop2
database:
  host: 127.0.0.1:3306 #設定資料庫伺服器IP位址和埠號
  dbname: yourDatabaseName #設定資料庫名稱
  username: yourDatabaseUsername #設定資料庫使用者名稱
  password: yourDatabasePassword #設定資料庫使用者密碼
```

## 開始
開啟終端機，Windows版本執行`EnvMonitoring.exe`，Linux版本執行`EnvMonitoring`