# simple short link go

简单的短链接后端，需手动修改 [links.json](links.json) ，支持热更新链接。

```
git clone https://github.com/bddjr/simple-short-link-go
cd simple-short-link-go
./run.sh
```

短链名称不得包含 `/` `\` `.`，开头 `@` 用于后端内部重定向

安装服务：

```
su
cd /opt
git clone https://github.com/bddjr/simple-short-link-go
cd simple-short-link-go
go build -trimpath -ldflags "-w -s"
cp shortlink.service /etc/systemd/system/
chmod 744 /etc/systemd/system/shortlink.service
systemctl enable shortlink
systemctl start shortlink
```
