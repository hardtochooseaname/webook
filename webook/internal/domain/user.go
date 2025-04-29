package domain


// user 领域对象，代表 DDD 中的实体或者聚合根
// 领域对象是业务逻辑的核心，包含了业务规则和行为
// 其他地方有人叫 BO（业务对象）等
type User struct {
	Email	string
	Password	string
}