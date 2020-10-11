# GoAdmin 官方主题

- [adminlte](https://github.com/HongJaison/ninja-agent-themes3/tree/master/adminlte)
- [sword](https://github.com/HongJaison/ninja-agent-themes3/tree/master/sword)

## 如何使用

- 导入主题
- 在全局配置中设置

```go

package main

import (
	...
	_ "github.com/HongJaison/ninja-agent-themes3/adminlte"
	...
)

func main()  {
	
	...
	
	cfg := config.Config{
    		...
    		
    		Theme: "adminlte",
    		
    		...
    	}
	
	...
 
}

```

## 如何修改，自定义

使用每个主题下面的 Makefile 命令