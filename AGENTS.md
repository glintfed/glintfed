# Glintfed

Glintfed 是一個以 Go 實作的服務，與 Pixelfed 專案在 API 與 Database 層級相容。

## 開發規範

- Go 的程式風格遵循 `.agents/skills/go-style-guide` 中的指示。
- 專案結構遵循 `.agents/skills/go-project-layout` 中的指示。
- `pixelfed/` 資料夾是 [pixelfed/pixelfed](https://github.com/pixelfed/pixelfed) 的既有實作，該資料夾是**唯讀 READ ONLY**的

## 本地 Pixelfed 環境

- 專案提供 `.pixelfed-local/` 作為本地開發用的 Pixelfed docker compose 工作目錄。
- 啟動方式為 `docker compose -f .pixelfed-local/docker-compose.yml up -d`。
- 停止方式為 `docker compose -f .pixelfed-local/docker-compose.yml down`。
- 查看服務狀態可使用 `docker compose -f .pixelfed-local/docker-compose.yml ps`。
- 查看應用日誌可使用 `docker compose -f .pixelfed-local/docker-compose.yml logs -f pixelfed`。
- `.pixelfed-local/` 內的檔案屬於本地產生或本地環境設定，不應修改 `pixelfed/` 內的唯讀內容來取代它。
- 未來若需要確認 Pixelfed API 的實際行為、驗證 request/response、檢查 migration 後的 schema，或直接查看 DB 內資料，應優先使用 `pixelfed-local` services 進行確認。
- 若需要檢查資料庫內容，可直接透過 `pixelfed-local` 的 MySQL service 連線；若需要驗證背景工作、排程或快取行為，也可直接使用同一組 `pixelfed-local` 的 `horizon`、`scheduler`、`redis` services。
