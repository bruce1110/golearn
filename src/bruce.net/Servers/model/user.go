package model

func GetOne(uid uint64) (string,error) {
	sql := "select nickname from user_1 where uid = ?";
	var name string
	err := Db.QueryRow(sql, uid).Scan(&name)
	if err != nil {
		return err.Error(),err
	}
	return name,nil
}
