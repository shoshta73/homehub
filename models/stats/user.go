package stats

import "github.com/shoshta73/homehub/log"

func CheckUserStats(id int64) {
	log.Info("Checking user stats", "id", id)

	has, err := pastebinStatsExist(id)
	if err != nil {
		log.Error(err)
	}

	if !has {
		log.Info("User stats do not exist")
		err = createEmptyPastebinStats(id)
		if err != nil {
			log.Error(err)
		}
	}
}

func InitUserStats(id int64) {
	log.Info("Initializing user stats", "id", id)
	has, err := pastebinStatsExist(id)
	if err != nil {
		log.Error(err)
	}

	if has {
		log.Warn("User stats already exist")
		return
	}

	log.Info("User stats do not exist, creating")
	err = createEmptyPastebinStats(id)
	if err != nil {
		log.Error(err)
	}
	log.Info("User stats created", "id", id)
}

func IncrementPasteCreated(id int64) {
	log.Info("Incrementing paste created", "id", id)
	paste := &PastebinStats{}
	_, err := orm.Where("id = ?", id).Get(paste)
	if err != nil {
		log.Error(err)
		return
	}

	paste.Created++
	_, err = orm.Where("id = ?", id).Update(paste)
	if err != nil {
		log.Error(err)
	}
}
