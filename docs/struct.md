# 项目结构

```plain
.
├── default.nix          # nix 项目管理 (存放项目名称等)
├── docs                 # 文档
│   ├── Interface.md     # 接口文档
│   └── struct.txt       # 项目结构
├── flake.lock           # nix 版本管理
├── flake.nix            # nix 表达式 用于生成合适的环境shell
├── go.mod               # 依赖库以及依赖库的版本
├── gomod2nix.toml       # 在nix下类似go.mod的东西
├── go.sum               # 依赖的 module 的校验信息
├── main.go              # 主程序
├── README.md            # 项目说明文件
└── shell.nix            # nix 表达式 用于生成合适的环境shell

2 directories, 11 files
```
