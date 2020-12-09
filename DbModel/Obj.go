package DbModel

import "NFUShop/Config"

func FindTableById(tableName string, id int, out interface{}) bool {
	db := Config.GetOneDB()
	if db == nil {
		return false
	}
	defer db.Close()
	var err error
	switch v := out.(type) {
	case *User:
		err = db.Table(tableName).Where("id = ?", id).First(v).Error
		break
	case *Address:
		err = db.Table(tableName).Where("id = ?", id).First(v).Error
		break
	case *Goods:
		err = db.Table(tableName).Where("id = ?", id).First(v).Error
		break
	case *SubGoods:
		err = db.Table(tableName).Where("id = ?", id).First(v).Error
		break
	case *Order:
		err = db.Table(tableName).Where("id = ?", id).First(v).Error
		break
	case *Cart:
		err = db.Table(tableName).Where("id = ?", id).First(v).Error
		break
	}
	return err == nil
}

func InsertDBObj(in interface{}) bool {
	db := Config.GetOneDB()
	if db == nil {
		return false
	}
	defer db.Close()
	var err error
	switch v := in.(type) {
	case *User:
		err = db.Create(v).Error
		break
	case *Address:
		err = db.Create(v).Error
		break
	case *Goods:
		err = db.Create(v).Error
		break
	case *Cart:
		err = db.Create(v).Error
		break
	case *Order:
		err = db.Create(v).Error
		break
	case *SubGoods:
		err = db.Create(v).Error
		break
	}
	return err == nil
}

func UpdateDBObj(in interface{}) bool {
	db := Config.GetOneDB()
	if db == nil {
		return false
	}
	defer db.Close()
	var err error
	switch v := in.(type) {
	case *User:
		db.Model(&User{}).Where("id = ?", v.Id).Update(v)
		break
	case *Address:
		db.Model(&User{}).Where("id = ?", v.Id).Update(v)
		break
	case *Goods:
		db.Model(&User{}).Where("id = ?", v.Id).Update(v)
		break
	case *SubGoods:
		db.Model(&User{}).Where("id = ?", v.Id).Update(v)
		break
	case *Cart:
		db.Model(&User{}).Where("id = ?", v.Id).Update(v)
		break
	case *Order:
		db.Model(&User{}).Where("id = ?", v.Id).Update(v)
		break
	}
	return err == nil
}
