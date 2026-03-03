package bilireq

import (
	"math/big"
)

func TalkerUser(uid int64) Talker {
	return Talker{ID: uid, SessionType: SessionTalkerTypeUser}
}

type Talker struct {
	ID          int64             `json:"talker_id" url:"talker_id"`       // 聊天对象的id. session_type 为 1 时表示用户 mid，为 2 时表示粉丝团 id
	SessionType SessionTalkerType `json:"session_type" url:"session_type"` // 聊天对象的类型
}

// 会话对象
type MsgSession struct {
	Talker
	AtSeqNo           uint64             `json:"at_seqno"`             // 最近一次未读at自己的消息的序列号. 在粉丝团会话中有效，若没有未读的 at 自己的消息则为 0
	TopTs             int64              `json:"top_ts"`               // 置顶该会话的时间. 微秒级时间戳；若未置顶该会话则为 0；用于判断是否置顶了会话
	GroupName         string             `json:"group_name"`           // 粉丝团名称. 在粉丝团会话中有效，其他会话中为空字符串
	GroupCover        string             `json:"group_cover"`          // 粉丝团头像. 在粉丝团会话中有效，其他会话中为空字符串
	IsFollow          int64              `json:"is_follow"`            // 是否关注了对方. 在用户会话中有效，系统会话中为 1, 其他会话中为 0
	IsDnd             int64              `json:"is_dnd"`               // 是否对会话设置了免打扰
	AckSeqNo          uint64             `json:"ack_seqno"`            // 最近一次已读的消息序列号. 用于快速跳转到首条未读的消息
	AckTs             int64              `json:"ack_ts"`               // 最近一次已读时间
	SessionTs         int64              `json:"session_ts"`           // 会话时间
	UnreadCount       int64              `json:"unread_count"`         // 未读消息数
	LastMsg           *PrivateMsgItem    `json:"last_msg"`             // 最近的一条消息
	GroupType         GroupType          `json:"group_type"`           // 粉丝团类型. 在粉丝团时有效
	CanFold           int64              `json:"can_fold"`             // 会话是否可被折叠入未关注人消息. 在用户会话中有效
	Status            int64              `json:"status"`               // 会话状态. 详细信息有待补充
	MaxSeqNo          uint64             `json:"max_seqno"`            // 最近一条消息的序列号
	NewPushMsg        int64              `json:"new_push_msg"`         // 是否有新推送的消息
	Setting           int64              `json:"setting"`              // 推送设置. 0：接收推送; 1：不接收推送; 2：（？）;
	IsGuardian        int64              `json:"is_guardian"`          // 自己是否为对方的骑士（？）. 在用户会话中有效. 0：否; 2：是（？）;
	IsIntercept       int64              `json:"is_intercept"`         // 会话是否被拦截
	IsTrust           int64              `json:"is_trust"`             // 是否信任此会话. 若为 1，则表示此会话之前被拦截过，但用户选择信任本会话
	SystemMsgType     SystemMsgType      `json:"system_msg_type"`      // 系统会话类型
	AccountInfo       *SystemAccountInfo `json:"account_info"`         // 会话信息. 仅在系统会话中出现
	BizMsgUnreadCount int64              `json:"biz_msg_unread_count"` // 未读通知消息数
}

func (m *MsgSession) NoNew() bool {
	if m.LastMsg == nil {
		return false
	}
	return m.LastMsg.MsgSeqNo == m.AckSeqNo
}

type SessionTalkerType int64 // 聊天对象的类型

const (
	SessionTalkerTypeUser SessionTalkerType = 1 // 用户
	SessionTalkerTypeFans SessionTalkerType = 2 // 粉丝团
)

type GroupType int64 // 粉丝团类型

const (
	GroupType0 GroupType = 0 // 应援团
	GroupType2 GroupType = 2 // 官方群（如：ID 为 10 的粉丝团）
)

type SystemMsgType int64 // 系统会话类型

const (
	SystemMsgType0      SystemMsgType = iota // 非系统会话
	SystemMsgTypeLive                        // 主播小助手
	_                                        //
	_                                        //
	_                                        //
	SystemMsgTypeNotify                      // 系统通知（？）
	_                                        //
	SystemMsgTypeUp                          // UP主小助手
	SystemMsgTypeKefu                        // 客服消息
	SystemMsgTypePay                         // 支付小助手
)

type SystemAccountInfo struct {
	Name string `json:"name"`    // 会话名称
	Pic  string `json:"pic_url"` // 会话头像
}

// 私信主体对象
type PrivateMsgItem struct {
	Sender         int64             `json:"sender_uid"`       // 发送者mid
	ReceiverType   SessionTalkerType `json:"receiver_type"`    // 接收者类型
	Receiver       int64             `json:"receiver_id"`      // 接收者id. receiver_type 为 1 时表示用户 mid，为 2 时表示粉丝团 id
	MsgType        MsgSendType       `json:"msg_type"`         // 消息类型
	Content        string            `json:"content"`          // 消息内容. 私信内容对象经过 JSON 序列化后的文本
	MsgSeqNo       uint64            `json:"msg_seqno"`        // 消息序列号. 按照时间顺序从小到大
	Timestamp      int64             `json:"timestamp"`        // 消息发送时间. 秒级时间戳
	AtUIDs         []int64           `json:"at_uids"`          // at的成员mid. 在粉丝团时有效；此项为 null 或 [0] 均表示没有 at 成员
	MsgKey         big.Int           `json:"msg_key"`          // 消息唯一id. 部分库在解析JSON对象中的大数时存在数值的精度丢失问题，因此在处理此字段时可能会出现问题，建议使用修复了这一问题的库（如将大数转换成文本）
	MsgStatus      MsgStatus         `json:"msg_status"`       // 消息状态
	SysCancel      bool              `json:"sys_cancel"`       // 是否为系统撤回
	NotifyCode     string            `json:"notify_code"`      // 通知代码. 发送通知时使用，以下划线 _ 分割，第 1 项表示主业务 id，第 2 项表示子业务 id；若这条私信非通知则为空文本；详细信息有待补充
	NewFaceVersion int64             `json:"new_face_version"` // 表情包版本. 为 0 或无此项表示旧版表情包，此时 B 站会自动转换成新版表情包，例如 [doge] -> [tv_doge]；1 为新版
	MsgSource      MsgSource         `json:"msg_source"`       // 消息来源
}

// 消息状态
type MsgStatus int64

const (
	MsgStatusOK             MsgStatus = iota // 正常
	MsgStatusRevoke                          // 被撤回（接口仍能返回被撤回的私信内容）
	MsgStatusRevokeBySystem                  // 被系统撤回（如：消息被举报；私信将不会显示在前端，B站接口也不会返回被系统撤回的私信的信息）

	MsgStatusImgExipred MsgStatus = 50 //
)

// 消息来源列表
type MsgSource int64

const (
	MsgSourceUnkown           MsgSource = iota // 未知来源
	MsgSourceIOS                               // iOS
	MsgSourceAndroid                           // Android
	MsgSourceH5                                // H5
	MsgSourcePC                                // PC客户端
	MsgSourcePush                              // 官方推送消息. 包括：官方向大多数用户自动发送的私信（如：UP主小助手的推广）等
	MsgSourceNotification                      // 推送/通知消息. 包括：特别关注时稿件的自动推送、因成为契约者而自动发送的私信、包月充电回馈私信、官方发送的特定于自己的消息（如：UP主小助手的稿件审核状态通知）等
	MsgSourceWeb                               // Web
	MsgSourceAutoFollow                        // 自动回复 - 被关注回复. B站前端会显示“此条消息为自动回复”
	MsgSourceAutoReply                         // 自动回复 - 收到消息回复
	MsgSourceAutoKeywords                      // 自动回复 - 关键词回复
	MsgSourceAutoGuard                         // 自动回复 - 大航海上船回复
	MsgSourceAutoUpVideo                       // 自动推送 - UP 主赠言. 在以前稿件推送消息与其附带的 UP 主赠言是 2 条不同的私信（其中 UP 主赠言的消息来源代码为 12），现在 UP 主赠言已并入为稿件自动推送消息的一部分（attach_msg）
	MsgSourceSystemFans                        // 粉丝团系统提示. 如：粉丝团中的提示信息“欢迎xxx入群”
	_                                          //
	_                                          //
	MsgSourceSystem                            // 系统. 目前仅在 msg_type 为 51 时使用该代码
	MsgSourceAutoFollowMutual                  // 互相关注. 互相关注时自动发送的私信“我们已互相关注，开始聊天吧~”
	MsgSourceSystemTip                         // 系统提示. 目前仅在 msg_type 为 18 时使用该代码，如：“对方主动回复或关注你前，最多发送1条消息”
	MsgSourceAI                                // AI. 如：给搜索AI助手测试版发送私信时对方的自动回复
)
