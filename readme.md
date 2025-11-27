# smartping

原作者已经多年不更新了，邮件发送只支持25号端口，而很多云主机封禁了25号端口。

改动：

- 增加SMTP 465端口支持
- 把项目改成 Go modules
- `github.com/mattn/go-sqlite3` 这个包换成 `modernc.org/sqlite` 不再依赖 libc 。

执行 build.sh 构建，产物路径：bin/smartping 。
