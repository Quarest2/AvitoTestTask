package db

import (
	"backend/db"
	"database/sql"
	"errors"
	"time"
)

func Get(params GetTendersParams) (tenders []db.Tender, err error) {
	if db.MyDB == nil {
		err = errors.New("database is not connected")
		return
	}

	var rows *sql.Rows
	if params.Limit > 0 {
		if params.Offset > 0 {
			if len(params.ServiceType) == 2 {
				if rows, err = db.MyDB.Query(`SELECT id, name, description, status, service_type, version, created_at FROM tender
                WHERE service_type = $1 OR service_type = $2 
                OFFSET $3 ROWS
				FETCH NEXT $4 ROWS ONLY`,
					params.ServiceType[0], params.ServiceType[1], params.Offset, params.Limit); err != nil {
					return
				}
			}
			if len(params.ServiceType) == 1 {
				if rows, err = db.MyDB.Query(`SELECT id, name, description, status, service_type, version, created_at FROM tender
                WHERE service_type = $1 
                OFFSET $2 ROWS
				FETCH NEXT $3 ROWS ONLY`,
					params.ServiceType[0], params.Offset, params.Limit); err != nil {
					return
				}
			}
			if len(params.ServiceType) == 0 || len(params.ServiceType) >= 3 {
				if rows, err = db.MyDB.Query(`SELECT id, name, description, status, service_type, version, created_at FROM tender 
                OFFSET $2 ROWS
				FETCH NEXT $3 ROWS ONLY`,
					params.ServiceType[0], params.Offset, params.Limit); err != nil {
					return
				}
			}
		} else {
			if len(params.ServiceType) == 2 {
				if rows, err = db.MyDB.Query(`SELECT id, name, description, status, service_type, version, created_at FROM tender
                WHERE service_type = $1 OR service_type = $2 
				FETCH NEXT $3 ROWS ONLY`,
					params.ServiceType[0], params.ServiceType[1], params.Limit); err != nil {
					return
				}
			}
			if len(params.ServiceType) == 1 {
				if rows, err = db.MyDB.Query(`SELECT id, name, description, status, service_type, version, created_at FROM tender
                WHERE service_type = $1 
				FETCH NEXT $3 ROWS ONLY`,
					params.ServiceType[0], params.Limit); err != nil {
					return
				}
			}
			if len(params.ServiceType) == 0 || len(params.ServiceType) >= 3 {
				if rows, err = db.MyDB.Query(`SELECT id, name, description, status, service_type, version, created_at FROM tender 
				FETCH NEXT $3 ROWS ONLY`, params.Limit); err != nil {
					return
				}
			}
		}
	} else {
		if params.Offset > 0 {
			if len(params.ServiceType) == 2 {
				if rows, err = db.MyDB.Query(`SELECT id, name, description, status, service_type, version, created_at FROM tender
                WHERE service_type = $1 OR service_type = $2 
                OFFSET $3 ROWS`,
					params.ServiceType[0], params.ServiceType[1], params.Offset); err != nil {
					return
				}
			}
			if len(params.ServiceType) == 1 {
				if rows, err = db.MyDB.Query(`SELECT id, name, description, status, service_type, version, created_at FROM tender
                WHERE service_type = $1 
                OFFSET $2 ROWS`,
					params.ServiceType[0], params.Offset); err != nil {
					return
				}
			}
			if len(params.ServiceType) == 0 || len(params.ServiceType) >= 3 {
				if rows, err = db.MyDB.Query(`SELECT id, name, description, status, service_type, version, created_at FROM tender 
                OFFSET $2 ROWS`, params.Offset); err != nil {
					return
				}
			}
		} else {
			if len(params.ServiceType) == 2 {
				if rows, err = db.MyDB.Query(`SELECT id, name, description, status, service_type, version, created_at FROM tender
                WHERE service_type = $1 OR service_type = $2`,
					params.ServiceType[0], params.ServiceType[1]); err != nil {
					return
				}
			}
			if len(params.ServiceType) == 1 {
				if rows, err = db.MyDB.Query(`SELECT id, name, description, status, service_type, version, created_at FROM tender
                WHERE service_type = $1`,
					params.ServiceType[0]); err != nil {
					return
				}
			}
			if len(params.ServiceType) == 0 || len(params.ServiceType) >= 3 {
				if rows, err = db.MyDB.Query(`SELECT id, name, description, status, service_type, version, created_at FROM tender`); err != nil {
					return
				}
			}
		}
	}

	defer rows.Close()
	if rows == nil {
		return nil, errors.New("no rows returned")
	}

	for rows.Next() {
		var t db.Tender
		if err = rows.Scan(&t.Id, &t.Name, &t.Description, &t.Status, &t.ServiceType, &t.Version, &t.CreatedAt); err != nil {
			return
		}

		tenders = append(tenders, t)
	}

	return
}

func New(params NewTenderParams) (tender db.Tender, err error) {
	if db.MyDB == nil {
		err = errors.New("database is not connected")
		return
	}

	timeNow := time.Now()

	var row *sql.Rows
	if row, err = db.MyDB.Query(`INSERT INTO tender (name, description, service_type, organization_id, creator_username, created_at) VALUES
                ($1, $2, $3, $4, $5, $6) RETURNING id`, params.Name, params.Description, params.ServiceType, params.OrganizationId, params.CreatorUsername, timeNow); err != nil {
		return
	}

	var id string
	for row.Next() {
		if err = row.Scan(&id); err != nil {
			return
		}
	}

	tender = db.Tender{Id: id, Name: params.Name, Description: params.Description, Status: "Created", ServiceType: params.ServiceType, Version: 1, CreatedAt: timeNow}

	return
}

func GetMyTenders(params GetMyTendersParams) (tenders []db.Tender, err error) {
	if db.MyDB == nil {
		err = errors.New("database is not connected")
		return
	}

	var rows *sql.Rows
	if params.Limit > 0 && params.Offset > 0 {
		if rows, err = db.MyDB.Query(`SELECT id, name, description, status, service_type, version, created_at FROM tender WHERE creator_username = $1
                OFFSET $2 ROWS
				FETCH NEXT $3 ROWS ONLY`,
			params.Username, params.Offset, params.Limit); err != nil {
			return
		}
	}
	if params.Limit > 0 && params.Offset <= 0 {
		if rows, err = db.MyDB.Query(`SELECT id, name, description, status, service_type, version, created_at FROM tender WHERE creator_username = $1
				FETCH NEXT $2 ROWS ONLY`,
			params.Username, params.Limit); err != nil {
			return
		}
	}
	if params.Limit <= 0 && params.Offset > 0 {
		if rows, err = db.MyDB.Query(`SELECT id, name, description, status, service_type, version, created_at FROM tender WHERE creator_username = $1
                OFFSET $2 ROWS`,
			params.Username, params.Offset); err != nil {
			return
		}
	}
	if params.Limit <= 0 && params.Offset <= 0 {
		if rows, err = db.MyDB.Query(`SELECT id, name, description, status, service_type, version, created_at FROM tender WHERE creator_username = $1 ORDER BY id`, params.Username); err != nil {
			return
		}
	}

	defer rows.Close()
	if rows == nil {
		return nil, errors.New("no rows returned")
	}

	for rows.Next() {
		var t db.Tender
		if err = rows.Scan(&t.Id, &t.Name, &t.Description, &t.Status, &t.ServiceType, &t.Version, &t.CreatedAt); err != nil {
			return
		}

		tenders = append(tenders, t)
	}

	return
}

func GetStatus(params GetTenderStatusParams) (status string, err error) {
	if db.MyDB == nil {
		err = errors.New("database is not connected")
		return
	}

	var rows *sql.Rows

	if rows, err = db.MyDB.Query(`SELECT status FROM tender
    WHERE id = $1`, params.TenderId); err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&status); err != nil {
			return
		}
	}

	return
}

func ChangeStatus(params ChangeTenderStatusParams) (tender db.Tender, err error) {
	if db.MyDB == nil {
		err = errors.New("database is not connected")
		return
	}

	if _, err = db.MyDB.Exec(`UPDATE tender
    SET status = $1
    WHERE id = $2 AND creator_username = $3`, params.Status, params.TenderId, params.Username); err != nil {
		return
	}

	var rows *sql.Rows
	if rows, err = db.MyDB.Query(`SELECT id, name, description, status, service_type, version, created_at FROM tender
                WHERE id = $1`,
		params.TenderId); err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&tender.Id, &tender.Name, &tender.Description, &tender.Status, &tender.ServiceType, &tender.Version, &tender.CreatedAt); err != nil {
			return
		}
	}
	return
}

func Update(params UpdateTenderParams) (tender db.Tender, err error) {
	if db.MyDB == nil {
		err = errors.New("database is not connected")
		return
	}

	if params.Name != "" {
		if params.Description != "" {
			if params.ServiceType != "" {
				if _, err = db.MyDB.Exec(`UPDATE tender
    				SET name = $1, description = $2, service_type = $3
    				WHERE id = $4 AND creator_username = $5`, params.Name, params.Description, params.ServiceType, params.TenderId, params.Username); err != nil {
					return
				}
			}
			if _, err = db.MyDB.Exec(`UPDATE tender
    				SET name = $1, description = $2
    				WHERE id = $3 AND creator_username = $4`, params.Name, params.Description, params.TenderId, params.Username); err != nil {
				return
			}
		}
		if _, err = db.MyDB.Exec(`UPDATE tender
    				SET name = $1
    				WHERE id = $2 AND creator_username = $3`, params.Name, params.TenderId, params.Username); err != nil {
			return
		}
	} else if params.Description != "" {
		if params.ServiceType != "" {
			if _, err = db.MyDB.Exec(`UPDATE tender
    				SET description = $1, service_type = $2
    				WHERE id = $3 AND creator_username = $4`, params.Description, params.ServiceType, params.TenderId, params.Username); err != nil {
				return
			}
		}
		if _, err = db.MyDB.Exec(`UPDATE tender
    				SET description = $1
    				WHERE id = $2 AND creator_username = $3`, params.Description, params.TenderId, params.Username); err != nil {
			return
		}
	} else if params.ServiceType != "" {
		if _, err = db.MyDB.Exec(`UPDATE tender
    				SET service_type = $1
    				WHERE id = $2 AND creator_username = $3`, params.ServiceType, params.TenderId, params.Username); err != nil {
			return
		}
	}

	var rows *sql.Rows
	if rows, err = db.MyDB.Query(`SELECT id, name, description, status, service_type, version, created_at FROM tender
                WHERE id = $1`,
		params.TenderId); err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&tender.Id, &tender.Name, &tender.Description, &tender.Status, &tender.ServiceType, &tender.Version, &tender.CreatedAt); err != nil {
			return
		}
	}
	return
}

//func Rollback(params RollbackTenderParams) (tender db.Tender, err error) {
//	if db.MyDB == nil {
//		err = errors.New("database is not connected")
//		return
//	}
//
//	var rows *sql.Rows
//	if rows, err = db.MyDB.Query(`SELECT id, name, description, status, service_type, version, created_at FROM tender_history
//               WHERE id = $1 AND version = $2 AND creator_username = $3`,
//		params.TenderId, params.Version, params.Username); err != nil {
//		return
//	}
//
//	defer rows.Close()
//
//	for rows.Next() {
//		if err = rows.Scan(&tender.Id, &tender.Name, &tender.Description, &tender.Status, &tender.ServiceType, &tender.Version, &tender.CreatedAt); err != nil {
//			return
//		}
//	}
//
//	if _, err = db.MyDB.Exec(`UPDATE tender
//   				SET name = $1, description = $2, service_type = $3, status = $4, organization_id = $5
//   				WHERE id = $6 AND creator_username = $7`, tender.Name, tender.Description, tender.ServiceType, tender.Status, tender.OrganizationId, tender.Id, tender.CreatorUsername); err != nil {
//		return
//	}
//
//	return
//}
