# 框架设计
- main.go       函数入口，一般只会调用NewGame函数
- game.go       Game接口实现的地方。 实现了Layout Draw Update三个函数
- update.go     用于表示 后台逻辑更新的代码
- draw.go       用于表示绘制页面的代码