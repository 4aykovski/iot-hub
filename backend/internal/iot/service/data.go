package service

import (
	"context"
	"errors"
	"fmt"
	"net/smtp"
	"time"

	"github.com/4aykovski/iot-hub/backend/internal/iot/config"
	"github.com/4aykovski/iot-hub/backend/internal/iot/model"
	"github.com/4aykovski/iot-hub/backend/internal/iot/repo/repoerrs"
)

type DataRepository interface {
	GetDeviceData(ctx context.Context, id string) ([]model.Data, error)
	GetDeviceDataForPeriod(
		ctx context.Context,
		id string,
		start, end time.Time,
	) ([]model.Data, error)
}

type Data struct {
	dataRepo   DataRepository
	deviceRepo DeviceRepository

	smtpAuth smtp.Auth
	smtpHost string
	smtpPort int
	smtpMail string
}

func (da *Data) GetDeviceData(ctx context.Context, id string, interval int) ([]model.Data, error) {
	if interval <= 0 {
		return da.dataRepo.GetDeviceData(ctx, id)
	}

	dateTo := time.Now()
	dateFrom := dateTo.Add(time.Duration(interval) * -1 * time.Second)

	fmt.Println(dateFrom, dateTo)

	return da.dataRepo.GetDeviceDataForPeriod(ctx, id, dateFrom, dateTo)
}

type GetDataForPeriodDTO struct {
	ID   string
	From time.Time
	To   time.Time
}

func (da *Data) GetDataFromPeriod(
	ctx context.Context,
	dto GetDataForPeriodDTO,
) ([]model.Data, error) {
	data, err := da.dataRepo.GetDeviceDataForPeriod(ctx, dto.ID, dto.From, dto.To)
	if err != nil {
		if errors.Is(err, repoerrs.ErrNoData) {
			return []model.Data{}, ErrNoData
		}

		return nil, err
	}
	return data, nil
}

func (da *Data) SendEmail(
	ctx context.Context,
	id string,
	limit int,
	value string,
	timestamp string,
) error {
	device, err := da.deviceRepo.GetDevice(ctx, id)
	if err != nil {
		return err
	}

	message := "From: " + da.smtpMail + "\r\n" +
		"To: " + device.Email + "\r\n" +
		"Subject: Limit Exceeded\r\n" +
		"\r\n" +
		"Device: " + device.Name + "\r\n" +
		"Limit: " + fmt.Sprintf("%d", limit) + "\r\n" +
		"Value: " + value + "\r\n" +
		"Timestamp: " + timestamp + "\r\n"

	err = smtp.SendMail(
		fmt.Sprintf("%s:%d", da.smtpHost, da.smtpPort),
		da.smtpAuth,
		da.smtpMail,
		[]string{device.Email},
		[]byte(message),
	)

	return err
}

func (da *Data) SaveData(ctx context.Context, data model.Data) error {
	panic("not implemented") // TODO: Implement
}

func NewData(
	dataRepo DataRepository,
	smtpAuth smtp.Auth,
	mail config.Mail,
	device DeviceRepository,
) *Data {
	return &Data{
		dataRepo:   dataRepo,
		deviceRepo: device,
		smtpAuth:   smtpAuth,
		smtpHost:   mail.SmtpHost,
		smtpPort:   mail.SmtpPort,
		smtpMail:   mail.MailFrom,
	}
}
