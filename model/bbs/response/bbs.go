package response

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/bbs"
)

type BbsResponse struct {
	//Article example.ExaCustomer `json:"customer"`
	Article  bbs.BbsArticle  `json:"article"`
	Category bbs.BbsCategory `json:"category"`
}
