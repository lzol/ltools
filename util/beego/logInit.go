package log

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type logConfigJson struct {
	FileName string `json:"filename"`
	Level    int    `json:"level"`
	MaxLines int    `json:"maxlines"`
	MaxDays  int    `json:"maxdays"`
}

func init() {
	logConfig := logConfigJson{}
	logConfig.FileName = beego.AppConfig.String("logFile")
	logConfig.Level, _ = beego.AppConfig.Int("logLevel")
	logConfig.MaxLines, _ = beego.AppConfig.Int("logMaxLines")
	logConfig.MaxDays, _ = beego.AppConfig.Int("logMaxDays")
	b, _ := json.Marshal(logConfig)
	s := string(b)
	beego.SetLogger(logs.AdapterMultiFile, s)



	//logOrmFile := beego.AppConfig.String("logOrmFile")
	//f, _ := os.OpenFile(logOrmFile, os.O_WRONLY|os.O_CREATE|os.O_SYNC,0755)
	//os.Stdout = f
	//os.Stderr = f
	//w := bufio.NewWriter(f)
	//defer w.Flush()
	//orm.DebugLog = orm.NewLog(w)
	dbLog, _ := beego.AppConfig.Bool("logOrmDebug")
	beego.Info("dbLog:", dbLog)
	orm.Debug = dbLog


	beego.Info("日志设置完成")
}
