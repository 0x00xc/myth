/**
 * singsen
 * SNQU
 * laixingchen@sndo.com
 * 2020/10/30 16:12
 */
package dao

type Follower struct {
	Model
	UID        int    `json:"uid"`         //
	FollowedID int    `json:"followed_id"` //
	Both       bool   `json:"both"`        //
	Group      string `json:"group"`       //
	Mark       string `json:"mark"`        //
}

func Follow(uid, followedID int, mark string, group string) error {
	f := new(Follower)
	if e := db.Model(&Follower{}).Where("uid = ? AND followed_id = ?", followedID, uid).Update("both", true).Error;
		e == nil {
		f.Both = true
	}
	f.UID = uid
	f.FollowedID = followedID
	f.Mark = mark
	f.Group = group
	return db.Save(f).Error
}

func EditFollowMark(id int, mark string) error {
	f := &Follower{}
	f.ID = id
	return db.Model(f).Update("mark", mark).Error
}

func EditFollowGroup(id int, group string) error {
	f := &Follower{}
	f.ID = id
	return db.Model(f).Update("group", group).Error
}

func Unfollow(uid, followedID int) error {
	f := new(Follower)
	db.Model(&Follower{}).Where("uid = ? AND followed_id = ?", followedID, uid).Update("both", false)
	return db.Delete(f, "uid = ? AND followed_id = ?", uid, followedID).Error
}

func FollowList(uid int, lastID int, limit int) ([]*Follower, error) {
	//var data []*Follower
	//err := db.Find(&data, "uid = ?", uid).Order("id DESC").Error
	//return data, err
	if limit == 0 {
		limit = 20
	}
	var data []*Follower
	var err error
	if lastID == 0 {
		err = db.Where("uid = ?", uid).
			Order("id DESC").Limit(limit).Find(&data).Error
	} else {
		err = db.Where("uid = ? AND id < ?", uid, lastID).
			Order("id DESC").Limit(limit).Find(&data).Error
	}
	return data, err
}

func FollowerList(uid int, lastID int, limit int) ([]*Follower, error) {
	if limit == 0 {
		limit = 20
	}
	var data []*Follower
	var err error
	if lastID == 0 {
		err = db.Where("followed_id = ?", uid).
			Order("id DESC").Limit(limit).Find(&data).Error
	} else {
		err = db.Where("followed_id = ? AND id < ?", uid, lastID).
			Order("id DESC").Limit(limit).Find(&data).Error
	}
	return data, err
}

func FindFollower(from, to int) (*Follower, error) {
	f := new(Follower)
	err := db.First(f, "uid = ? AND followed_id = ?", from, to).Error
	return f, err
}

func GetRelation(from, to int) int {
	if from == 0 {
		return RelationStranger
	}
	if from == to {
		return RelationOwn
	}
	f, err := FindFollower(from, to)
	if err != nil {
		return RelationStranger
	}
	if f.Both {
		return RelationFollowed
	}
	return RelationFollower
}

//func RemoveFollower(uid int, followerID int) error {
//
//}
