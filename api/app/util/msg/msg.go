package msg

// func Tpl(msg string) string {
// 	return `{"errors": {"Password" : {"message": "` + msg + `"}}}`
// }

func Tpl(column string, msg string) string {
	return `{"errors": {"` + column + `" : {"message": "` + msg + `"}}}`
}
