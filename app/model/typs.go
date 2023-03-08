package model

type SearchBase struct {
	Id int64
	// 游标
	Cursor    int64
	CursorAsc bool
	// 每页条数
	PageSize int
	// 当前页
	PageCurrent int
	// 排序 eg： ["create_time asc", "id desc"]
	OrderSort []string
	// 开始时间
	StartTime int64
	// 结束时间
	EndTime int64
	// plusCond  Benefits: It is not necessary to define each field separately, and the query conditions are flexible
	// [[field,  symbol, value, dataType]]，symbol:= != > >= < <= in like..., dataType:dataType of the value
	//[["name", "=", "服务器程序", "string"], ["complete_time", ">=", 1674373544, "date"],["id", "in", "1,2,3", "numberArray"]]
	SearchPlus [][]string
}

type MArtifactQuery struct {
	SearchBase
	*MArtifact
}

type MArtifactTypeQuery struct {
	SearchBase
	*MArtifactType
}

type MBugQuery struct {
	SearchBase
	*MBug
}

type MEnvQuery struct {
	SearchBase
	*MEnv
}

type MExternalDiffRecordsQuery struct {
	SearchBase
	*MExternalDiffRecords
}

type MExternalServerQuery struct {
	SearchBase
	*MExternalServer
}

type MExternalTaskQuery struct {
	SearchBase
	*MExternalTask
}

type MMasterVersionQuery struct {
	SearchBase
	*MMasterVersion
}

type MProjectQuery struct {
	SearchBase
	*MProject
}

type MTypeMapQuery struct {
	SearchBase
	*MTypeMap
}

type MVersionQuery struct {
	SearchBase
	*MVersion
}
