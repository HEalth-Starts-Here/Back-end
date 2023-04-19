package mlservicesusecase

import (
	"hesh/internal/pkg/domain"
	// recordusecase "hesh/internal/pkg/record/usecase"
	"hesh/internal/pkg/utils/filesaver"
	"path/filepath"
)

type MLServicesUsecase struct {
	mlservicesRepo domain.MLServicesRepository
}

func InitMLServicesUsc(mlsr domain.MLServicesRepository) domain.MLServicesUsecase {
	return &MLServicesUsecase{
		mlservicesRepo: mlsr,
	}
}

func (msu MLServicesUsecase) CreateMedicRecordDiarisations(userId uint64, recordId uint64, diarisationBeforeCompletingInfo domain.DiarisationBeforeCompletingInfo) (domain.DiarisationResponse, error) {
	// TODO uncomment check
	// var u domain.RecordUsecase
	// u = recordusecase.RecordUsecase{}
	// err := u.CheckMedicAndDiaryAndRecordExistAndMedicHaveAccess(medicId, recordId)
	// if err != nil {
	// 		return domain.DiarisationResponse{}, err
	// }
	alreadyUsed, err := msu.mlservicesRepo.GetAudioNames()

	audioNames := filesaver.GetUniqueFileNames(1, alreadyUsed)
	for i := 0; i < len(audioNames); i++ {
		diarisationBeforeCompletingInfo.Filename = audioNames[i] + filepath.Ext(diarisationBeforeCompletingInfo.Filename)
		audioNames[i] = diarisationBeforeCompletingInfo.Filename
	}

	RecordCreateResponse, err := msu.mlservicesRepo.CreateMedicRecordDiarisation(recordId, diarisationBeforeCompletingInfo)
	if err != nil {
		return domain.DiarisationResponse{}, err
	}
	return RecordCreateResponse, nil
}

func (msu MLServicesUsecase) SetDiarisationText (diarisationId uint64, diarisationText string) (error) {
	err := msu.mlservicesRepo.SetDiarisationText(diarisationId, diarisationText)
	if err != nil {
		return err
	}
	return nil
}
