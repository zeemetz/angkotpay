package services

import (
	"angkotpay/dto/request"
	"angkotpay/dto/response"
	"angkotpay/engine"
	"angkotpay/model"
)

func getID(source model.User, destination model.User, status string) (uint, uint, string) {
	return source.ID, destination.ID, status
}

func TransferService(request request.TransferRequest) response.Response {
	var source model.User
	source.Phone = request.From
	source.Pin = request.SourcePIN
	var destination model.User
	destination.Phone = request.To

	var transaction model.Transaction
	transaction.Amount = request.Amount
	transaction.Notes = request.Notes
	transaction.TransactionType = model.CASHIN
	transaction.Status = model.FAILED

	var res response.Response
	var error response.ErrorResponse

	tx := engine.GetORM().Begin()

	if tx.Set("gorm:query_option", "FOR UPDATE").Find(&source, source).Error != nil {
		transaction.Notes += " source of fund not found "
		res.Header = 400
		error.ErrorCode = 2
	} else if tx.Set("gorm:query_option", "FOR UPDATE").Find(&destination, destination).Error != nil {
		transaction.Notes += " destination of fund not found "
		res.Header = 400
		error.ErrorCode = 3
	} else if source.Balance = source.Balance - request.Amount; source.Balance < source.LowerLimit {
		transaction.Notes += " not enough balance "
		res.Header = 400
		error.ErrorCode = 4
	} else if destination.Balance = destination.Balance + request.Amount; destination.Balance > destination.UpperLimit {
		transaction.Notes += " balance exceed limit "
		res.Header = 400
		error.ErrorCode = 5
	} else if tx.Model(&source).Updates(source).Error != nil {
		transaction.Notes += " cannot save updates (source) to db "
		res.Header = 400
		error.ErrorCode = 6
	} else if tx.Model(&destination).Updates(destination).Error != nil {
		transaction.Notes += " cannot save updates (destination) to db "
		res.Header = 400
		error.ErrorCode = 6
	} else if transaction.From, transaction.To, transaction.Status = getID(source, destination, model.SUCCESS); engine.GetORM().Create(&transaction).Error != nil {
		res.Header = 500
		error.ErrorCode = 7
		// } else if transaction.From, transaction.To, transaction.Status = getID(source, destination, model.SUCCESS); tx.Create(&transaction).Error != nil {
		// 	res.Header = 500
		// 	error.ErrorCode = 7
	} else {
		res.Header = 200
		error.ErrorCode = 0
	}

	if res.Header != 200 {
		tx.Rollback()
		error.Message = transaction.Notes
		res.Body = error
	} else {
		tx.Commit()
		res.Body = transaction
	}

	return res
}
