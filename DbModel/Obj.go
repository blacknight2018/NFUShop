package DbModel

import "NFUShop/Config"

func SelectTableRecordById(tableName string, id int, condition map[string]interface{}, out interface{}) bool {
	db := Config.GetRandomSlaveDB()
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
	case *Banner:
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
	case *Banner:
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
		err = db.Model(&Order{}).Where("id = ?", v.Id).Update(v).Error
		break
	case *Address:
		err = db.Model(&Address{}).Where("id = ?", v.Id).Update(v).Error
		break
	case *Goods:
		err = db.Model(&Goods{}).Where("id = ?", v.Id).Update(v).Error
		break
	case *SubGoods:
		err = db.Model(&SubGoods{}).Where("id = ?", v.Id).Update(v).Error
		break
	case *Cart:
		err = db.Model(&Cart{}).Where("id = ?", v.Id).Update(v).Error
		break
	case *Order:
		err = db.Model(&Order{}).Where("id = ?", v.Id).Update(v).Error
		break
	case *Banner:
		err = db.Model(&Banner{}).Where("id = ?", v.Id).Update(v).Error
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
	case *Banner:
		err = db.Delete(v).Error
		break
	case *Order:
		err = db.Delete(v).Error
		break
	}
	return err == nil
}

func SelectTableRecordSet(tableName string, out interface{}, condition map[string]interface{}, likeCondition map[string]string, limit *int, offset *int, order string) bool {
	db := Config.GetRandomSlaveDB()
	if db == nil {
		return false
	}
	dbCondition := db.Table(tableName).Where(condition).Order(order)
	if likeCondition != nil {
		for s := range likeCondition {
			dbCondition = dbCondition.Where(s+" like ?", "%"+likeCondition[s]+"%")
		}
	}
	if limit != nil {
		dbCondition = dbCondition.Limit(*limit)
	}
	if offset != nil {
		dbCondition = dbCondition.Offset(*offset)
	}
	var err error
	switch v := out.(type) {
	case *[]User:
		err = dbCondition.Find(v).Error
		break
	case *[]Order:
		err = dbCondition.Find(v).Error
		break
	case *[]Cart:
		err = dbCondition.Find(v).Error
		break
	case *[]Goods:
		err = dbCondition.Find(v).Error
		break
	case *[]SubGoods:
		err = dbCondition.Find(v).Error
		break
	case *[]Address:
		err = dbCondition.Find(v).Error
		break
	case *[]Banner:
		err = dbCondition.Find(v).Error
		break
	}
	defer db.Close()

	return err == nil
}
