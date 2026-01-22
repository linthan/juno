# build

build为斗鱼Go应用通用编译组件，开发者可以通过引入此repo作为自己应用的子树来对应用进行编译。

应用的文件布局应当满足如下所示结构：
```shell
.
├── bin
├── main.go
├── vendor
│   └── git.dz11.com
│       └── vega
│           └── minerva
```

由于需要编译时打入版本信息、编译时间、应用名称、应用id等，应用vendor目录下需要有git.dz11.com/vega/minerva(minerva>=v1.5.0)子包。

## 快速使用
**以下指令会删除旧build子树，并新增新build子树，若$project中build子树版本较陈旧建议采用以下指令。**
```
cd $project                                                    # 进入到应用工作目录
git remote add -f build git@git.dz11.com:vega/build.git        # 强制添加vega/build远程仓库
rm -r ./build && git add build && git commit -m "rm old build" # 删除旧build子树
git subtree add --prefix=build build master --squash           # 添加新build子树
```

## 指令说明
**若$project中build子树版本并不陈旧，可采用以下指令来进行更新**
```
git subtree pull --prefix=build build master --squash # 更新build子树
```
