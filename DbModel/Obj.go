package DbModel

import "NFUShop/Config"

func SelectTableRecordById(tableName string, id int, condition map[string]interface{}, out interface{}) bool {
	db := Config.GetOneDB()
	if db == nil {
		return false
	}
	defer db.Close()
	var err error
	switch v := out.(type) {
	case *User:
		err = db.Table(tableName).Where("id = ?", id).Where(condition).First(v).Error
		break
	case *Address:
		err = db.Table(tableName).Where("id = ?", id).Where(condition).First(v).Error
		break
	case *Goods:
		err = db.Table(tableName).Where("id = ?", id).Where(condition).First(v).Error
		break
	case *SubGoods:
		err = db.Table(tableName).Where("id = ?", id).Where(condition).First(v).Error
		break
	case *Order:
		err = db.Table(tableName).Where("id = ?", id).Where(condition).First(v).Error
		break
	case *Cart:
		err = db.Table(tableName).Where("id = ?", id).Where(condition).First(v).Error
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
		db.Model(&Address{}).Where("id = ?", v.Id).Update(v)
		break
	case *Goods:
		db.Model(&Goods{}).Where("id = ?", v.Id).Update(v)
		break
	case *SubGoods:
		db.Model(&SubGoods{}).Where("id = ?", v.Id).Update(v)
		break
	case *Cart:
		db.Model(&Cart{}).Where("id = ?", v.Id).Update(v)
		break
	case *Order:
		db.Model(&Order{}).Where("id = ?", v.Id).Update(v)
		break
	}
	return err == nil
}

func DeleteDBObj(in interface{}) bool {
	db := Config.GetOneDB()
	if db == nil {
		return false
	}
	defer db.Close()
	var err error
	switch v := in.(type) {
	case *Cart:
		err = db.Delete(v).Error
		break
	case *Address:
		err = db.Delete(v).Error
		break
	}
	return err == nil
}

func SelectTableRecordSet(tableName string, out interface{}, condition map[string]interface{}, limit int, offset int, order string) bool {
	db := Config.GetOneDB()
	if db == nil {
		return false
	}
	var err error
	switch v := out.(type) {
	case *[]User:
		err = db.Table(tableName).Where(condition).Limit(limit).Offset(offset).Order(order).Find(v).Error
		break
	case *[]Order:
		err = db.Table(tableName).Where(condition).Limit(limit).Offset(offset).Order(order).Find(v).Error
		break
	case *[]Cart:
		err = db.Table(tableName).Where(condition).Limit(limit).Offset(offset).Order(order).Find(v).Error
		break
	case *[]Goods:
		err = db.Table(tableName).Where(condition).Limit(limit).Offset(offset).Order(order).Find(v).Error
		break
	case *[]SubGoods:
		err = db.Table(tableName).Where(condition).Limit(limit).Offset(offset).Order(order).Find(v).Error
		break
	case *[]Address:
		err = db.Table(tableName).Where(condition).Limit(limit).Offset(offset).Order(order).Find(v).Error
		break
	}
	defer db.Close()

	return err == nil
}
