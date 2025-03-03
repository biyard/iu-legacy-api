package heroesSongContest

type HeroesSongContestCandidate struct {
	Key       string `bson:"key,omitempty" json:"key,omitempty"`
	Title     string `bson:"title,omitempty" json:"title,omitempty"`
	Lyrics    string `bson:"lyrics,omitempty" json:"lyrics,omitempty"`
	Address   string `bson:"address,omitempty" json:"address,omitempty"`
	AudioURL  string `bson:"audioUrl,omitempty" json:"audioUrl,omitempty"`
	ImageURL  string `bson:"imageUrl,omitempty" json:"imageUrl,omitempty"`
	GoogleID  string `bson:"googleId,omitempty" json:"googleId,omitempty"`
	Email     string `bson:"email,omitempty" json:"email,omitempty"`
	Nickname  string `bson:"nickname,omitempty" json:"nickname,omitempty"`
	Name      string `bson:"name,omitempty" json:"name,omitempty"`
	Phone     string `bson:"phone,omitempty" json:"phone,omitempty"`
	NFTID     int64  `bson:"nftId,omitempty" json:"nftId,omitempty"`
	PlayCount int64  `bson:"playCount,omitempty" json:"playCount,omitempty"`
	LikeCount int64  `bson:"likeCount,omitempty" json:"likeCount,omitempty"`
}

func (r HeroesSongContestCandidate) Collection() string {
	return "heroesSongContestCandidate"
}

func (r HeroesSongContestCandidate) Table() string {
	return r.Collection()
}
