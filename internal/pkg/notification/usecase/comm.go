package ratusecase

import (
	"fmt"
	"hesh/internal/pkg/domain"
	"hesh/internal/pkg/utils/log"
	"os"
	"strings"
	"time"

	"github.com/SevereCloud/vksdk/v2/api"
	// "github.com/SevereCloud/vksdk/v2/api/params"
)

const (
	textContinue = " ..."
)

type NotificationUsecase struct {
	notificationRepo domain.NotificationRepository
}

func InitNotificationUsc(nr domain.NotificationRepository) domain.NotificationUsecase {
	return &NotificationUsecase{
		notificationRepo: nr,
	}
}

func (u NotificationUsecase) PatientRecordRemind() () {
	for {
		// newRating, err := nu.notificationRepo.PatientRecordReminder(data.GetMovieId(), data.GetUserId(), int(data.GetRating()))
		// if err != nil {
		// 	return nil, err
		// }

		// vk := api.NewVK("e80e2119e80e2119e80e21198ceb1d081fee80ee80e21198c168a958ccfd793e077d5da")
		vk := api.NewVK(os.Getenv("vkToken"))


		// users, err := vk.UsersGet(api.Params{
		// 	"user_ids": 165523569,
		// 	"fields": "photo_50,verified,photo_id,bdate",
		// 	// "fields": "photo_50,verified",
		// })
		// if err != nil {
		// 	log.Error(err)
		// 	return
		// }
		// fmt.Printf("users: %v\n", users)


		// patientReminders := nu.notificationRepo.PatientRecordRemind()
		patientsReminders, err := u.notificationRepo.PatientRecordRemind()
		if err != nil {
			log.Error(err)
			return
		}

		// var userIds []int
		// userIds = append(userIds, 165523569)
		// userIds = append(userIds, 165523569)
		for i, _ := range(patientsReminders) {
			var message strings.Builder
			message.WriteString("Доброе утро! Сегодня вам по плану желательно сделать: ")
			for j := range(patientsReminders[i].NotificationInfo) {
				// usersMessages[i][j].Message
				diaryName := patientsReminders[i].NotificationInfo[j].DiaryName
				frequency := patientsReminders[i].NotificationInfo[j].Frequency
				if len(patientsReminders[i].NotificationInfo[j].DiaryName) + len(message.String()) > domain.MaxVkNotificationLength + len(textContinue) {
					message.WriteString(textContinue)
					break
				}
				message.WriteString(fmt.Sprintf("%d запись(и) в дневнике \"%s\"", frequency, diaryName))
				if j != len(patientsReminders[i].NotificationInfo) - 1 {
					message.WriteString(", ")
				} else {
					message.WriteString(".")
				}
			}
			println(patientsReminders[i].PatientId)
			notidications, err := vk.NotificationsSendMessage(api.Params{
				// "user_ids": patientsReminders[i].PatientId,
				"user_ids": 165523569,
				"message":  message.String(),
	
				// "message":  "Вам пришла электронная повестка! Узнать подробности можно в личном кабинете на портале ГосУслуг (gosuslugi.ru)",
				// "sending_mode":  5,
			})
			log.Info(message.String())
			if err != nil {
				log.Error(err)
			}
			fmt.Printf("notidications: %v\n", notidications)
		}
		println()
		println()
		println("Before sleep")
		time.Sleep(86400 * time.Second) // 1 day
		println("After sleep")
		// recursiveCall
		// u.PatientRecordRemind()
		// TODO check is shift is more than 10 minutes
	}

}
