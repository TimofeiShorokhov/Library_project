package repo

import (
	"Library_project/other"
	"fmt"
	"strconv"
)

type Instance struct {
	InstanceId    uint16  `json:"instance_id"`
	InstanceName  string  `json:"instance_name"`
	Damage        string  `json:"damage"`
	InstancePrice uint16  `json:"instance_price"`
	ReturnDate    string  `json:"return_date"`
	DmgPhoto      string  `json:"dmg_photo"`
	Rating        uint16  `json:"rating"`
	FinalPrice    float64 `json:"final_price"`
}

func SaveInstanceInDB(book Book) {
	db := other.ConnectDB()
	defer db.Close()
	var i uint16
	for i = 0; i < book.Quantity; i++ {
		ins, err := db.Query(fmt.Sprintf("INSERT INTO `instances` (`instance_name`,`instance_price`) VALUES ('%s','%d')", book.BookName, book.Price))
		other.CheckErr(err)
		if i == book.Quantity {
			ins.Close()
		}
	}

}

func UpdateInstancesInDB(instance Instance) {
	db := other.ConnectDB()
	defer db.Close()

	updBook := db.QueryRow("UPDATE `instances` SET damage= ?, rating = ?, return_date =?, instance_price =?, dmg_photo=?  where instance_id =? ", instance.Damage, instance.Rating, instance.ReturnDate, instance.InstancePrice, instance.DmgPhoto, instance.InstanceId)
	updBook.Err()
}

func GetInstancesFromDB(Instances *[]Instance) {
	db := other.ConnectDB()
	defer db.Close()
	get, err := db.Query("Select * from `instances`")
	other.CheckErr(err)

	for get.Next() {
		var instance Instance
		err = get.Scan(&instance.InstanceId, &instance.InstanceName, &instance.Damage, &instance.InstancePrice, &instance.ReturnDate, &instance.DmgPhoto, &instance.Rating)
		other.CheckErr(err)
		*Instances = append(*Instances, instance)
	}
}

func GetInstancesFromDBWithPage(Instances *[]Instance, page string, limit string) {
	db := other.ConnectDB()
	defer db.Close()
	p, _ := strconv.Atoi(page)
	pageForSql := (p - 1) * 5
	l, _ := strconv.Atoi(limit)
	get, err := db.Query(fmt.Sprintf("Select * from `instances` order by instance_id LIMIT %d OFFSET %d", l, pageForSql))
	other.CheckErr(err)

	for get.Next() {
		var instance Instance
		err = get.Scan(&instance.InstanceId, &instance.InstanceName, &instance.Damage, &instance.InstancePrice, &instance.ReturnDate, &instance.DmgPhoto, &instance.Rating)
		other.CheckErr(err)
		*Instances = append(*Instances, instance)
	}
}
