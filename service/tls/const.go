package tls

const (
	defaultLogUserAgent = "tls-golang-sdk-v0.1.0"
	RequestIDHeader     = "x-tls-requestid"
	AgentHeader         = "User-Agent"
	ContentMd5Header    = "Content-MD5"
	ServiceName         = "TLS"

	CompressLz4  = "lz4"
	CompressNone = "none"

	FullTextIndexKey = "__content__"

	PathCreateProject    = "/CreateProject"
	PathDescribeProject  = "/DescribeProject"
	PathDeleteProject    = "/DeleteProject"
	PathModifyProject    = "/ModifyProject"
	PathDescribeProjects = "/DescribeProjects"

	PathCreateTopic    = "/CreateTopic"
	PathDescribeTopic  = "/DescribeTopic"
	PathDeleteTopic    = "/DeleteTopic"
	PathModifyTopic    = "/ModifyTopic"
	PathDescribeTopics = "/DescribeTopics"

	PathCreateIndex   = "/CreateIndex"
	PathDescribeIndex = "/DescribeIndex"
	PathDeleteIndex   = "/DeleteIndex"
	PathModifyIndex   = "/ModifyIndex"
	PathSearchLogs    = "/SearchLogs"

	PathDescribeShards = "/DescribeShards"

	PathPutLogs        = "/PutLogs"
	PathDescribeCursor = "/DescribeCursor"
	PathConsumeLogs    = "/ConsumeLogs"

	PathCreateRule               = "/CreateRule"
	PathDeleteRule               = "/DeleteRule"
	PathModifyRule               = "/ModifyRule"
	PathDescribeRule             = "/DescribeRule"
	PathDescribeRules            = "/DescribeRules"
	PathApplyRuleToHostGroups    = "/ApplyRuleToHostGroups"
	PathDeleteRuleFromHostGroups = "/DeleteRuleFromHostGroups"

	PathCreateHostGroup        = "/CreateHostGroup"
	PathDeleteHostGroup        = "/DeleteHostGroup"
	PathModifyHostGroup        = "/ModifyHostGroup"
	PathDescribeHostGroup      = "/DescribeHostGroup"
	PathDescribeHostGroups     = "/DescribeHostGroups"
	PathDescribeHostGroupRules = "/DescribeHostGroupRules"

	PathDeleteHost    = "/DeleteHost"
	PathDescribeHosts = "/DescribeHosts"

	PathCreateAlarmNotifyGroup    = "/CreateAlarmNotifyGroup"
	PathDeleteAlarmNotifyGroup    = "/DeleteAlarmNotifyGroup"
	PathDescribeAlarmNotifyGroups = "/DescribeAlarmNotifyGroups"
	PathModifyAlarmNotifyGroup    = "/ModifyAlarmNotifyGroup"
	PathCreateAlarm               = "/CreateAlarm"
	PathDeleteAlarm               = "/DeleteAlarm"
	PathModifyAlarm               = "/ModifyAlarm"
	PathDescribeAlarms            = "/DescribeAlarms"
)
