package mongo

import "time"

// ClientName mongodb client名称枚举
type ClientName string

// ClientName 营销平台数据库集群连接名称
const (
	MKBiz   ClientName = "markting"
	MKWat   ClientName = "mk_wat"
	WP      ClientName = "wp"
	QwCrmDB            = "qw_crm" // 企业微信crm库
	Mall    ClientName = "mall"   // 商城
	QWChat ClientName = "qw_chat" // 企微 会话存档
)

// 默认超时时间配置
const (
	DefaultConnectTimeout time.Duration = 5 // mongo connect 建立mongodb连接默认超时时间 5s
	DefaultQueryTimeout   time.Duration = 5 // mongo query 查询默认超时时间 5s
)
