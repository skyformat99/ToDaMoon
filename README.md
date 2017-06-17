# ToDaMoon
ToDaMoon是我的自己的虚拟币交易系统。

## 使用方式
1. ``git clone https://github.com/aQuaYi/ToDaMoon.git``到你的``$GOPATH``目录。
1. 按照注释，修改`Makefile`文件的第2行的`BINARY`和第4行的`FILE_PATH`设置。
1. 在命令行进入``ToDaMoon/main``目录后，使用``make``命令生成运行程序`ToDaMoon`。
1. 在命令行进入`FILE_PATH`目录，使用`./ToDaMoon`运行程序。


## Tips
>注意：以下方式适用于Linux操作系统
1. [后台运行ToDaMoon](./Tips/Run-TDM-in-Background.md)
1. [开机启动ToDaMoon](./Tips/Run-TDM-at-Startup.md)