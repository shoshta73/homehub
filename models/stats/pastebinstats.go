package stats

import "github.com/shoshta73/homehub/log"

type PastebinStats struct {
	Id           int64 `xorm:"unique"`
	Created      int64
	SharedWithMe int64
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
