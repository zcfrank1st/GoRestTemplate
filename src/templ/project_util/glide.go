package project_util


func GlideTemplate () string {
    return `package: .
import:
- package: github.com/gin-gonic/gin
- package: golang.org/x/sys
  repo: https://github.com/golang/sys.git
- package: github.com/go-ini/ini
  version: v1.28.0
- package: github.com/go-sql-driver/mysql
  version: v1.3
`
}