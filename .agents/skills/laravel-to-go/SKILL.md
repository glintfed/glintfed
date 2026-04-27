---
name: laravel-to-go
description: 將 Laravel 專案遷移到 Go 應用程式的實作指南，特別適用於分析 Laravel migrations、Eloquent 模型、既有 MySQL schema、routes 與 controllers，並規劃 Go 端的資料庫遷移、schema 對齊、資料回填、handler/routing 遷移與驗證流程。當任務涉及 Laravel 到 Go 的資料庫設計轉譯、migration 策略、資料相容性檢查、Pixelfed API/controller/route 行為遷移、或以既有 Laravel/Pixelfed schema 作為 Go 服務資料層基準時使用。
---

# Laravel To Go

在 glintfed 中，Laravel 到 Go 的遷移分成兩條互補工作線：資料庫/schema 遷移，以及 route/controller 到 Go handler 的 API surface 遷移。資料庫遷移要先從 Pixelfed 的 Laravel model 類別建立 ent schema 骨架，再根據已啟動的 MySQL service 取得真實 schema，將 tables 與 columns 補註回 ent models。route/controller 遷移要先從 Laravel routes 與 controller action 建立 Go handler surface，再註冊 handler dependencies，最後才綁定 chi routes。不要跳過其中任一步。

## 核心原則

1. 先找 canonical schema，再寫 Go migration；不要只看單一 migration 檔就直接重建。
2. 優先維持資料庫行為相容，而不是追求 Laravel 寫法的逐行翻譯。
3. 先做 additive change，再做 destructive change；任何刪欄位、改型別、改索引都要先證明可安全回滾或可重新建構。
4. 把 migration、backfill、驗證拆成獨立步驟；不要把 schema 變更與資料修補混在同一個不可重跑的流程。
5. 對於這個專案，若需要確認 Pixelfed/Laravel 的實際資料庫狀態，優先使用 `.pixelfed-local/` 啟動的本地服務驗證，不要修改唯讀的 `pixelfed/` 內容。

## 工作流程

### 0. 遷移 Laravel routes 與 controllers 到 Go handlers

當任務涉及 Pixelfed API routing、controller action、或 `internal/server/handler` 時，先建立 handler surface 與 DI 註冊，再處理實作細節。不要一開始就把所有 route 綁到不存在或未註冊的 handler 上。

來源與對照方式：

- 以 `pixelfed/routes/*.php` 與 `pixelfed/app/Http/Controllers/**/*.php` 作為 Laravel route/controller 的主要來源；`pixelfed/` 只讀，不可修改。
- 先整理每條 route 的 HTTP method、完整 path、controller class、action method、middleware/group prefix，再決定 Go 端 handler package 與 method。
- Laravel route group 的 prefix 要在 Go chi route 中展開或用 `mux.Route`/`r.Route` 保留結構；Laravel `{id}` 這類 path parameter 可直接對應 chi 的 `{id}`。
- 同一個 Laravel controller 若 action surface 很大，可依現有 Go 目錄慣例切成具語意的 handler package，例如 `api/apiv1`、`api/apiv1dot1`、`api/apiv2`、`groups/*`、`stories/storyapiv1`。

Go handler package 慣例：

- 每個 handler package 放在 `internal/server/handler/.../handler.go`，package name 使用小寫且符合目錄語意。
- 對外暴露 `type Handler interface { ... }`，列出該 package 負責的 route action。
- 提供 `func New() Handler { return &handler{} }` 與未導出的 `type handler struct{}`。
- 每個 action 方法簽名固定為 `func (h *handler) Action(w http.ResponseWriter, r *http.Request)`。
- 方法名稱使用 Go MixedCaps，通常由 Laravel controller action 轉換而來；不要使用底線。
- 尚未實作業務邏輯時保留 stub，但要建立 trace span：`internal.T.Start(r.Context(), "Domain.Action")`，span name 使用穩定且可搜尋的 handler domain 與 action。

Handler 聚合與 DI 註冊：

- `internal/server/handler/handler.go` 是 Go handler surface 的聚合點。新增 handler package 後，必須把它加入 `APIHandlers` struct 與 `NewAPIHandlers` constructor。
- `APIHandlers` 欄位名稱要符合 `internal/server/api.go` route 綁定會使用的名稱，例如 `APIv1`、`APIv1Dot1`、`GroupAPI`、`GroupPost`、`StoryAPIv1`。
- `NewAPIHandlers` 的參數順序、struct 欄位順序與回傳 literal 順序應保持一致，避免 kessoku 產生與人工檢查時難以追蹤。
- 遇到 package name 衝突時使用明確 alias，例如 `groupsapi`、`groupscreate`、`admindomainblocks`；不要同時 import 同一路徑兩次。
- `cmd/glintfed/kessoku.go` 必須為每個 handler interface 加上 `kessoku.Bind[package.Handler](kessoku.Provide(package.New))`，並保留 `kessoku.Provide(handler.NewAPIHandlers)`。
- 更新 kessoku provider graph 後執行 `go generate ./cmd/glintfed` 重新產生 `kessoku_band.go`；若只被要求修改特定檔案，先遵守使用者限制，不要順手改 generated file。

Route 綁定慣例：

- `internal/server/api.go` 只負責把已存在且已註冊的 handler 綁到 chi router；不要在這裡建立 handler 或放業務邏輯。
- 綁 route 前先確認 `APIHandlers` 有對應欄位、handler interface 有對應 method、kessoku 有對應 bind。
- Laravel 的 `Route::get/post/put/patch/delete` 分別對應 chi 的 `Get/Post/Put/Patch/Delete`。
- 對於相同 prefix 的大量 route，優先使用 `mux.Route("/prefix", func(r chi.Router) { ... })` 保留結構。
- 若使用者要求「先不要綁定路由」，只能建立 handler package、`APIHandlers` 聚合與 kessoku 註冊，不要修改 `internal/server/api.go` 的 route 清單。

驗證：

- 至少跑與修改範圍相符的 Go 編譯檢查，例如 `go test ./internal/server/handler`、`go test ./internal/server` 或 `go test ./cmd/glintfed`。
- 若更新 kessoku graph，執行 `go generate ./cmd/glintfed` 後再跑 `go test ./cmd/glintfed`。
- 檢查 `rg "NewAPIHandlers|APIHandlers|kessoku.Bind" internal/server/handler cmd/glintfed internal/server/api.go`，確保 handler surface、DI 與 route 綁定沒有脫節。

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
- 資料庫任務的第一優先輸出不是 SQL migration，而是可對齊 live schema 的 ent models。
- API 任務的第一優先輸出不是完整業務實作，而是可對齊 Laravel route/controller surface 的 Go handler、DI 註冊與 route mapping。

## 輸出格式

處理這類任務時，優先輸出：

1. 現況摘要：找到哪些 Laravel models、哪些 live tables，以及缺口在哪裡。
2. ent 建模結果：哪些 ent models 已建立、哪些是補建的。
3. 差異清單：Laravel model、live schema、ent model 之間的差異與風險。
4. 驗證計畫：如何確認 ent models 已完整反映 live MySQL schema。

處理 route/controller 遷移任務時，優先輸出：

1. route/controller 對照：Laravel route、controller action、Go handler package/method 的映射。
2. handler surface 結果：新增或更新了哪些 `Handler` interface、stub method 與 trace span。
3. DI 與 route 狀態：哪些 handler 已加入 `APIHandlers`、kessoku provider graph，哪些 route 已綁定或依使用者要求暫不綁定。
4. 驗證結果：執行過哪些 `go test`、`go generate` 或靜態檢查，以及剩餘缺口。

如果使用者要求直接實作，先依上述兩步流程完成 ent models 建立與補註，再處理後續 migration 或資料層程式碼。
