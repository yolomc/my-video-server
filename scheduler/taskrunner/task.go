package taskrunner

import (
	"sync"
)

func VideoClearDispatcher(dc dataChan) error {
	//获取要删除的videoid list

	return nil
}

func VideoClearExecutor(dc dataChan) error {

	errMap := &sync.Map{}
	var err error

forloop:
	for {
		select {
		case vid := <-dc:
			go func(id interface{}) {
				//删除文件 os.Remove
				//删除数据
				//保存报错信息
				//errMap.Store(id, err)
			}(vid)
		default:
			break forloop
		}
	}

	errMap.Range(func(k, v interface{}) bool {
		err = v.(error)
		if err != nil {
			return false
		}
		return true
	})
	return err
}
