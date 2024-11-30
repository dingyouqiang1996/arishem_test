# arishem_test
- arishem规则引擎调试
- 后端: `gin`
- 前端: `gin模版引擎`
- 引擎
```shell
go get github.com/bytedance/arishem@latest
go get -u ./... # 更新arishem相关的依赖
```
- `gin` 热加载工具
```shell
go install github.com/air-verse/air@latest
```
## 前端
- 前端传入数据
```json
{
    "OpLogic": "&&",
    "Conditions": [
        {
            "Operator": "==",
            "Lhs": {
                "Const": {
                    "NumConst": 1
                }
            },
            "Rhs": {
                "Const": {
                    "NumConst": 1
                }
            }
        }
    ]
}
```