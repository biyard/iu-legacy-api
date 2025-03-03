package heroesSongContestSubmission

type HeroesSongContestSubmission struct {
	Key         string `bson:"key,omitempty" json:"key,omitempty"`
	Email       string `bson:"email,omitempty" json:"email,omitempty"`
	Address     string `bson:"address,omitempty" json:"address,omitempty"`
	Name        string `bson:"name,omitempty" json:"name,omitempty"`
	Phone       string `bson:"phone,omitempty" json:"phone,omitempty"`
	MainNFTID   int64  `bson:"mainNftId,omitempty" json:"mainNftId,omitempty"`
	AudioURL    string `bson:"audioUrl,omitempty" json:"audioUrl,omitempty"`
	ImageURL    string `bson:"imageUrl,omitempty" json:"imageUrl,omitempty"`
	Lyrics      string `bson:"lyrics,omitempty" json:"lyrics,omitempty"`
	SubmittedAt int64  `bson:"submittedAt,omitempty" json:"submittedAt,omitempty"`
	CreatedAt   int64  `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
}

func (r HeroesSongContestSubmission) Collection() string {
	return "heroesSongContestSubmission"
}

func (r HeroesSongContestSubmission) Table() string {
	return r.Collection()
}
