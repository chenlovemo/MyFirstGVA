package response

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/file"
)

type FileManageResponse struct {
	FileManage file.FileManageStruct `json:"fileManage"`
}
