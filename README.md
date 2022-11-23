# dockerfiles

This repo is inspired by [marcel-dempers](https://github.com/marcel-dempers/my-desktop/blob/master/README.md)


## windows
[Download VcXsrv Windows X Server from SourceForge.net](https://sourceforge.net/projects/vcxsrv/files/latest/download)


## windows wsl2 -> cd ~ -> cat .bash_login

```bash
# .bash_login
export DISPLAY=$(cat /etc/resolv.conf | grep nameserver | awk '{print $2}'):0
```