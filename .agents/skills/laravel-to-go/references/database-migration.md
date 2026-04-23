# 資料庫遷移參考

## 目錄

- Laravel 到 Go 的資料庫遷移目標
- 來源盤點順序
- Laravel 常見 schema 特性對照
- 風險檢查清單
- 建議的遷移切法
- 驗證方法

## Laravel 到 Go 的資料庫遷移目標

目標不是把 Laravel migration API 逐字翻譯成另一套 migration DSL，而是保留以下能力：

- 資料可被既有 Laravel/Pixelfed 與新 Go 服務共同理解。
- schema 可以由 Go 端穩定重建。
- migration 可以重跑、排序清楚、可在 CI 或本地環境驗證。
- 在切換期間，資料回填與讀寫兼容路徑可被明確觀察。

## 來源盤點順序

先依序檢查下列來源：

1. `database/migrations/`：看歷史演進，不假設最後狀態一定正確。
2. Eloquent model 與 relation：看 table、primary key、casts、hidden/appends、soft delete。
3. seeders/factories：看資料形狀與隱含預設值。
4. 實際資料庫：用 `SHOW CREATE TABLE`、`information_schema`、索引與 constraint 查詢確認最終狀態。
5. 真正的讀寫程式碼：看 controller、service、query builder、raw SQL。

若 Laravel 專案存在手動 SQL、hotfix migration、或曾經直接改 DB，live DB 幾乎一定比 migration 檔更可靠。

## Laravel 常見 schema 特性對照

### 主鍵與整數欄位

- `increments` 通常對應 `INT UNSIGNED AUTO_INCREMENT PRIMARY KEY`。
- `bigIncrements` 通常對應 `BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY`。
- `foreignId` 在 MySQL 常見為 `BIGINT UNSIGNED`，必須和被參照欄位一致。
- Go 端若使用 ORM 或 migration 工具，先確認是否真的支援 unsigned；若不支援，要明確記錄風險。

### 時間欄位

- Laravel 常見 `created_at`、`updated_at`、`deleted_at` 為 nullable timestamp/datetime。
- 不要假設 `updated_at` 一定有 DB 層 `ON UPDATE CURRENT_TIMESTAMP`；很多專案其實由應用層更新。
- 如果 Go 端會自己管理時間欄位，需確認與既有行為一致，尤其是零值、timezone、precision。

### Soft delete

- `softDeletes()` 實質上通常只是新增 `deleted_at`。
- 真正的刪除語意在應用層 query scope，不在 schema 本身。
- Go 端若沒有實作對應的 default filter，很容易把邏輯刪除資料誤當有效資料。

### 多型關聯

- `morphs()` 常見拆成 `<name>_type` 與 `<name>_id`。
- `nullableMorphs()` 允許兩欄皆為 null。
- 這類欄位通常需要複合索引，不一定有 foreign key。
- Go 端不要嘗試強制轉成單一外鍵模型，除非使用者明確要求重構資料模型。

### Pivot tables

- Laravel pivot table 經常沒有獨立 ID，而是複合唯一鍵。
- 要確認是否包含額外欄位，例如排序、狀態、建立時間。
- 若 Go 端 ORM 預設假設單欄主鍵，需先避開錯誤建模。

### JSON 與序列化欄位

- Laravel casts 可能把 DB `json`、`text`、甚至 PHP serialized blob 包成陣列或物件。
- 只看 migration 無法知道真實資料格式；要抽樣看 live data。
- Go 端若改用 struct decode，先確認欄位內容是否真的一致。

### Enum 與狀態欄位

- Laravel `enum()` 在 MySQL 可行，但跨工具與跨資料庫可攜性差。
- 遷移到 Go 時，優先考慮：
  - 保留字串欄位，並在應用層驗證合法值。
  - 或建立 lookup table / check constraint。
- 若使用者沒有要求重構，先維持現況相容，再規劃第二階段清理。

### 預設值與 nullability

- Laravel migration 的 `default()` 不代表 live DB 一定相同。
- 某些專案會先允許 null，之後靠 backfill 再改成 not null。
- Go migration 寫法若直接一步到位，可能在既有資料上失敗。

### 索引與 constraint

- 檢查 unique index 是否同時承擔業務規則。
- 檢查複合索引欄位順序是否對應實際查詢條件。
- 檢查 foreign key 是否真的存在；很多 Laravel 專案只建欄位與索引，沒有建 FK。
- 檢查 collation 與 charset，尤其是帳號、slug、username、domain 等比對欄位。

## 風險檢查清單

看到以下情況時，要先停下來補證據：

- migration 檔案與 live schema 不一致。
- 欄位型別看起來相同，但 signed/unsigned、precision、collation 不一致。
- 應用層有 raw SQL 或 DB trigger，但 migration 沒有表達。
- 既有資料中存在 null、空字串、非法 enum 值或壞資料。
- Laravel 依賴 model cast、accessor、mutator 才能正確解讀資料。
- 同一張表同時被 API、背景工作、排程任務讀寫。

## 建議的遷移切法

### 階段 1：建立事實

- 匯出 live schema。
- 記錄與 migration 檔差異。
- 列出高風險欄位與高風險表格。

### 階段 2：建立 Go migration 基線

- 先讓 Go migration 能重建「現況 schema」。
- 不在這一階段順便做命名優化或資料模型重構。
- 如果需要，可把歷史 migration 壓成少數幾個 baseline migration。

### 階段 3：新增相容欄位或索引

- 以 additive 方式新增新欄位、新索引、新 table。
- 保留舊欄位與舊讀寫契約。

### 階段 4：回填資料

- 用可重跑的程式或 job 做 backfill。
- 為 backfill 記錄範圍、批次大小、重試方式與完成條件。

### 階段 5：切換讀寫

- 先切 read path，再切 write path，或反過來，但要明確。
- 若是雙寫，定義一致性檢查與停止條件。

### 階段 6：清理舊契約

- 只有在所有依賴都切換完成後，才刪除舊欄位、舊索引、舊表格。
- destructive migration 要獨立成最後階段。

## 驗證方法

### Schema 驗證

- 在本地用 `.pixelfed-local/` 啟動 Pixelfed 服務。
- 對 MySQL 執行 schema 匯出，確認表格、欄位、索引、constraint。
- 在 Go 測試資料庫重建 migration，與 live schema 對照。

### 資料抽樣驗證

- 抽查高風險欄位：JSON、enum、polymorphic、soft delete、時間欄位。
- 抽查高流量表：users、posts、media、follows、notifications 之類的核心表。
- 比對 Laravel 與 Go 對同一筆資料的序列化結果。

### 行為驗證

- 驗證建立、更新、刪除、列表查詢與排序。
- 驗證背景工作是否仍能理解新欄位或新索引。
- 驗證 rollback 或重新初始化資料庫時 migration 能否穩定重建。

### 輸出建議

在分析結果中至少列出：

1. 目前採信的 canonical schema 來源。
2. 會先實作的 Go migration 範圍。
3. 需要另外做 backfill 的欄位與原因。
4. 會用哪些查詢或 API 驗證遷移成功。
