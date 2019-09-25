package rpc

// ByUploadTime  根据上传日期排序(最新上传排在前)
type ByUploadTime []UserFile

func (a ByUploadTime) Len() int           { return len(a) }
func (a ByUploadTime) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByUploadTime) Less(i, j int) bool { return a[j].UploadAt < a[i].UploadAt }
