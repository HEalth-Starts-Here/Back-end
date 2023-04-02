package diaryrepository

const (
	queryCreateDiary = `
	INSERT INTO
    diaries (medicId, patientId, creatingDate, title, complaints, anamnesis, objectively, diagnosis)
	VALUES
    (
		$1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7,
        $8
    )
	RETURNING id, medicId, patientId, creatingDate, title, complaints, anamnesis, objectively, diagnosis;
	`
	
	queryDeleteDiary = `
	DELETE FROM diaries
	WHERE id = $1;
	`

	queryDiaryList = `
	SELECT id, medicid, medics.name, patientid, patients.name, creatingdate, title, objectively
	FROM diaries
	JOIN patients ON diaries.patientid = patients.vkid
	JOIN medics ON diaries.medicid = medics.vkid
	ORDER BY creatingdate;
	`

	queryGetCertainDiaryMainInfo = `
	SELECT id, medicId, patientId, creatingDate, title, description
	FROM diaries
	WHERE id = $1;
	`

	// queryGetCertainDiaryRecords = `
	// SELECT (id, diaryId, creatingDate, title, description, area, dryness, edema, itching, pain, peeling, redness)
	// FROM records
	// WHERE diaryid = $1;
	// `
	queryGetCertainDiaryRecords = `
	SELECT id, diaryid, creatingdate, description, title, area, dryness, edema, itching, pain, peeling, redness
	FROM records
	WHERE diaryid = $1;
	`

	queryCreateRecord = `
	INSERT INTO
    records (diaryId, creatingDate, title, description, area, dryness, edema, itching, pain, peeling, redness)
	VALUES
    (
        $1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7,
        $8,
        $9,
        $10,
        $11
    )
	RETURNING id, diaryId, creatingDate, title, description, area, dryness, edema, itching, pain, peeling, redness;
	`
// 	queryCheckEvent = `
// 	SELECT count(*)
// 	FROM events 
// 	WHERE title = $1 and longitude = $2 and latitude = $3;
// 	`
// 	queryCreateEvent = `
// 	INSERT INTO
//     events (title, description, userId, longitude, latitude, maxmembersquantity, minmembersquantity, creatingdate, startdate, enddate, minage, maxage, price)
// 	VALUES
//     (
// 		$1,
//         $2,
//         $3, 
//         $4,
// 		$5,
//         $6,
//         $7,
//         $8,
//         $9,
//         $10,
//         $11,
//         $12,
//         $13
//     )
// 	RETURNING id, poster, title, rating, votesnum, description, userId, longitude, latitude, currentmembersquantity, maxmembersquantity, minmembersquantity, creatingdate, startdate, enddate, minage, maxage, price;
// 	`
	
// 	queryGetEventListFirstPart = `
// 	SELECT id, poster, title, rating, votesnum, description, userId, longitude, latitude, currentmembersquantity, maxmembersquantity, minmembersquantity, creatingdate, startdate, enddate, minage, maxage, price
// 	FROM events
// 	JOIN events_categories ON events.id = events_categories.eventId
// 	WHERE events_categories.category IN `
// 	queryGetEventListSecondPart = ` 
// 	ORDER BY events.id;
// 	`

// 	queryGetAllEventList = `
// 	SELECT id, poster, title, rating, votesnum, description, userId, longitude, latitude, currentmembersquantity, maxmembersquantity, minmembersquantity, creatingdate, startdate, enddate, minage, maxage, price
// 	FROM events
// 	ORDER BY events.id;
// 	`

// 	queryGetCategoryList = `
// 	SELECT name, imagePath
// 	FROM categories
// 	ORDER BY categories.name;
// 	`
// // 	queryCreateEventCategoryFirstPart = `
// // 	INSERT INTO
// //     events_categories (eventId, category)
// // 	VALUES`
// //     queryCreateEventCategorySecondPart = `
// // 	(
// // 		$`
// // 	queryCreateEventCategoryThirdPart = `,
// //         $`
// // 	queryCreateEventCategoryForthPart = `
// //     )`
// // 	queryCreateEventCategoryFifthPart = `
// // 	RETURNING eventId, category;`

// 	queryCreateEventCategory = `
// 	INSERT INTO
// 	events_categories (eventId, category)
// 	VALUES
// 	(
// 		$1,
// 		$2
// 	)
// 	RETURNING eventId, category;`

// 	queryGetCertainEvent = `
// 	SELECT id, poster, title, rating, votesnum, description, userId, longitude, latitude, currentmembersquantity, maxmembersquantity, minmembersquantity, creatingdate, startdate, enddate, minage, maxage, price
// 	FROM events
// 	WHERE id = $1;
// 	`

// 	querySignUpUserForEvent = `
// 	INSERT INTO 
// 	users_events (eventId, userId)
// 	VALUES
// 	(
//         $1,
//         $2
// 	);
// 	`

// 	queryCancelSignUpUserForEvent = `
// 	DELETE FROM 
// 	users_events
// 	WHERE eventid = $1 and userid = $2;
// 	`

// 	queryGetUserAge = `
// 	SELECT age 
// 	FROM users
// 	WHERE id = $1;
// 	`

// 	queryGetEventAges = `
// 	SELECT minage, maxage
// 	FROM events
// 	WHERE id = $1;
// 	`

queryGetImageList = `
SELECT id, recordid, name, area
FROM images;
`

queryCreateRecordImageListFirstPart = `
INSERT INTO
images (recordid, name, area)
VALUES
`

queryCreateRecordImageListSecondPart = `
RETURNING id, recordid, name, area;
`

queryUpdateDiary = `
UPDATE diaries
SET title = $1, description = $2
WHERE id = $3
RETURNING id, title, description;
`
)
