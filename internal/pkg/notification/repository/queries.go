package notificationrepository

const (
	//  выбрать  дневники, у которых вариант = false (несколько раз в день), variant = true, (текущий день - startDate) mod frequency == 0
	getPatientsReminders = `
	SELECT convert_from(decrypt(d.patientId::text::bytea,'secret','aes'),'SQL_ASCII')::bigint, d.title, d.variant, d.frequency
	FROM diaries d
	WHERE patientid is NOT NULL AND (variant = true OR (current_date - startdate::date) % frequency = 0)
	GROUP BY convert_from(decrypt(d.patientid::text::bytea,'secret','aes'),'SQL_ASCII')::bigint, d.title, d.variant, d.frequency;
	` 
)
