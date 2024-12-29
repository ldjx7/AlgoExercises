# AlgoExercises
用于记录一些练习题目

在 Golang 中，用于接收用户输入的方法主要包括以下几种方式，每种方式都有特定的适用场景。以下是常见方法及其区别的详细介绍：

1. fmt.Scan 系列函数
   包括 fmt.Scan、fmt.Scanln 和 fmt.Sscanf。
   这些函数用于从标准输入读取数据并赋值给变量。  
   使用场景  
   适合读取格式化的输入，通常用于简单交互。  
   示例代码
    ```go
    package main
    
    import "fmt"
    
    func main() {
        var name string
        var age int
    
        // Scan 以空格为分隔符读取
        fmt.Println("请输入姓名和年龄，用空格分隔：")
        fmt.Scan(&name, &age)
        fmt.Printf("姓名: %s, 年龄: %d\n", name, age)
    
        // Scanln 在遇到换行符时结束
        fmt.Println("请输入姓名（单独一行）：")
        fmt.Scanln(&name)
        fmt.Printf("姓名: %s\n", name)
    
        // Sscanf 用于从字符串中解析输入
        input := "John 30"
        fmt.Sscanf(input, "%s %d", &name, &age)
        fmt.Printf("姓名: %s, 年龄: %d\n", name, age)
    }
    ```
    特点  
    fmt.Scan：按空格分隔输入。  
    fmt.Scanln：按换行符分隔输入。  
    fmt.Sscanf：从字符串中按格式解析。  
    缺点  
    输入时必须匹配变量的类型和数量，否则会出错。  
    需要逐个定义变量，处理复杂输入会显得繁琐。  
2. bufio.Scanner  
   bufio.Scanner 是一种基于缓冲的输入方法，适合逐行读取数据。  
   使用场景  
   适用于需要逐行读取用户输入或大文本的场景。  
   示例代码  
    ```go
    package main
    
    import (
    "bufio"
    "fmt"
    "os"
    )
    
    func main() {
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Println("请输入一行文本：")
    
        if scanner.Scan() {
            input := scanner.Text()
            fmt.Println("你输入的是：", input)
        }
    
        // 错误处理
        if err := scanner.Err(); err != nil {
            fmt.Println("读取输入时出错：", err)
        }
    }
   ```
    特点  
    支持逐行读取。  
    使用 scanner.Text() 获取输入内容。  
    可以通过 scanner.Scan() 在循环中持续读取输入。  
    缺点  
    默认的缓冲区大小限制为 64 KB。如果输入过大，需要手动调整缓冲区大小。  
3. os.Stdin 和 ioutil.ReadAll  
   直接从标准输入读取所有数据。  
   使用场景  
   适合读取多行文本或需要处理标准输入中的所有数据。  
   示例代码  
    ```go
    package main
    
    import (
        "fmt"
        "io/ioutil"
        "os"
    )
    
    func main() {
        fmt.Println("请输入多行文本，按 Ctrl+D 结束：")
        input, err := ioutil.ReadAll(os.Stdin)
        if err != nil {
            fmt.Println("读取输入时出错：", err)
            return
        }
        fmt.Println("你输入的是：")
        fmt.Println(string(input))
    }
    ```
    特点
    一次性读取标准输入的全部内容。
    适合处理大块数据或连续输入。
    缺点
    可能占用大量内存，尤其是处理大输入时。
    输入结束时需要用户按下 EOF（Linux: Ctrl+D，Windows: Ctrl+Z）。
4. 命令行参数：os.Args  
   从程序启动时的命令行参数中读取数据。  
   使用场景  
   适合处理启动参数，例如传递文件路径或标志。  
   示例代码  
    ```go
    package main
    
    import (
        "fmt"
        "os"
    )
    
    func main() {
        args := os.Args
        fmt.Println("命令行参数：", args)
    
        if len(args) > 1 {
            fmt.Println("第一个参数：", args[1])
        } else {
            fmt.Println("未提供参数")
        }
    }
    ```
    特点
    os.Args[0] 是程序本身的路径，os.Args[1:] 是实际的参数。  
    不需要用户在程序运行时提供输入。  
    缺点  
    仅能在程序启动时提供参数，不能用于交互式输入。  

**区别总结**

| 方法                     | 使用场景                       | 特点                                   | 缺点                                         |
|--------------------------|--------------------------------|----------------------------------------|----------------------------------------------|
| fmt.Scan 系列            | 简单格式化输入                 | 快速，语法简单，支持直接解析到变量     | 格式严格，输入需匹配变量类型，处理复杂输入繁琐 |
| bufio.Scanner            | 按行读取或需要循环读取的场景   | 灵活，支持逐行处理输入，内存占用少     | 默认缓冲区限制 64 KB，大量数据需手动调整      |
| os.Stdin 和 ioutil.ReadAll | 一次性读取所有输入             | 简单快捷，适合处理多行或大块文本       | 内存占用较高，输入结束需按 EOF               |
| os.Args                  | 读取命令行参数                 | 非交互式，直接通过命令行传参，适合传递启动参数 | 仅能在程序启动时提供，不能动态输入           |
