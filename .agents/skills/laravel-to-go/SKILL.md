---
name: laravel-to-go
description: 將 Laravel 專案遷移到 Go 應用程式的實作指南，特別適用於分析 Laravel migrations、Eloquent 模型與既有 MySQL schema，並規劃 Go 端的資料庫遷移、schema 對齊、資料回填與驗證流程。當任務涉及 Laravel 到 Go 的資料庫設計轉譯、migration 策略、資料相容性檢查、或以既有 Laravel/Pixelfed schema 作為 Go 服務資料層基準時使用。
---

# Laravel To Go

在 glintfed 中，Laravel 到 Go 的資料庫遷移分成兩個明確步驟：先從 Pixelfed 的 Laravel model 類別建立 ent schema 骨架，再根據已啟動的 MySQL service 取得真實 schema，將 tables 與 columns 補註回 ent models。不要跳過其中任一步。

## 核心原則

1. 先找 canonical schema，再寫 Go migration；不要只看單一 migration 檔就直接重建。
2. 優先維持資料庫行為相容，而不是追求 Laravel 寫法的逐行翻譯。
3. 先做 additive change，再做 destructive change；任何刪欄位、改型別、改索引都要先證明可安全回滾或可重新建構。
4. 把 migration、backfill、驗證拆成獨立步驟；不要把 schema 變更與資料修補混在同一個不可重跑的流程。
5. 對於這個專案，若需要確認 Pixelfed/Laravel 的實際資料庫狀態，優先使用 `.pixelfed-local/` 啟動的本地服務驗證，不要修改唯讀的 `pixelfed/` 內容。

## 工作流程

### 1. 從 Laravel model 建立 ent schema 骨架

先掃描 `pixelfed/app/*.php` 與 `pixelfed/app/Models/*.php` 中所有 `extends Model` 或 `extends Authenticatable` 的 PHP 檔案，並以 `ent new {model_name}` 建立對應的 ent models。這一步的目標是先把 Laravel 端的 model surface 映射成 Go/ent 的初始骨架，不要求一次補齊所有欄位。

需要自動建立 ent schema 骨架時，直接執行 [scripts/generate_ent_from_laravel_models.sh](scripts/generate_ent_from_laravel_models.sh) 並傳入 Laravel app root。

執行這一步時遵循以下要求：

- 只把 `extends Model` 或 `extends Authenticatable` 的類別視為候選來源。
- 以 PHP 檔名作為 `ent new` 的 model 名稱來源。
- 把 Laravel model 視為 ent schema 的命名種子，不要在這一步先做大量重命名。
- 讀取 Laravel model 內的 `$hidden` 定義；凡是出現在 `hidden` 內的欄位，後續在 ent 欄位上應標示為 `Sensitive()`。
- 若有 model 名稱衝突、保留字或明顯不符合 Go/ent 命名慣例，再單獨處理。

### 2. 以 live MySQL schema 補註 ent models

在 `.pixelfed-local/` 的 docker mysql service 已啟動的前提下，直接從 live database 取得完整的 tables 與 columns，將真實 schema 補註到 ent models 中。這一步的基準是資料庫的實際狀態，不是僅靠 Laravel migration 檔推測。

執行這一步時至少要完成：

- 盤點所有 tables、columns、型別、nullability、default、index、foreign key 與必要的備註。
- 將已存在的 ent models 補上對應欄位與必要設定。
- 若發現某些 tables 沒有對應的 ent model，使用 `go tool ent new {model_name}` 補建對應 ent models，再繼續補欄位。
- 若 Laravel model 與 live table 對不上，以 live MySQL schema 為主，並在分析中標記差異。

需要從 `.pixelfed-local/` 的 MySQL service 查詢特定 table 或批次列出 Laravel model 對應 table 定義時，使用 [scripts/describe_pixelfed_tables.sh](scripts/describe_pixelfed_tables.sh)。

### 3. 對齊 Laravel 語意與資料庫真實語意

在補註 ent models 時，不要只複製欄位名稱，要一併確認 Laravel 與資料庫的語意：

- `increments`, `bigIncrements` 是否對應正確的主鍵型別、unsigned 與 auto increment。
- `foreignId`, `unsignedBigInteger` 是否與 referenced key 完全一致。
- Laravel model 中若某欄位存在於 `$hidden`，則對應的 ent field 應加上 `Sensitive()`，不要因為 live DB 看不出這個語意就遺漏。
- `softDeletes` 是否需要反映 `deleted_at` 與對應查詢慣例。
- `timestamps`, `timestampsTz` 的精度、default 與更新方式是否與 live schema 一致。
- `json`, `text`, `enum`, polymorphic relation 等欄位是否需要特殊處理。

細節與檢查清單見 [references/database-migration.md](references/database-migration.md)。

### 4. 驗證 ent models 完整性

至少執行以下檢查：

- 所有 live tables 都有對應 ent models，或已明確標示為暫不處理。
- 所有 ent models 的欄位定義都能對應到 live columns。
- 所有 Laravel model `$hidden` 欄位都已在對應 ent fields 上標示 `Sensitive()`。
- 高風險欄位如 JSON、enum、soft delete、polymorphic relation 已被標記或處理。
- Laravel model、live schema、ent model 三者的差異有被列出，而不是默默忽略。

## 本專案特別要求

- `pixelfed/` 是唯讀，僅作為既有實作參考，不要修改它來完成遷移。
- 需要確認實際 schema、migration 結果、request/response、背景工作或快取行為時，優先使用 `.pixelfed-local/` 的 docker compose 服務。
- 若發現 Laravel migration、模型定義與 live DB 有落差，應先在分析中列出差異，再提出 Go 端採信的基準。
- 這個 skill 在 glintfed 的第一優先輸出不是 SQL migration，而是可對齊 live schema 的 ent models。

## 輸出格式

處理這類任務時，優先輸出：

1. 現況摘要：找到哪些 Laravel models、哪些 live tables，以及缺口在哪裡。
2. ent 建模結果：哪些 ent models 已建立、哪些是補建的。
3. 差異清單：Laravel model、live schema、ent model 之間的差異與風險。
4. 驗證計畫：如何確認 ent models 已完整反映 live MySQL schema。

如果使用者要求直接實作，先依上述兩步流程完成 ent models 建立與補註，再處理後續 migration 或資料層程式碼。
