package databases

// BaseDatabase is a base database
type BaseDatabase struct {
}

// func (d *BaseDatabase) exists(tx *sql.Tx, query string, args ...interface{}) (bool, error) {
// 	var exists bool
// 	query = fmt.Sprintf("SELECT exists (%v)", query)
// 	err := tx.QueryRow(query, args...).Scan(&exists)
// 	if err != nil && err != sql.ErrNoRows {
// 		return false, err
// 	}
// 	return exists, nil
// }
//
// func (d *BaseDatabase) selectID(tx *sql.Tx, query string, args ...interface{}) (int64, error) {
// 	var id int64
// 	err := tx.QueryRow(query, args...).Scan(&id)
// 	if err != nil {
// 		return 0, err
// 	}
// 	return id, nil
// }
