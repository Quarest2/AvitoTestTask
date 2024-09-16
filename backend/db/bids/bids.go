package db

import (
	"backend/db"
	"database/sql"
	"errors"
	"time"
)

// Не знаю будет ли кто это читать, но я принял решение, потребовать еще и username от создателя для будущих запросов
func New(params NewBidParams) (bid db.Bid, err error) {
	if db.MyDB == nil {
		err = errors.New("database is not connected")
		return
	}

	timeNow := time.Now()

	var row *sql.Rows
	if row, err = db.MyDB.Query(`INSERT INTO bid (name, creator_username, description, tender_id, author_type, author_id, created_at) VALUES
                ($1, $2, $3, $4, $5, $6, $7) RETURNING id`, params.Name, params.CreatorUsername, params.Description, params.TenderId, params.AuthorType, params.AuthorId, timeNow); err != nil {
		return
	}

	var id string
	for row.Next() {
		if err = row.Scan(&id); err != nil {
			return
		}
	}

	bid = db.Bid{Id: id, Name: params.Name, Status: "Created", AuthorType: params.AuthorType, AuthorId: params.AuthorId, Version: 1, CreatedAt: timeNow}

	return
}

func GetMyBids(params GetMyBidsParams) (bids []db.Bid, err error) {
	if db.MyDB == nil {
		err = errors.New("database is not connected")
		return
	}

	var rows *sql.Rows
	if params.Limit > 0 && params.Offset > 0 {
		if rows, err = db.MyDB.Query(`SELECT id, name, creator_username, status, author_type, author_id, version, created_at FROM bid WHERE creator_username = $1
                OFFSET $2 ROWS
				FETCH NEXT $3 ROWS ONLY`,
			params.Username, params.Offset, params.Limit); err != nil {
			return
		}
	}
	if params.Limit > 0 && params.Offset <= 0 {
		if rows, err = db.MyDB.Query(`SELECT id, name, creator_username, status, author_type, author_id, version, created_at FROM bid WHERE creator_username = $1
				FETCH NEXT $2 ROWS ONLY`,
			params.Username, params.Limit); err != nil {
			return
		}
	}
	if params.Limit <= 0 && params.Offset > 0 {
		if rows, err = db.MyDB.Query(`SELECT id, name, creator_username, status, author_type, author_id, version, created_at FROM bid WHERE creator_username = $1
                OFFSET $2 ROWS`,
			params.Username, params.Offset); err != nil {
			return
		}
	}
	if params.Limit <= 0 && params.Offset <= 0 {
		if rows, err = db.MyDB.Query(`SELECT id, name, creator_username, status, author_type, author_id, version, created_at FROM bid WHERE creator_username = $1`, params.Username); err != nil {
			return
		}
	}

	defer rows.Close()
	if rows == nil {
		return nil, errors.New("no rows returned")
	}

	for rows.Next() {
		var b db.Bid
		if err = rows.Scan(&b.Id, &b.Name, &b.CreatorUsername, &b.Status, &b.AuthorType, &b.AuthorId, &b.Version, &b.CreatedAt); err != nil {
			return
		}

		bids = append(bids, b)
	}

	return
}

func GetListOfBidsByTender(params GetListOfBidsByTenderParams) (bids []db.Bid, err error) {
	if db.MyDB == nil {
		err = errors.New("database is not connected")
		return
	}

	var rows *sql.Rows
	if params.Limit > 0 && params.Offset > 0 {
		if rows, err = db.MyDB.Query(`SELECT id, name, status, author_type, author_id, version, created_at FROM bid WHERE tender_id = $1
                ORDER BY id
                OFFSET $2 ROWS
				FETCH NEXT $3 ROWS ONLY`,
			params.TenderId, params.Offset, params.Limit); err != nil {
			return
		}
	}
	if params.Limit > 0 && params.Offset <= 0 {
		if rows, err = db.MyDB.Query(`SELECT id, name, description, status, service_type, version, created_at FROM bid WHERE tender_id = $1
                ORDER BY id
				FETCH NEXT $2 ROWS ONLY`,
			params.TenderId, params.Limit); err != nil {
			return
		}
	}
	if params.Limit <= 0 && params.Offset > 0 {
		if rows, err = db.MyDB.Query(`SELECT id, name, description, status, service_type, version, created_at FROM bid WHERE tender_id = $1
                ORDER BY id
                OFFSET $2 ROWS`,
			params.TenderId, params.Offset); err != nil {
			return
		}
	}
	if params.Limit <= 0 && params.Offset <= 0 {
		if rows, err = db.MyDB.Query(`SELECT id, name, description, status, service_type, version, created_at FROM bid WHERE tender_id = $1 ORDER BY id`, params.TenderId); err != nil {
			return
		}
	}

	defer rows.Close()
	if rows == nil {
		return nil, errors.New("no rows returned")
	}

	for rows.Next() {
		var b db.Bid
		if err = rows.Scan(&b.Id, &b.Name, &b.Status, &b.AuthorType, &b.AuthorId, &b.Version, &b.CreatedAt); err != nil {
			return
		}

		bids = append(bids, b)
	}

	return
}

func GetStatus(params GetBidStatusParams) (status string, err error) {
	if db.MyDB == nil {
		err = errors.New("database is not connected")
		return
	}

	var rows *sql.Rows

	if rows, err = db.MyDB.Query(`SELECT status FROM bid
    WHERE id = $1`, params.BidId); err != nil {
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

func ChangeStatus(params ChangeBidStatusParams) (bid db.Bid, err error) {
	if db.MyDB == nil {
		err = errors.New("database is not connected")
		return
	}

	if _, err = db.MyDB.Exec(`UPDATE bid SET status = $1
    WHERE id = $2 AND creator_username = $3`, params.Status, params.BidId, params.Username); err != nil {
		return
	}

	var rows *sql.Rows
	if rows, err = db.MyDB.Query(`SELECT id, name, status, author_type, author_id, version, created_at FROM bid
                WHERE id = $1`,
		params.BidId); err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&bid.Id, &bid.Name, &bid.Status, &bid.AuthorType, &bid.AuthorId, &bid.Version, &bid.CreatedAt); err != nil {
			return
		}
	}
	return
}

func Update(params UpdateBidParams) (bid db.Bid, err error) {
	if db.MyDB == nil {
		err = errors.New("database is not connected")
		return
	}

	if params.Name != "" {
		if params.Description != "" {
			if _, err = db.MyDB.Exec(`UPDATE bid
    			SET name = $1, description = $2
    			WHERE id = $3 AND creator_username = $4`, params.Name, params.Description, params.BidId, params.Username); err != nil {
				return
			}
		}
		if _, err = db.MyDB.Exec(`UPDATE bid
    				SET name = $1
    				WHERE id = $2 AND creator_username = $3`, params.Name, params.BidId, params.Username); err != nil {
			return
		}
	} else if params.Description != "" {
		if _, err = db.MyDB.Exec(`UPDATE bid
    			SET description = $1
    			WHERE id = $2 AND creator_username = $3`, params.Description, params.BidId, params.Username); err != nil {
			return
		}
	}

	var rows *sql.Rows
	if rows, err = db.MyDB.Query(`SELECT id, name, status, author_type, author_id, version, created_at FROM bid
                WHERE id = $1`,
		params.BidId); err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&bid.Id, &bid.Name, &bid.Status, &bid.AuthorType, &bid.AuthorId, &bid.Version, &bid.CreatedAt); err != nil {
			return
		}
	}
	return
}

func SubmitDecision(params SubmitDecisionParams) (bid db.Bid, err error) {
	if db.MyDB == nil {
		err = errors.New("database is not connected")
		return
	}

	switch params.Decision {
	case Approved:
		if _, err = db.MyDB.Exec(`UPDATE bid
    			SET approve_count = approve_count + 1
    			WHERE id = $1 AND creator_username = $2`, params.BidId, params.Username); err != nil {
			return
		}
	case Rejected:
		if _, err = db.MyDB.Exec(`UPDATE bid
    			SET reject_count = reject_count + 1
    			WHERE id = $1 AND creator_username = $2`, params.BidId, params.Username); err != nil {
			return
		}
	}

	var rows *sql.Rows
	if rows, err = db.MyDB.Query(`SELECT id, name, status, author_type, author_id, version, created_at FROM bid
                WHERE id = $1`,
		params.BidId); err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&bid.Id, &bid.Name, &bid.Status, &bid.AuthorType, &bid.AuthorId, &bid.Version, &bid.CreatedAt); err != nil {
			return
		}
	}
	return
}

func Feedback(params FeedbackParams) (bid db.Bid, err error) {
	if db.MyDB == nil {
		err = errors.New("database is not connected")
		return
	}

	if _, err = db.MyDB.Query(`INSERT INTO bid_feedback (bid_id, comment, author_username, created_at) VALUES
                ($1, $2, $3, $4)`, params.BidId, params.BidFeedback, params.Username, time.Now()); err != nil {
		return
	}

	var rows *sql.Rows
	if rows, err = db.MyDB.Query(`SELECT id, name, status, author_type, author_id, version, created_at FROM bid
                WHERE id = $1`,
		params.BidId); err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&bid.Id, &bid.Name, &bid.Status, &bid.AuthorType, &bid.AuthorId, &bid.Version, &bid.CreatedAt); err != nil {
			return
		}
	}

	return
}

func GetReviews(params GetReviewsParams) (feedbacks []db.Feedback, err error) {
	if db.MyDB == nil {
		err = errors.New("database is not connected")
		return
	}

	var rows *sql.Rows
	if params.Limit > 0 && params.Offset > 0 {
		if rows, err = db.MyDB.Query(`SELECT id, description, created_at FROM bid_feedback WHERE author_username = $1
                OFFSET $2 ROWS
				FETCH NEXT $3 ROWS ONLY`,
			params.AuthorUsername, params.Offset, params.Limit); err != nil {
			return
		}
	}
	if params.Limit > 0 && params.Offset <= 0 {
		if rows, err = db.MyDB.Query(`SELECT id, description, created_at FROM bid_feedback WHERE author_username = $1
				FETCH NEXT $2 ROWS ONLY`,
			params.AuthorUsername, params.Limit); err != nil {
			return
		}
	}
	if params.Limit <= 0 && params.Offset > 0 {
		if rows, err = db.MyDB.Query(`SELECT id, description, created_at FROM bid_feedback WHERE author_username = $1
                OFFSET $2 ROWS`,
			params.AuthorUsername, params.Offset); err != nil {
			return
		}
	}
	if params.Limit <= 0 && params.Offset <= 0 {
		if rows, err = db.MyDB.Query(`SELECT id, description, created_at FROM bid_feedback WHERE author_username = $1`, params.AuthorUsername); err != nil {
			return
		}
	}

	defer rows.Close()
	if rows == nil {
		return nil, errors.New("no rows returned")
	}

	for rows.Next() {
		var f db.Feedback
		if err = rows.Scan(&f.Id, &f.Comment, &f.CreatedAt); err != nil {
			return
		}

		feedbacks = append(feedbacks, f)
	}

	return
}

//func Rollback(params RollbackBidParams) (bid db.Bid, err error) {
//	if db.MyDB == nil {
//		err = errors.New("database is not connected")
//		return
//	}
//
//	var rows *sql.Rows
//	if rows, err = db.MyDB.Query(`SELECT id, name, status, author_type, author_id, version, created_at FROM bid_history
//               WHERE id = $1 AND version = $2 AND creator_username = $3`,
//		params.BidId, params.Version, params.Username); err != nil {
//		return
//	}
//
//	defer rows.Close()
//
//	for rows.Next() {
//		if err = rows.Scan(&bid.Id, &bid.Name, &bid.Status, &bid.AuthorType, &bid.AuthorId, &bid.Version, &bid.CreatedAt); err != nil {
//			return
//		}
//	}
//
//	if _, err = db.MyDB.Exec(`UPDATE bid
//   				SET name = $1, name = $2, status = $3, author_type = $4, author_id = $5, version = version + 1, created_at = $6
// 				WHERE id = $7 AND creator_username = $8`, bid.Name, bid.Name, bid.Status, bid.AuthorType, bid.AuthorId, time.Now(), bid.Id, params.Username); err != nil {
//		return
//	}
//
//	return
//}
