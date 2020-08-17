package database

import "github.com/wonderivan/logger"

type t_sentence struct {
	Id					int
	Sentence_type		string
	Sentence_text		string
}
type TSentence t_sentence

// 获取句子列表
func GetSentenceList(index int, perpage int, stype int) ([]*TSentence, error) {
	t := []*TSentence{}
	err := Select(&t, "SELECT * FROM sentence WHERE sentence_type = ? ORDER BY id LIMIT ?,?",
		stype, index, perpage)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	return t, nil
}

// 根据类型获取总记录数
func GetSentenceCountByType(stype int) int {
	cnt := 0
	err := Get(&cnt, "SELECT count(*) FROM sentence WHERE sentence_type = ?",
		stype)
	if err != nil {
		logger.Error(err)
		return 0
	}

	return cnt
}

func GetSentenceByID(id int) (*t_sentence,error){
	t := t_sentence{}
	err := Get(&t, "SELECT * FROM sentence WHERE id = ?", id)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, nil
		}
	}
	return &t,nil
}

