package model

import (
	"encoding/json"
	"errors"
	"os"
)

type User struct {
	UserId   string `json:"userId"`
	UserName string `json:"userName"`
}

func CreatUser(u User) error {
	id := u.UserId
	data, err := ReadData("users_data")
	if err != nil {
		return err
	}
	var users map[string]User
	_ = json.Unmarshal(data, &users)
	_, flag := users[id]
	if flag {
		return errors.New("用户ID已存在")
	}
	if users == nil {
		users = make(map[string]User)
	}
	users[id] = u
	data, _ = json.Marshal(users)
	_ = os.WriteFile("./data/users_data.txt", data, 0666)
	return nil
}

func GetAllUsers() map[string]User {
	data, _ := ReadData("users_data")
	var users map[string]User
	_ = json.Unmarshal(data, &users)
	return users
}

func GetAllRolesOfUser(u User) ([]Role, error) {
	data, err := ReadData("user_role_data")
	if err != nil {
		return nil, err
	}
	var m map[string][]Role
	_ = json.Unmarshal(data, &m)
	return m[u.UserId], nil
}

func GetAllPermsOfUser(u User) ([]Perm, error) {
	data1, _ := ReadData("user_role_data")
	data2, _ := ReadData("role_perm_data")
	var urmap map[string][]Role
	var rpmap map[string][]Perm
	_ = json.Unmarshal(data1, &urmap)
	_ = json.Unmarshal(data2, &rpmap)
	set := make(map[Perm]struct{})
	for _, r := range urmap[u.UserId] {
		for _, pm := range rpmap[r.RoleId] {
			set[pm] = struct{}{}
		}
	}
	res := make([]Perm, len(set))
	index := 0
	for key := range set {
		res[index] = key
		index++
	}
	return res, nil
}

func AddRole(u User, r Role) error {
	data, err := ReadData("user_role_data")
	if err != nil {
		return err
	}
	var m map[string][]Role
	_ = json.Unmarshal(data, &m)
	if m == nil {
		m = make(map[string][]Role)
	}
	_, flag := m[u.UserId]
	if !flag {
		m[u.UserId] = make([]Role, 1)
		m[u.UserId][0] = r
	} else {
		//去重
		for _, ri := range m[u.UserId] {
			if ri.RoleId == r.RoleId {
				return errors.New("该用户已被分配过该角色")
			}
		}
		m[u.UserId] = append(m[u.UserId], r)
	}
	data, _ = json.Marshal(m)
	_ = os.WriteFile("./data/user_role_data.txt", data, 0666)
	return nil
}

func DeleteRoleOfUser(u User, r Role) error {
	data, err := ReadData("user_role_data")
	if err != nil {
		return err
	}
	var m map[string][]Role
	_ = json.Unmarshal(data, &m)

	rs, flag := m[u.UserId]
	if !flag {
		return nil
	} else {
		nrs := rs[:0]
		for _, role := range rs {
			if role.RoleId != r.RoleId {
				nrs = append(nrs, role)
			}
		}
		m[u.UserId] = nrs
		data, _ = json.Marshal(m)
		_ = os.WriteFile("./data/user_role_data.txt", data, 0666)
		return nil
	}
}

func FindUserByName(userName string) (*User, error) {
	data, err := ReadData("users_data")
	if err != nil {
		return nil, err
	}
	var users map[string]User
	_ = json.Unmarshal(data, &users)
	for _, u := range users {
		if u.UserName == userName {
			return &u, nil
		}
	}
	return nil, errors.New("该用户不存在")
}

func DeleteUserById(id string) error {
	//删除用户数据
	data, err := ReadData("users_data")
	if err != nil {
		return err
	}
	var users map[string]User
	_ = json.Unmarshal(data, &users)
	delete(users, id)
	data, _ = json.Marshal(users)
	_ = os.WriteFile("./data/users_data.txt", data, 0666)
	//删除用户角色数据
	data, err = ReadData("user_role_data")
	if err != nil {
		return err
	}
	var urmap map[string][]Role
	_ = json.Unmarshal(data, &urmap)
	delete(urmap, id)
	data, _ = json.Marshal(urmap)
	_ = os.WriteFile("./data/user_role_data.txt", data, 0666)
	return nil
}

func DeleteUserByName(userName string) error {
	u, err := FindUserByName(userName)
	if err != nil {
		return err
	}
	return DeleteUserById(u.UserId)
}

func IsPrmitted(u User, p Perm) bool {
	data1, _ := ReadData("user_role_data")
	data2, _ := ReadData("role_perm_data")
	var urmap map[string][]Role
	var rpmap map[string][]Perm
	_ = json.Unmarshal(data1, &urmap)
	_ = json.Unmarshal(data2, &rpmap)
	urs, flag := urmap[u.UserId]
	if !flag {
		return false
	}
	for _, r := range urs {
		for _, pm := range rpmap[r.RoleId] {
			if pm.PermName == p.PermName {
				return true
			}
		}
	}
	return false
}
