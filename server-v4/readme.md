# 目录结构
```
-global
-instruction
-log
-model
-persistence
  -aof
-viper
main.go
```

1. global: 保存所有的全局变量。包括服务器配置、存储数据的哈希表
2. instruction： 实现对命令的执行函数
3. log：日志记录
4. model：类
5. persistence：持久化保存
   1. aof：aof持久化方法
6. viper：读取配置文件，根据配置选择策略
7. main：启动函数

# 日志服务
使用zap包
```go
package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var (
	SystemLog *zap.Logger // SystemLog is the logger for system logs.
)

func init() {
	// Create a new production encoder configuration.
	config := zap.NewProductionEncoderConfig()
	// Set the time encoder to ISO8601.
	config.EncodeTime = zapcore.ISO8601TimeEncoder

	// Open the info log file with necessary permissions.
	infoFile, _ := os.OpenFile("./log/file"+"/info-log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	// Open the error log file with necessary permissions.
	errorFile, _ := os.OpenFile("./log/file"+"/error-log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	// Create a new core for the logger that writes to both the info and error log files.
	treeCore := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewJSONEncoder(config), zapcore.AddSync(infoFile), zapcore.InfoLevel),
		zapcore.NewCore(zapcore.NewJSONEncoder(config), zapcore.AddSync(errorFile), zapcore.ErrorLevel),
	)

	// Initialize the logger with the created core and add the caller option.
	SystemLog = zap.New(treeCore, zap.AddCaller())
}

```

# viper
使用viper包，配置文件采用yaml
```go
package viper

import (
    "github.com/spf13/viper"
    "go.uber.org/zap"
    "server-v4/log"
)

func init() {
    viper.SetConfigName("config")
    viper.AddConfigPath("./viper")
    viper.SetConfigType("yaml")

    err := viper.ReadInConfig()
    if err != nil {
        log.SystemLog.Fatal("Failed to read the configuration file.",
                           zap.Error(err))
    }
}
```


# 数据结构

1. sds

使用golang自带的string

2. set
```go
type setInterface interface {
    Init()
    Add(str ...string)
    Delete(str ...string)
    Has(str string) bool
    ToString() string
    Intersection(set *Set) *Set
    Union(set *Set) *Set
    Difference(set *Set) *Set
}

type Set struct {
    data map[string]bool
}
```

3. instruction
```go
type Command struct {
	Value          string // instruction name
	BelongTo       string // belong to which data structure：such as sds, list, hash, set, zset
	WillChangeData bool   // whether it will change data
}

type Instruction struct {
	Cmd   Command  // command information
	Key   string   // the key of the command
	Value []string // the value of the command
	// the function to execute the command, return the result to print and the error
	Execute func(*Instruction) ([]string, error)
}
```

# 全局变量
```go
var (
	// WholeMap store all the map
	WholeMap MapSet
	// Config store all the system config
	Config RunTimeConfig
)
```

1. 系统配置
```go
type RunTimeConfig struct {
	Os string
}
```
获取操作系统信息。
ps：设计的目的是由于每个操作系统的换行符不一样，会影响读取命令行后的字符串切割等操作。
strings.Field函数能很好的处理换行符问题。

2. hash存储

该版本暂时将每个不同类型用不同的哈希表存储
```go
type MapSet struct {
	SdsMap    map[string]model.Sds
	ListMap   map[string][]list.List
	HashMap   map[string]map[string]model.Sds
	SetMap    map[string]model.Set
	ExpireMap map[string]time.Time
}
```

# AOF存储
# aof写入
> 包变量

```go
const (
	// aofFileNane is the path of the aof file
	aofFileNane = "./persistence/aof.aof"
)

var (
	tactic         string   // Tactic is the way to enter the aof mode
	fileConnector  *os.File // connection to the aof file
	instructionBuf []string // buffer for instructions
)

var (
	// WriteToAof is the function to write the instruction to the aof file
	WriteToAof func(str string)
	// GoRoutineMethod is the method to run the aof go routine
	GoRoutineMethod func()
)
```

1. tactic：aof缓存写入磁盘的策略
2. fileConnector：与文件的连接
3. instructionBuf：aof缓冲区
4. WriteToAof func(str string)：根据策略选择向外暴露的写入AOF函数
5. GoRoutineMethod func()：根据策略选择向外暴露的后台方法
   1. always：该方法为空
   2. everysecond：每一秒都将缓冲区文件写入磁盘

> 问题

该版本中，fileConnector的生命周期与redis运行的生命周期相同，因此一直没有close。但这样会导致aof重写时，替换旧aof会出问题


## aof重写（未完善）
> 流程

1. 开启go协程
2. 开启一个缓冲区，保存在重写过程中新增的写操作
3. 读取全局hash表，编写新的aof文件
4. 删除旧aof
5. 将新aof改名
6. 将缓冲区写入新aof


# 命令执行
命令的结构
```go
type Command struct {
	Value          string // instruction name
	BelongTo       string // belong to which data structure：such as sds, list, hash, set, zset
	WillChangeData bool   // whether it will change data
}

type Instruction struct {
	Cmd   Command  // command information
	Key   string   // the key of the command
	Value []string // the value of the command
	// the function to execute the command, return the result to print and the error
	Execute func(*Instruction) ([]string, error)
}
```
该版本的策略是将所有的命令都预先创建，然后存储在一个hash表中

执行方法：func ExecuteInstruction(str string, flag int)

- str是命令的字符串
- flag是标识符，判断该命令是由系统aof发起还是用户发起，0-system；1-user

> 执行流程

1. 处理命令字符串，并将其按照空格进行分割
```go
splits := strings.Fields(str)
```

2. 根据命令行的命令，从预置的命令hash表中获取命令结构
3. 根据获取的命令结构提供的执行函数进行执行命令
4. 根据命令的发起者，决定是否打印输出/aof写入等
