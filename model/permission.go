package model

import (
	"encoding/json"
	"errors"
	"os"
)

type Perm struct {
	PermId   string `json:"permId"`
	PermName string `json:"permName"`
}

func CreatPerm(p Perm) error {
	id := p.PermId
	data, err := ReadData("perms_data")
	if err != nil {
		return err
	}
	var perms map[string]Perm
	_ = json.Unmarshal(data, &perms)
	_, flag := perms[id]
	if flag {
		return errors.New("ID已存在")
	}
	if perms == nil {
		perms = make(map[string]Perm)
	}
	perms[id] = p
	data, _ = json.Marshal(perms)
	_ = os.WriteFile("./data/perms_data.txt", data, 0666)
	return nil
}

func ShowAllPerms() map[string]Perm {
	data, _ := ReadData("perms_data")
	var perms map[string]Perm
	_ = json.Unmarshal(data, &perms)
	return perms
}

func FindPermByName(permName string) (*Perm, error) {
	data, err := ReadData("perms_data")
	if err != nil {
		return nil, err
	}
	var perms map[string]Perm
	_ = json.Unmarshal(data, &perms)
	for _, p := range perms {
		if p.PermName == permName {
			return &p, nil
		}
	}
	return nil, errors.New("该权限不存在")
}

func DeletePermById(id string) error {
	//删除权限数据
	data, err := ReadData("perms_data")
	if err != nil {
		return err
	}
	var perms map[string]Perm
	_ = json.Unmarshal(data, &perms)
	//_, flag := perms[id]
	//if !flag {
	//	return errors.New("该权限(" + id + ")不存在")
	//}
	delete(perms, id)
	data, _ = json.Marshal(perms)
	_ = os.WriteFile("./data/perms_data.txt", data, 0666)
	//删除角色权限数据
	data, err = ReadData("role_perm_data")
	if err != nil {
		return err
	}
	var rpmap map[string][]Perm
	_ = json.Unmarshal(data, &rpmap)

	for index, ps := range rpmap {
		nps := ps[:0]
		for _, p := range ps {
			if p.PermId != id {
				nps = append(nps, p)
			}
		}
		rpmap[index] = nps
	}
	data, _ = json.Marshal(rpmap)
	err = os.WriteFile("./data/role_perm_data.txt", data, 0666)
	if err != nil {
		return err
	}
	return nil
}

func DeletePermByName(roleName string) error {
	r, err := FindRoleByName(roleName)
	if err != nil {
		return err
	}
	return DeleteUserById(r.RoleId)
}
