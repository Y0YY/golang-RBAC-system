package model

import (
	"encoding/json"
	"errors"
	"os"
)

type Role struct {
	RoleId   string `json:"roleId"`
	RoleName string `json:"roleName"`
}

func CreatRole(r Role) error {
	id := r.RoleId
	data, err := ReadData("roles_data")
	if err != nil {
		return err
	}
	var roles map[string]Role
	_ = json.Unmarshal(data, &roles)
	_, flag := roles[id]
	if flag {
		return errors.New("角色ID已存在")
	}
	if roles == nil {
		roles = make(map[string]Role)
	}
	roles[id] = r
	data, _ = json.Marshal(roles)
	_ = os.WriteFile("./data/roles_data.txt", data, 0666)
	return nil
}

func GetAllRoles() map[string]Role {
	data, _ := ReadData("roles_data")
	var roles map[string]Role
	_ = json.Unmarshal(data, &roles)
	return roles
}

/*
	func GetAllPermsOfRole(r Role) ([]Perm, error) {
		data, err := ReadData("role_perm_data")
		if err != nil {
			return nil, err
		}
		var m map[string][]Perm
		_ = json.Unmarshal(data, &m)
		return m[r.RoleId], nil
	}
*/
func GetAllPermsOfRole(r Role) ([]byte, error) {
	data, err := ReadData("role_perm_data")
	if err != nil {
		return nil, err
	}
	var m map[string][]Perm
	_ = json.Unmarshal(data, &m)
	return data, nil
}

// 通过角色名找角色
func FindRoleByName(roleName string) (*Role, error) {
	data, err := ReadData("roles_data")
	if err != nil {
		return nil, err
	}
	var roles map[string]Role
	_ = json.Unmarshal(data, &roles)
	for _, u := range roles {
		if u.RoleName == roleName {
			return &u, nil
		}
	}
	return nil, errors.New("该角色不存在")
}

func AddPerm(r Role, p Perm) error {
	data, err := ReadData("role_perm_data")
	if err != nil {
		return err
	}
	var m map[string][]Perm
	_ = json.Unmarshal(data, &m)
	if m == nil {
		m = make(map[string][]Perm)
	}
	_, flag := m[r.RoleId]
	if !flag {
		m[r.RoleId] = make([]Perm, 1)
		m[r.RoleId][0] = p
	} else {
		//去重
		for _, pi := range m[r.RoleId] {
			if pi.PermId == p.PermId {
				return errors.New("该角色已被分配过该权限")
			}
		}
		m[r.RoleId] = append(m[r.RoleId], p)
	}
	data, err = json.Marshal(m)
	if err != nil {
		return err
	}
	err = os.WriteFile("./data/role_perm_data.txt", data, 0666)
	if err != nil {
		return err
	}
	//return errors.New(fmt.Sprintf("%d", len(m[r])))
	return nil
}

func DeletePermOfRole(r Role, p Perm) error {
	data, err := ReadData("role_perm_data")
	if err != nil {
		return err
	}
	var m map[string][]Perm
	_ = json.Unmarshal(data, &m)

	ps, flag := m[r.RoleId]
	if !flag {
		return nil
	} else {
		nps := ps[:0]
		for _, perm := range ps {
			if perm.PermId != p.PermId {
				nps = append(nps, perm)
			}
		}
		m[r.RoleId] = nps
		data, _ = json.Marshal(m)
		_ = os.WriteFile("./data/role_perm_data.txt", data, 0666)
		return nil
	}
}

func DeleteRoleById(id string) error {
	//删除用户数据
	data, err := ReadData("roles_data")
	if err != nil {
		return err
	}
	var roles map[string]Role
	_ = json.Unmarshal(data, &roles)
	delete(roles, id)
	data, _ = json.Marshal(roles)
	_ = os.WriteFile("./data/roles_data.txt", data, 0666)
	//删除用户角色数据
	data, err = ReadData("user_role_data")
	if err != nil {
		return err
	}
	var urmap map[string][]Role
	_ = json.Unmarshal(data, &urmap)
	for _, roles := range urmap {
		for index, r := range roles {
			if r.RoleId == id {
				roles = append(roles[:index], roles[index+1:]...)
			}
		}
	}

	data, _ = json.Marshal(urmap)
	_ = os.WriteFile("./data/user_role_data.txt", data, 0666)
	//删除角色权限数据
	data, err = ReadData("role_perm_data")
	if err != nil {
		return err
	}
	var rpmap map[string][]Perm
	_ = json.Unmarshal(data, &rpmap)
	delete(rpmap, id)
	data, _ = json.Marshal(rpmap)
	_ = os.WriteFile("./data/role_perm_data.txt", data, 0666)
	return nil
}

func DeleteRoleByName(roleName string) error {
	r, err := FindRoleByName(roleName)
	if err != nil {
		return err
	}
	return DeleteUserById(r.RoleId)
}
