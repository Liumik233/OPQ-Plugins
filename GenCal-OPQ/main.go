package main

import (
	"github.com/mcoo/OPQBot"
	"log"
	"strconv"
	"time"
)

var sflag int
var qq int64
var atk, cr, cratk, sadd, skill, defend int64

func cal() (int64, int64, int64) {
	un:=0.52
	atk1:=float64(atk)
	cr1:=float64(cr)/100
	cratk1:=float64(cratk)/100
	sadd1:=float64(sadd)/100
	skill1:=float64(skill)/100
	defend1:=float64(defend)/100
	relh:=atk1 * skill1 * (1 + sadd1) * un
	crh:=relh * (1 + cratk1)
	relall:=relh + (crh - relh) * cr1
	return int64(relh*defend1)*10,int64(crh*defend1)*10,int64(relall*defend1)*10
}
func main() {
	opqBot := OPQBot.NewBotManager(2037403857,"http://127.0.0.1:8888")
	err := opqBot.Start()
	if err != nil {
		log.Println(err.Error())
	}
	defer opqBot.Stop()
	err = opqBot.AddEvent(OPQBot.EventNameOnFriendMessage, func(botQQ int64, packet OPQBot.FriendMsgPack) {
		log.Println(botQQ,packet)
		if packet.Content=="calstart"{
			qq= packet.FromUin
			opqBot.Send(OPQBot.SendMsgPack{
				ToUserUid: packet.FromUin,
				Content: OPQBot.SendTypeTextMsgContent{Content: "欢迎使用原神计算器\n作者：Liumik\n版本：V0.01a\n计算结果存在误差，仅供参考"},
				SendToType: OPQBot.SendToTypeFriend,
				SendType: OPQBot.SendTypeTextMsg,
			})
			opqBot.Send(OPQBot.SendMsgPack{
				ToUserUid: packet.FromUin,
				Content: OPQBot.SendTypeTextMsgContent{Content: "请输入攻击力"},
				SendToType: OPQBot.SendToTypeFriend,
				SendType: OPQBot.SendTypeTextMsg,
			})
		sflag=1
	}else if packet.FromUin==qq{
			if sflag==1{
				atk,_=strconv.ParseInt(packet.Content,10,64)
				opqBot.Send(OPQBot.SendMsgPack{
					ToUserUid: packet.FromUin,
					Content: OPQBot.SendTypeTextMsgContent{Content: "请输入暴击率"},
					SendToType: OPQBot.SendToTypeFriend,
					SendType: OPQBot.SendTypeTextMsg,
				})
				sflag=2
			}else if sflag==2{
				cr,_=strconv.ParseInt(packet.Content,10,64)
				opqBot.Send(OPQBot.SendMsgPack{
					ToUserUid: packet.FromUin,
					Content: OPQBot.SendTypeTextMsgContent{Content: "请输入暴击伤害"},
					SendToType: OPQBot.SendToTypeFriend,
					SendType: OPQBot.SendTypeTextMsg,
				})
				sflag=3
			}else if sflag==3{
				cratk,_=strconv.ParseInt(packet.Content,10,64)
				opqBot.Send(OPQBot.SendMsgPack{
					ToUserUid: packet.FromUin,
					Content: OPQBot.SendTypeTextMsgContent{Content: "请输入技能倍率"},
					SendToType: OPQBot.SendToTypeFriend,
					SendType: OPQBot.SendTypeTextMsg,
				})
				sflag=4
			}else if sflag==4{
				skill,_=strconv.ParseInt(packet.Content,10,64)
				opqBot.Send(OPQBot.SendMsgPack{
					ToUserUid: packet.FromUin,
					Content: OPQBot.SendTypeTextMsgContent{Content: "请输入元素/物理伤害加成"},
					SendToType: OPQBot.SendToTypeFriend,
					SendType: OPQBot.SendTypeTextMsg,
				})
				sflag=5
			}else if sflag==5{
				sadd,_=strconv.ParseInt(packet.Content,10,64)
				opqBot.Send(OPQBot.SendMsgPack{
					ToUserUid: packet.FromUin,
					Content: OPQBot.SendTypeTextMsgContent{Content: "请输入抗性"},
					SendToType: OPQBot.SendToTypeFriend,
					SendType: OPQBot.SendTypeTextMsg,
				})
				sflag=6
			}else if sflag==6{
				defend,_=strconv.ParseInt(packet.Content,10,64)
				r1, r2, r3 :=cal()
				opqBot.Send(OPQBot.SendMsgPack{
					ToUserUid: packet.FromUin,
					Content: OPQBot.SendTypeTextMsgContent{Content: "计算结果:"+"\n非暴击伤害期望："+strconv.FormatInt(r1,10)+"\n暴击伤害期望："+strconv.FormatInt(r2,10)+"\n综合伤害期望："+strconv.FormatInt(r3,10)},
					SendToType: OPQBot.SendToTypeFriend,
					SendType: OPQBot.SendTypeTextMsg,
				})
				sflag=0
			}
		}
	})

	err = opqBot.AddEvent(OPQBot.EventNameOnGroupShut, func(botQQ int64, packet OPQBot.GroupShutPack) {
		log.Println(botQQ, packet)
	})
	if err != nil {
		log.Println(err.Error())
	}
	err = opqBot.AddEvent(OPQBot.EventNameOnConnected, func() {
		log.Println("连接成功！！！")
	})
	if err != nil {
		log.Println(err.Error())
	}
	err = opqBot.AddEvent(OPQBot.EventNameOnDisconnected, func() {
		log.Println("连接断开！！")
	})
	if err != nil {
		log.Println(err.Error())
	}
	err = opqBot.AddEvent(OPQBot.EventNameOnOther, func(botQQ int64, e interface{}) {
		log.Println(e)
	})
	if err != nil {
		log.Println(err.Error())
	}
	time.Sleep(1 *time.Hour)
}
