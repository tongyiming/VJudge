/**
 * Created by shiyi on 2017/10/16.
 * Email: shiyi@fightcoder.com
 */

package dispatch

//
//import (
//	"encoding/json"
//	"fmt"
//
//	"self/judger"
//
//	"github.com/nsqio/go-nsq"
//	log "github.com/sirupsen/logrus"
//)
//
//type Handler struct {
//	Topic string
//}
//
//func (this *Handler) HandleMessage(m *nsq.Message) error {
//	fmt.Println(string(m.Body))
//
//	judgeEvent := new(judger.JudgeEvent)
//	if err := json.Unmarshal(m.Body, judgeEvent); err != nil {
//		log.Errorf("unmarshal judgeEvent from nsqdata failed, err: %v, event:%s", err, m.Body)
//		return err
//	}
//
//	fmt.Printf("%#v\n", judgeEvent)
//
//	handlerCount <- 1
//	go this.doJudge(judgeEvent)
//
//	return nil
//}
//
//func (this *Handler) doJudge(judgeEvent *judger.JudgeEvent) {
//	defer func() {
//		<-handlerCount
//	}()
//
//	judgeEvent.DoJudge()
//}
