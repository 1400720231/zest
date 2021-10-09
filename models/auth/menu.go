package auth

//权限树
type Tree struct {
	Id       int
	AuthName string
	UrlFor   string
	Weight   int
	Children []*Tree
}
