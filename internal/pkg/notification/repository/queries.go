package notificationrepository

const (
	//  выбрать  дневники, у которых вариант = false (несколько раз в день), variant = true, (текущий день - startDate) mod frequency == 0
	getPatientsReminders = `
	SELECT d.patientId, d.title, d.variant, d.frequency
	FROM diaries d
	WHERE patientid is NOT NULL AND (variant = true OR (current_date - startdate::date) % frequency = 0)
	GROUP BY d.patientid, d.title, d.variant, d.frequency;
	` 
)
