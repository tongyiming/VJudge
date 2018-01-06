/**
 * Created by shiyi on 2017/10/16.
 * Email: shiyi@fightcoder.com
 */

package dispatch

import (
	"encoding/json"

	"self/vjudge"

	"github.com/nsqio/go-nsq"
	log "github.com/sirupsen/logrus"
)

type Handler struct {
	Topic string
}

func (this *Handler) HandleMessage(m *nsq.Message) error {
	vjudgeData := new(vjudge.VJudge)
	if err := json.Unmarshal(m.Body, vjudgeData); err != nil {
		log.Errorf("unmarshal JudgerData from NsqMessage failed, err: %v, event:%s", err, m.Body)
		return nil
	}

	log.Infof("consume Message from dispatch: %#v", vjudgeData)

	handlerCount <- 1
	go this.doJudge(vjudgeData)

	return nil
}

func (this *Handler) doJudge(vjudgeData *vjudge.VJudge) {
	defer func() {
		<-handlerCount
	}()

	vjudgeData.DoJudge()
}
