package service

import (
    "fmt"
    // "reflect"

    "around/backend"
    "around/constants"
    "around/model"

    "github.com/olivere/elastic/v7"
)

func CheckUser(username, password string) (bool, error){
	query := elastic.NewBoolQuery()
    query.Must(elastic.NewTermQuery("username", username))
    query.Must(elastic.NewTermQuery("password", password))
    searchResult, err := backend.ESBackend.ReadFromES(query, constants.USER_INDEX)
    if err != nil {
        return false, err
    }

	return searchResult.TotalHits() > 0, nil
}

func AddUser(user *model.User) (bool, error){
	// 判断用户名是否冲突
	query := elastic.NewTermQuery("username", user.Username)
    searchResult, err := backend.ESBackend.ReadFromES(query, constants.USER_INDEX)
	if err != nil {
        return false, err
    }
	// 冲突, 有duplicate user
	if searchResult.TotalHits() > 0 {
        return false, nil
    }
	// 不冲突 saveToES
	err = backend.ESBackend.SaveToES(user, constants.USER_INDEX, user.Username)
    if err != nil {
        return false, err
    }
    fmt.Printf("User is added: %s\n", user.Username)
    return true, nil
}
