# simple short link go

简单的短链接后端，需手动修改 [links.json](links.json) ，支持热更新链接。

未来会部署到 https://q8p.cc

```
git clone https://github.com/bddjr/simple-short-link-go
cd simple-short-link-go
./run.sh
```

短链名称不得包含 `/` `\` `.`，开头 `@` 用于后端内部重定向
