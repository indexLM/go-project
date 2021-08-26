package service

import (
	"github.com/xuri/excelize/v2"
	"go-project/global"
	"go-project/model/po"
	"go-project/model/resp"
	"golang.org/x/crypto/bcrypt"
	"mime/multipart"
	"strconv"
)

func DistrictList(cityName string) *[]resp.DistrictListRes {
	sql := "select distinct\n       hb.city_id,\n       hb.city_name,\n       hb.district_id,\n       hb.district_name\nfrom his_branch hb\nwhere city_name like ?"
	districtListRes := make([]resp.DistrictListRes, 0)
	err := global.MySqlx.Select(&districtListRes, sql, cityName+"%")
	if err != nil {
		panic(err.Error())
	}
	return &districtListRes
}
func BranchList(cityId string, districtId string, branchName string) *[]resp.BranchListRes {
	sql := "select hb.id        as branch_id,\n       hb.name      as branch_name,\n       hb.address   as branch_address,\n       hb.longitude as lon,\n       hb.latitude  as lat\nfrom his_branch hb\n where city_id = ?\n  and district_id = ? \n  "
	res := make([]resp.BranchListRes, 0)
	if branchName != "" {
		sql = sql + " and hb.name like ?"
		err := global.MySqlx.Select(&res, sql, cityId, districtId, "%"+branchName+"%")
		if err != nil {
			panic(err.Error())
		}
	} else {
		err := global.MySqlx.Select(&res, sql, cityId, districtId)
		if err != nil {
			panic(err.Error())
		}
	}
	return &res
}

func BranchImport(file *multipart.FileHeader) error {
	open, err := file.Open()
	if err != nil {
		return err
	}
	defer open.Close()
	reader, err := excelize.OpenReader(open)
	if err != nil {
		return err
	}
	sheetName := "门店列表"
	rows, err := reader.Rows(sheetName)
	if err != nil {
		return err
	}
	branchList := make([]map[string]interface{}, 0)
	//reader.GetCellValue(sheetName,)
	num := 0
	for rows.Next() {
		if num == 0 {
			num++
			_, _ = rows.Columns()
			continue
		}
		cols, _ := rows.Columns()
		m := make(map[string]interface{}, 0)
		for i, e := range cols {
			if i == 0 {
				//名称
				m["name"] = e
			} else if i == 1 {
				//地址
				m["address"] = e
			} else if i == 2 {
				//经度
				m["lon"] = e
			} else if i == 3 {
				//维度
				m["lat"] = e
			} else if i == 4 {
				//所在区
				m["districtName"] = e
				sql := "select ha.id from his_area ha where ha.name like ? and ha.level=3 limit 1"
				id := new(string)
				_ = global.MySqlx.Get(id, sql, e+"%")
				m["districtId"] = id
			} else if i == 5 {
				//所在城市
				m["cityName"] = e
				sql := "select ha.id from his_area ha where ha.name like ? and ha.level=2 limit 1"
				id := new(string)
				_ = global.MySqlx.Get(id, sql, e+"%")
				m["cityId"] = id
			}
		}
		branchList = append(branchList, m)
	}
	tx := global.MySqlx.MustBegin()
	inSql := `insert into his_branch(name,address,latitude,longitude,district_id,district_name,city_id,city_name,type_id) values (:name,:address,:lat,:lon,:districtId,:districtName,:cityId,:cityName,1)`
	_, err = tx.NamedExec(inSql, branchList)
	if err != nil {
		panic(err.Error())
	}
	err = tx.Commit()
	if err != nil {
		panic(err.Error())
	}
	return nil
}

func BranchAdminReport(name string) []byte {
	sql := "select id from his_branch where id not in (select branch_id from his_branch_admin)"
	idList := make([]string, 0)
	err := global.MySqlx.Select(&idList, sql)
	if err != nil {
		panic(err.Error())
	}
	password, err := bcrypt.GenerateFromPassword([]byte("000000"), bcrypt.DefaultCost)
	if err != nil {
		panic(err.Error())
	}
	s := string(password)
	if len(idList) != 0 {
		inList := make([]map[string]interface{}, 0)
		for _, m := range idList {
			m2 := make(map[string]interface{}, 0)
			m2["name"] = name + m
			m2["password"] = s
			m2["branchId"] = m
			inList = append(inList, m2)
		}
		_, err := global.MySqlx.NamedExec("insert into his_branch_admin(branch_id, username, password) VALUES (:branchId,:name,:password)", inList)
		if err != nil {
			panic(err.Error())
		}
	}
	file := excelize.NewFile()
	sheetName := "账号列表"
	newSheet := file.NewSheet(sheetName)
	//查询数据库
	res := make([]po.BranchAdmin, 0)
	err = global.MySqlx.Select(&res, "select hb.name,hba.username,hba.password from his_branch hb left join his_branch_admin hba on hb.id=hba.branch_id ")
	if err != nil {
		panic(err.Error())
	}
	_ = file.SetCellValue(sheetName, "A1", "分院名称")
	_ = file.SetCellValue(sheetName, "B1", "账号")
	_ = file.SetCellValue(sheetName, "C1", "密码")
	for index, element := range res {
		for i := 1; i < 4; i++ {
			if i == 1 {
				_ = file.SetCellValue(sheetName, "A"+strconv.FormatUint(uint64(index+2), 10), *element.Name)
			}
			if i == 2 {
				_ = file.SetCellValue(sheetName, "B"+strconv.FormatUint(uint64(index+2), 10), *element.Username)
			}
			if i == 3 {
				_ = file.SetCellValue(sheetName, "C"+strconv.FormatUint(uint64(index+2), 10), "000000")
			}
		}
	}
	file.SetActiveSheet(newSheet)
	buffer, err := file.WriteToBuffer()
	if err != nil {
		panic(err.Error())
	}
	bytes := buffer.Bytes()
	return bytes
}
func DoctorImport(file *multipart.FileHeader) error {
	open, err := file.Open()
	if err != nil {
		return err
	}
	defer open.Close()
	reader, err := excelize.OpenReader(open)
	if err != nil {
		return err
	}
	sheetName := "医生列表"
	rows, err := reader.Rows(sheetName)
	if err != nil {
		return err
	}
	branchList := make([]map[string]interface{}, 0)
	//reader.GetCellValue(sheetName,)
	num := 0
	for rows.Next() {
		if num == 0 {
			num++
			_, _ = rows.Columns()
			continue
		}
		cols, _ := rows.Columns()
		m := make(map[string]interface{}, 0)
		for i, e := range cols {
			if i == 0 {
				//id
				m["id"] = e
			} else if i == 1 {
				//姓名
				m["name"] = e
			} else if i == 2 {
				//简介
				m["intro"] = e
			} else if i == 3 {
				//职级
				m["professional_name"] = e
			} else if i == 4 {
				//医院
				sql := "select id from his_holiday where name like ? limit 1"
				id := new(string)
				_ = global.MySqlx.Get(*id, sql, "%"+e+"%")
				m["hospital_id"] = id
			}
		}
		branchList = append(branchList, m)
	}
	tx := global.MySqlx.MustBegin()
	inSql := `INSERT INTO lejian_hospital.his_doctor (id, name, mobile, intro, professional_name,status,password) VALUES
(:id, :name, DEFAULT, :intro,:professional_name, 1,'$2a$10$WalS92/Bxc31GlgfL6UqGuKBg204EqQWVGAHpEMG2VZe9Orz.h/wC')`
	_, err = tx.NamedExec(inSql, branchList)
	if err != nil {
		panic(err.Error())
	}
	ninSql := `INSERT INTO lejian_hospital.his_channel_doctor (doctor_id, channel_id) VALUES (:id, 2);`
	_, err = tx.NamedExec(ninSql, branchList)
	if err != nil {
		panic(err.Error())
	}
	err = tx.Commit()
	if err != nil {
		panic(err.Error())
	}
	return nil
}
