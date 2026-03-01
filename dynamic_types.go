package bilireq

type Card struct {
	Card string    `json:"card,omitempty"`
	Desc *CardDesc `json:"desc,omitempty"`
}

type CardDesc struct {
	DynamicId    int64                `json:"dynamic_id,omitempty"`
	Type         DynamicDescType      `json:"type,omitempty"`
	Timestamp    int64                `json:"timestamp,omitempty"`
	DynamicIdStr string               `json:"dynamic_id_str,omitempty"`
	UserProfile  *CardDescUserProfile `json:"user_profile,omitempty"`
	BVID         *string              `json:"bvid,omitempty"`
}

type CardDescUserProfile struct {
	Info *CardDescUserProfileInfo `json:"info,omitempty"`
}

type CardDescUserProfileInfo struct {
	Uid   int64  `json:"uid,omitempty"`
	Uname string `json:"uname,omitempty"`
	Face  string `json:"face,omitempty"`
}

type DynamicDescType int32

const (
	DynamicDescType_DynamicDescTypeUnknown DynamicDescType = 0
	DynamicDescType_WithOrigin             DynamicDescType = 1
	DynamicDescType_WithImage              DynamicDescType = 2
	DynamicDescType_TextOnly               DynamicDescType = 4
	DynamicDescType_WithVideo              DynamicDescType = 8
	DynamicDescType_WithPost               DynamicDescType = 64
	DynamicDescType_WithMusic              DynamicDescType = 256
	DynamicDescType_WithAnime              DynamicDescType = 512
	// 该内容已经不见了哦
	DynamicDescType_WithMiss DynamicDescType = 1024
	// 评分、头像挂关注件，这种动态下面有一个小卡片的
	DynamicDescType_WithSketch DynamicDescType = 2048
	DynamicDescType_WithMovie  DynamicDescType = 4098
	// 电视剧、综艺
	DynamicDescType_WithDrama DynamicDescType = 4099
	// 4100去哪了捏
	DynamicDescType_WithDocumentary DynamicDescType = 4101
	DynamicDescType_WithLive        DynamicDescType = 4200
	// XXX的收藏夹，收藏夹居然也可以发动态？
	DynamicDescType_WithMylist DynamicDescType = 4300
	// (付费?)课程
	DynamicDescType_WithCourse DynamicDescType = 4302
	DynamicDescType_WithLiveV2 DynamicDescType = 4308
)
