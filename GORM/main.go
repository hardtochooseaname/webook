package main

import (
    "fmt"
    "log"

    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
)

type Product struct {
    gorm.Model
    Code  string
    Price uint
}

// printProducts 查询并打印 products 表的所有行
func printProducts(db *gorm.DB, note string) {
    var products []Product
    if err := db.Find(&products).Error; err != nil {
        log.Printf("查询失败 (%s): %v\n", note, err)
        return
    }
    fmt.Printf("%s: %d row(s)\n", note, len(products))
    for _, p := range products {
        fmt.Printf("  ID=%d  Code=%q  Price=%d\n", p.ID, p.Code, p.Price)
    }
    fmt.Println()
}

func main() {
    // 1. 打开数据库且关闭 GORM 自带的所有 SQL 日志
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Silent),
    })
    if err != nil {
        log.Fatalf("无法连接数据库: %v", err)
    }

    // 2. 自动迁移：创建或更新表结构
	//   i) 查系统表 sqlite_master 看对应表是否存在
	//	 ii) 如果表存在，则读出一行数据看表结构是否一致，决定是否需要修改表结构
	// 	 iii) 如果表不存在，则根据结构体字段创建新表
    if err := db.AutoMigrate(&Product{}); err != nil {
        log.Fatalf("迁移失败: %v", err)
    }
    printProducts(db, "迁移完成后")  // 这里应该是空表

    // 3. Create：插入一条新记录
    db.Create(&Product{Code: "D42", Price: 100})
    printProducts(db, "Create 之后")

    // 4. Update：单字段更新 Price → 200
    var p Product
    db.First(&p, "code = ?", "D42")
    db.Model(&p).Update("Price", 200)
    printProducts(db, "Update Price 之后")

    // 5. Update：多字段更新（用 struct）
    db.Model(&p).Updates(Product{Code: "F42", Price: 300})
    printProducts(db, "Updates struct 之后")

    // 6. Update：多字段更新（用 map）
    db.Model(&p).Updates(map[string]interface{}{"Code": "G42", "Price": 400})
    printProducts(db, "Updates map 之后")

    // 7. Delete：删除这条记录
    db.Delete(&p)
    printProducts(db, "Delete 之后")
}
