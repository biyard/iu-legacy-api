package heroesSongContest

type HeroesSongContestLike struct {
	Key             string   `bson:"key,omitempty" json:"key,omitempty"`
	CandidateIDList []string `bson:"candidateIdList,omitempty" json:"candidateIdList,omitempty"`
	CreatedAt       int64    `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
}

func (r HeroesSongContestLike) Collection() string {
	return "heroesSongContestLike"
}

func (r HeroesSongContestLike) Table() string {
	return r.Collection()
}
