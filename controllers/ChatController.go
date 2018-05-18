package controllers
//todo:加头像| 改处理方式为 websocket 加队列 并发不安全
import (
	"github.com/astaxie/beego"
	"strings"
	"time"
	"encoding/json"
	"bufio"
	"os"
)

var  chatersM  map[string]*Chater  //聊者管理

var chatRecordsM map[int64]*ChatRecord //用于所有的聊天管理 扩展持久化

var replyM map[int64]ChatRecord  //用于刷新时管理新记录

var replyEntity Reply  //用户刷新回复

var chaterNames []string  //用作刷新在线

var  bufW  *bufio.Writer


func init(){
	go initTimer()
	chatersM =make( map[string]*Chater,0)
	chatRecordsM=make(map[int64]*ChatRecord,0)  //step 为1
	replyM=make(map[int64]ChatRecord,0)  //时间戳  记录


	file,_:=os.OpenFile("log.txt",os.O_RDWR|os.O_CREATE|os.O_APPEND,0777)
	bufW =bufio.NewWriter(file)
}
const (
	OPERATE_OK=0          //添加聊天者成功
	CHATER_PARSE_ERR =1  //解析不出聊者
	CHATER_EXIST_ALREADY=2  //同名聊者已在线
	CHATER_NOT_EXIST=3     //同名聊者不存在（下线删除）
	RECORD_PARSE_ERR=4     //同名聊者不存在（下线删除）
	RESET_MINUTE_MIN=20
)

type  Reply struct {
	Joiner []string
	Records []ChatRecord
}

type IntervalReply struct {
		Online []string
}

type ChatRecord struct{
	Content string 	`form:"content"`
	Who  string   `form:"who"`
	Unix int64
}

type ChatController struct {
	beego.Controller
}

type Chater struct{
		Name string `form:"chatName"`
	}


// 加入聊天室验证
func (this *ChatController) Verify(){

	chater:=Chater{}

	if err:=this.ParseForm(&chater);err!=nil{
		this.Data["json"]=CHATER_PARSE_ERR
		this.ServeJSON()
	}

	this.Data["json"]=AddChater(&chater)
	this.ServeJSON()

}

//进入主页

func (this *ChatController) Index(){

	if name:=this.GetString("name");strings.TrimSpace(name)!="" {
		this.Data["who"]=strings.TrimSpace(name)
	}else{
		this.TplName="join.tpl"
	}

}

//填名字

func (this *ChatController) Join(){
}

//聊天处理
func (this *ChatController)  Chat(){

	chatRecord:=ChatRecord{}
	if err:=this.ParseForm(&chatRecord);err!=nil{
		this.Data["json"]=RECORD_PARSE_ERR
		this.ServeJSON()
		return
	}
	chatRecord.Unix=time.Now().Unix()

	chatRecordsM[chatRecord.Unix]=&chatRecord

	bts,_:=json.Marshal(chatRecord)
	this.Data["json"]= string(bts)
	this.ServeJSON()

}

//新人加入聊天室
func AddChater(chater *Chater) int {

	if _,exist:=chatersM[chater.Name];exist {
		return  CHATER_EXIST_ALREADY
	}

	chatersM[chater.Name]=chater
	chaterNames=append(chaterNames, chater.Name)
	bufW.WriteString(chater.Name+":online\n")
	bufW.Flush()
	return OPERATE_OK
}


//长轮训

func (this  *ChatController) Refresh(){
	reply:=[]ChatRecord{}

	//聊天记录中存在的比当前时间早的但不在回复管理中的
	for u,v:=range  chatRecordsM{
		if _,exist:=replyM[u];exist {
			continue
		}else{
			replyM[u]=*v
			reply=append(reply, *v)
		}
	}

	replyEntity.Records=append(replyEntity.Records, reply...)

	refreshJoiner()

	this.Data["json"]=replyEntity

	this.ServeJSON()

}

//离线

func (this *ChatController) Offline(){
	name:=this.GetString("who","nonameDeFaUlTName")



	if name== "nonameDeFaUlTName" {
		this.Data["json"]=CHATER_PARSE_ERR
	}


	this.Data["json"]=RemoveChatperByName(strings.TrimSpace(name))

	this.ServeJSON()

}

//刷新在线

func refreshJoiner(){

	replyEntity.Joiner=chaterNames

}

//helper
func updateChaterNames() []string {
	updated :=[]string{}

	for k,_:=range chatersM  {
		updated=append(updated, k)
	}

	return updated
}

//关闭页面清除内存信息

func RemoveChatperByName(name string) int{

	if _,exist:=chatersM[name];exist{
		bufW.WriteString(name+":offline\n")
		bufW.Flush()
		delete(chatersM,name)
		chaterNames=updateChaterNames()
		return OPERATE_OK
	}
	return CHATER_NOT_EXIST

}



func initTimer(){
	ticker:=time.NewTicker(time.Minute*RESET_MINUTE_MIN)
	for  {
		select {
		case <-ticker.C:
			chatRecordsM=nil
			chatRecordsM=make(map[int64]*ChatRecord,0)
		}
	}
}