# Glintfed

Glintfed 是一個以 Go 實作的服務，與 Pixelfed 專案在 API 與 Database 層級相容。

## 開發規範

- Go 的程式風格遵循 `.agents/skills/go-style-guide` 中的指示。
- 專案結構遵循 `.agents/skills/go-project-layout` 中的指示。
- service / repo 層級可以使用 model 層級的封裝進行資料庫存取，避免直接依賴或操作 `*ent.Client`；`*ent.Client` 的查詢細節應優先收斂在 `internal/model`。
- `pixelfed/` 資料夾是 [pixelfed/pixelfed](https://github.com/pixelfed/pixelfed) 的既有實作，該資料夾是**唯讀 READ ONLY**的

## 應用程式分層

應用程式請求順序一般為 `main.go` -> `server` -> `handler` -> (`service` / `repo` / `model`)。

- `server` 雖然目前只有 HTTP 服務，但未來可以視需要提供 gRPC 或 worker pool（例如 watermill）相關的初始化工作。
- `handler` 的職責為「驗證請求」與「構建回應」，其中可以視需要調用 `service` / `repo` 或 `model`。
- `service` 可以視需要使用 `repo` / `model`，但為了避免 cycle import，`repo` 及 `model` 不可以使用 `service`；同理，`repo` 也可以使用 `model`。
- `model` 目前僅作為 `*ent.Client` 的高階抽象，實際的 SQL 或資料存取邏輯會被寫在這裡；但未來不排除會有其它資料源（例如 MongoDB、Elasticsearch）的抽象。
    - 如果需要為 `*ent.*` 擴充功能，可以直接在 `ent/` 資料夾下寫入新檔案：例如，想要為 `*ent.Profile` 加入新的函式 `URL()`，可以在 `ent/profile_url.go` 中撰寫相關邏輯。
- `service`, `repo` 或 `model` 都會註冊到 kessoku 以注入應用程式中。
    - interface 的實作可能來源於多個不同的 servces, repo 或 model，這時可以靠 anonymous struct 來解決這個問題
    ```go
    kessoku.Bind[features](kessoku.Provide(
        func (svc1 *svc1, svc2 *svc2) features {
            return &struct{
                s1 *svc1
                s2 *svc2
            }{
                s1: svc1,
                s2: svc2,
            }
        }
    ))
    ```

## 本地 Pixelfed 環境

- 專案提供 `.pixelfed-local/` 作為本地開發用的 Pixelfed docker compose 工作目錄。
- 啟動方式為 `docker compose -f .pixelfed-local/docker-compose.yml up -d`。
- 停止方式為 `docker compose -f .pixelfed-local/docker-compose.yml down`。
- 查看服務狀態可使用 `docker compose -f .pixelfed-local/docker-compose.yml ps`。
- 查看應用日誌可使用 `docker compose -f .pixelfed-local/docker-compose.yml logs -f pixelfed`。
- `.pixelfed-local/` 內的檔案屬於本地產生或本地環境設定，不應修改 `pixelfed/` 內的唯讀內容來取代它。
- 未來若需要確認 Pixelfed API 的實際行為、驗證 request/response、檢查 migration 後的 schema，或直接查看 DB 內資料，應優先使用 `pixelfed-local` services 進行確認。
- 若需要檢查資料庫內容，可直接透過 `pixelfed-local` 的 MySQL service 連線；若需要驗證背景工作、排程或快取行為，也可直接使用同一組 `pixelfed-local` 的 `horizon`、`scheduler`、`redis` services。
