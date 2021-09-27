gorm查询主要由以下几个部分的函数组成，这些函数可以串起来组合sql语句，使用起来类似编写sql语句的习惯。

1.query 执行查询的函数，gorm提供下面几个查询函数：
- Take: 查询一条记录。 例如：  
```go
//定义接收查询结果的结构体变量
food := Food{}

//等价于：SELECT * FROM `foods`   LIMIT 1  
gorm.Take(&food)
```

- First: 查询一条记录，根据主键ID排序(正序)，返回第一条记录。 例如：  
```go
//等价于：SELECT * FROM `foods`   ORDER BY `foods`.`id` ASC LIMIT 1    
gorm.First(&food)
```

- Last: 查询一条记录, 根据主键ID排序(倒序)，返回第一条记录。 例如：  
```go
//等价于：SELECT * FROM `foods`   ORDER BY `foods`.`id` DESC LIMIT 1   
//语义上相当于返回最后一条记录
gorm.Last(&food)
```

- Find 查询多条记录，Find函数返回的是一个数组
例子：
```go
//因为Find返回的是数组，所以定义一个商品数组用来接收结果
var foods []Food
//等价于：SELECT * FROM `foods`
gorm.Find(&foods)
```

- Pluck: 查询一列值。 例如：  
```go
//商品标题数组
var titles []string

//返回所有商品标题
//等价于：SELECT title FROM `foods`
//Pluck提取了title字段，保存到titles变量
//这里Model函数是为了绑定一个模型实例，可以从里面提取表名。
gorm.Model(&Food{}).Pluck("title", &titles)
```

2.where
上面的例子都没有指定where条件，这里介绍下如何设置where条件，主要通过gorm.Where函数设置条件.
函数说明：
gorm.Where(query interface{}, args ...interface{})

参数说明:

参数名	说明
query	sql语句的where子句, where子句中使用问号(?)代替参数值，则表示通过args参数绑定参数
args	where子句绑定的参数，可以绑定多个参数
例子1:
//等价于: SELECT * FROM `foods`  WHERE (id = '10') LIMIT 1
//这里问号(?), 在执行的时候会被10替代
gorm.Where("id = ?", 10).Take(&food)

//例子2:
// in 语句 
//等价于: SELECT * FROM `foods`  WHERE (id in ('1','2','5','6')) LIMIT 1 
//args参数传递的是数组
gorm.Where("id in (?)", []int{1,2,5,6}).Take(&food)

//例子3:
//等价于: SELECT * FROM `foods`  WHERE (create_time >= '2018-11-06 00:00:00' and create_time <= '2018-11-06 23:59:59')
//这里使用了两个问号(?)占位符，后面传递了两个参数替换两个问号。
gorm.Where("create_time >= ? and create_time <= ?", "2018-11-06 00:00:00", "2018-11-06 23:59:59").Find(&foods)

//例子4:
//like语句
//等价于: SELECT * FROM `foods`  WHERE (title like '%可乐%')
gorm.Where("title like ?", "%可乐%").Find(&foods)
3.select
设置select子句, 指定返回的字段

//例子1:
//等价于: SELECT id,title FROM `foods`  WHERE `foods`.`id` = '1' AND ((id = '1')) LIMIT 1  
gorm.Select("id,title").Where("id = ?", 1).Take(&food)

//这种写法是直接往Select函数传递数组，数组元素代表需要选择的字段名
gorm.Select([]string{"id", "title"}).Where("id = ?", 1).Take(&food)


//例子2:
//可以直接书写聚合语句
//等价于: SELECT count(*) as total FROM `foods`
total := []int{}

//Model函数，用于指定绑定的模型，这里生成了一个Food{}变量。目的是从模型变量里面提取表名，Pluck函数我们没有直接传递绑定表名的结构体变量，gorm库不知道表名是什么，所以这里需要指定表名
//Pluck函数，主要用于查询一列值
gorm.Model(&Food{}).Select("count(*) as total").Pluck("total", &total)

fmt.Println(total[0])
4.order
设置排序语句，order by子句

//例子:
//等价于: SELECT * FROM `foods`  WHERE (create_time >= '2018-11-06 00:00:00') ORDER BY create_time desc
gorm.Where("create_time >= ?", "2018-11-06 00:00:00").Order("create_time desc").Find(&foods)
5.limit & Offset
设置limit和Offset子句，分页的时候常用语句。

//等价于: SELECT * FROM `foods` ORDER BY create_time desc LIMIT 10 OFFSET 0 
gorm.Order("create_time desc").Limit(10).Offset(0).Find(&foods)
6.count
Count函数，直接返回查询匹配的行数。

//例子:
var total int64 = 0
//等价于: SELECT count(*) FROM `foods` 
//这里也需要通过model设置模型，让gorm可以提取模型对应的表名
gorm.Model(Food{}).Count(&total)
fmt.Println(total)
7.分组
设置group by子句

//例子:
//统计每个商品分类下面有多少个商品
//定一个Result结构体类型，用来保存查询结果
type Result struct {
    Type  int
    Total int
}

var results []Result
//等价于: SELECT type, count(*) as  total FROM `foods` GROUP BY type HAVING (total > 0)
gorm.Model(Food{}).Select("type, count(*) as  total").Group("type").Having("total > 0").Scan(&results)

//scan类似Find都是用于执行查询语句，然后把查询结果赋值给结构体变量，区别在于scan不会从传递进来的结构体变量提取表名.
//这里因为我们重新定义了一个结构体用于保存结果，但是这个结构体并没有绑定foods表，所以这里只能使用scan查询函数。
提示：Group函数必须搭配Select函数一起使用