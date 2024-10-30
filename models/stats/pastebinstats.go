package stats

import "github.com/shoshta73/homehub/log"

type PastebinStats struct {
	Id           int64 `xorm:"unique" json:"id"`
	Created      int64 `json:"created"`
	SharedWithMe int64 `json:"sharedWithMe"`
}

func (ps PastebinStats) TableName() string {
	return "pastebin_stats"
}

func createEmptyPastebinStats(id int64) error {
	log.Info("Creating empty pastebin stats for user", "id", id)
	pbs := PastebinStats{
		Id:           id,
		Created:      0,
		SharedWithMe: 0,
	}

	_, err := orm.Insert(&pbs)
	if err != nil {
		return err
	}
	return nil
}

func pastebinStatsExist(id int64) (bool, error) {
	return orm.Get(&PastebinStats{Id: id})
}

func GetPastebinStats(id int64) (*PastebinStats, error) {
	pbs := &PastebinStats{}
	_, err := orm.Get(pbs, &PastebinStats{Id: id})
	if err != nil {
		return nil, err
	}

	return pbs, nil
}
