# Juno(朱诺)
朱诺号木星探测器是目前人类是制造出最快的宇宙飞行器。
这里，朱诺是一个通用的易用的高性能的内存型广告检索引擎

## 目标

1. 通用性： 能试用广告检索的大部分情况
2. 易用性： 可以极低的代价从0搭建搜索引擎
3. 高性能： 本身搜索性能20ms内，单机QPS>1-2K
4. 插件化，可扩展： 检索各模块都是接口的形式，可以根据需求轻松定制

## 主要特性

1. 支持倒排索引
   1. 数值型（int, double）
   2. 字符串型
2. 正排索引
   1. 数值型（int, double）
   2. 字符串型
   3. set集合
   4. List
   5. KV
3. 查询支持多索引查询、布尔查询、范围查询、集合查询

## 示例

```go
// 建立索引
index := Index.NewIndex("indexName")
index.Add(docInfo)

// 查询
query := NewQuery(NewAndExpress(
	NewEqExpression("country", "us"),
	NewRangeExpression("price", 1, 20)，
  NewOrExpress(
		NewEqExpression("country", "us"),
    NewInExpression("packageName", "package1", "package2")
	),
)	
searchResult := index.Search(query)
fmt.Println(searchResult)
```



## 设计

搜索引擎主要分为2个部分

1. 查询
2. 索引

### 一、查询

#### 查询语法

查询是类sql语法，有表达式组成（可嵌套），表达式有 and, or, not等操作

支持 =, >=, >, <=,<, !=, range, in

查询语法支持三种格式  string,  json, go stuct

```shell
country=us and (price range [1, 10]) and (platform=ios or package in ["pacakge1", "pacakge1"] )
```

```json
{
    "and": [
        {
            "=": {
                "field": "country",
                "value": "US"
            }
        },
        {
            "range": {
                "field": "price",
                "value": [
                    1,
                    20
                ]
            }
        },
        {
            "or": [
                {
                    "=": {
                        "field": "platform",
                        "value": "ios"
                    }
                },
                {
                    "in": {
                        "field": "packageName",
                        "value": [
                            "package1",
                            "package2"
                        ]
                    }
                }
            ]
        }
    ]
}
```

```go
// 构建查询
q := NewQuery(NewAndExpress(
  NewEqExpression("country", "us"),
  NewRangeExpression("price", 1, 20)，
  NewOrExpress(
    NewEqExpression("country", "us"),
    NewInExpression("packageName", "package1", "package2")
  ),          
))

// 遍历结果
for q.HasNext() {
  docid := q.Next()
}
```

#### 查询执行过程

1. 构建查询语法树

   ![](pic/search_tree.png)

2. 执行语法树

   1. 语法树本身可以抽象成一个迭代器，迭代的过程就是对倒排链查找的过程

3. 过滤

### 二、索引

#### 索引接口

index接口：

```go
type Index interface {
	Add(doc *DocInfo)  // 新增文档
	Del(doc *DocInfo)  // 删除文档
	UpDate(doc *DocInfo)  // 更新文档

	Dump(filename string)  // 将索引Dump到磁盘
	Load()  // 从磁盘加载索引

	Search(query *query.Query)  // 查询接口
}
```

DocInfo:  json结构

```json
{
    "Id": 12345,
    "Index": [
        {
            "field": "country",
            "value": "cn"
        },
        {
            "field": "platform",
            "value": "android"
        }
    ],
    "Storage": [
        {
            "field": "package",
            "type": "string",
            "value": "com.juno"
        },
        {
            "field": "price",
            "type": "float64",
            "value": 2.3
        }
    ]
}
```

#### 索引内存结构

##### 倒排

倒排索引是一个Key, InvertList结构

- Key： FieldName + Value（一个字符串、一个数值、也可以是一个数值范围）
- InvertList： 是一个有序的集合的接口，，可以是数组、跳表、排序树等
- Value:  一个字符串、一个数值、也可以是一个数值范围

###### 倒排接口：

倒排索引可以有不同的实现方式，只要满足下面的接口，都可以称之为倒排索引

```go
// InvertList 倒排结构的接口，仅负责查询，不负责索引更新
type InvertList interface {
	HasNext()
	Next() query.DocId
	GetGE(id query.DocId) query.DocId
}

type InvertIndex interface {
	GetInvertList(fieldName string) InvertList
}
```

##### 正排索引

正排分字段存储，结构为map<fieldname, <docid, value>>

###### 正排接口

```go
// 按字段存储正排信息
type StorageIndex interface {
   Get(filedName string, id query.DocId) interface{}
}
```



#### 索引构建

索引构建模块能方便的将数据源中的数据构建成索引，同时能感知数据源的变化，并将变化同步至索引中

示例：

```go
builder := IndexBuilder.NewBuilder(mongoBuilderCfg)  // 构造mongo的builder
index := builder.build()  // 开始构建索引
```

索引构建模块会支持多种数据源，如文件、mongo、mysql等

