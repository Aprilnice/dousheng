package snowflaker

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"sync"
	"time"
)

/*
雪花算法生成ID
*/

var (
	node *snowflake.Node
	once sync.Once
)

func Init(startDate string, machineID int64) (err error) {

	once.Do(func() {
		st, err := time.Parse("2006-01-01", startDate)
		if err != nil {
			fmt.Println("时间解析错误")
		}

		// 时间片偏移到指定的时间上
		snowflake.Epoch = st.UnixNano() / 1000000
		node, err = snowflake.NewNode(machineID)

	})
	return err
}

func NextID() int64 {
	return node.Generate().Int64()
}
