package mlservicesdelivery

import (
	"bytes"
	"hesh/internal/pkg/domain"
	"io"
	"mime/multipart"
	"path/filepath"
	"strconv"

	// "mime/multipart"

	// "path/filepath"
	// "eventool/internal/pkg/sessions"
	// "hesh/internal/pkg/utils/cast"
	// "hesh/internal/pkg/utils/cast"
	"hesh/internal/pkg/utils/config"
	"hesh/internal/pkg/utils/filesaver"
	"hesh/internal/pkg/utils/log"

	// "hesh/internal/pkg/utils/sanitizer"
	// "strconv"

	// "strings"
	"context"
	mlsgrpc "hesh/internal/pkg/mlservices/delivery/grpc"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/gorilla/mux"
	"github.com/mailru/easyjson"
	// "github.com/mailru/easyjson"
	// "encoding/json"
	// "io"
	// "os"
)

func validAudioExtenstions(files []*multipart.FileHeader) bool {
	availableExtensions := map[string]struct{}{".mp3": {}}
	for i := range files {
		extension := filepath.Ext(files[i].Filename)
		_, is := availableExtensions[extension]
		if !is {
			return false
		}

	}
	return true
}

func readMultipartDataAudio(r *http.Request) ([]string, int, error) {

	err := r.ParseMultipartForm(1 << 28) // maxMemory 256MB
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	formdata := r.MultipartForm
	//get the *fileheaders
	files := formdata.File["audio"] // grab the filenames
	audioInfo := make([]string, 0)
	for i, _ := range files { // loop through the files one by one
		file, err := files[i].Open()
		defer file.Close()
		if err != nil {
			// TODO: add mapping from error to http code
			return nil, http.StatusInternalServerError, domain.Err.ErrObj.InternalServer
		}

		if !validAudioExtenstions(files) {
			return nil, http.StatusBadRequest, domain.Err.ErrObj.BadFileExtension
		}
		// TODO: parse tags
		// tags := make([]string, 0)
		// for j := range (r.Form["tags"])[i]{
		// 	fmt.Sprintf("%v", (r.Form["title"])[0])
		// 	tags = append(tags, fmt.Sprintf("%v", (r.Form["tags"])[i][j]))
		// }
		if err != nil {
			return nil, http.StatusBadRequest, domain.Err.ErrObj.BadInput
		}
		audioInfo = append(audioInfo, filesaver.ExtractName(files[i].Filename))
	}
	return audioInfo, http.StatusCreated, nil
}

// func GetAudioSummarization(fileHeader []*multipart.FileHeader) (*mlsgrpc.DiarisationResponse, error) {
// 	if len(fileHeader) == 0 {
// 		return nil, nil
// 	}
// 	file, err := fileHeader[0].Open()
// 	defer file.Close()
// 	if err != nil {
// 		log.Error(err)
// 		return nil, err
// 	}

// 	buf := bytes.NewBuffer(nil)
// 	if _, err := io.Copy(buf, file); err != nil {
// 		log.Error(err)
// 		return nil, err

// 	}

// 	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
// 	if err != nil {
// 		log.Error(err)
// 		return nil, err

// 	}
// 	defer conn.Close()

// 	client := mlsgrpc.NewDiarisationClient(conn)
// 	response, err := client.TranscribeAudio(context.Background(), &mlsgrpc.DiarisationRequest{
// 		Audio: buf.Bytes(),
// 	})
// 	if err != nil {
// 		log.Error(err)
// 		return nil, err

// 	}
// 	return response, nil
// }

func (handler *MLServicesHandler) DetermineArea(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	err := r.ParseMultipartForm(32 << 20) // maxMemory 32MB
	if err != nil {
		log.Error(err)
		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
		return
	}
	formdata := r.MultipartForm
	fileHeader := formdata.File["image"]
	println(config.DevConfigStore.LoadedFilesPath + (fileHeader[0]).Filename)

	//TODO: check extension
	file, err := fileHeader[0].Open()
	defer file.Close()
	if err != nil {
		log.Error(err)
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		log.Error(err)
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error(err)
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	// client := mlsgrpc.AffectedAreaClient()
	client := mlsgrpc.NewAffectedAreaClient(conn)
	response, err := client.CalculateArea(context.Background(), &mlsgrpc.AffectedAreaRequest{
		Image: buf.Bytes(),
	})
	if err != nil {
		log.Error(err)
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// log.Printf("book list: %v", bookList)

	es := domain.DetermineAreaResponse{
		Area: response.Area,
	}

	out, err := easyjson.Marshal(es)
	if err != nil {
		log.Error(err)
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write(out)
	if err != nil {
		log.Error(err)
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (handler *MLServicesHandler) ImageQualityAssesment(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	err := r.ParseMultipartForm(32 << 20) // maxMemory 32MB
	if err != nil {
		log.Error(err)
		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
		return
	}
	formdata := r.MultipartForm
	fileHeader := formdata.File["image"]
	println(config.DevConfigStore.LoadedFilesPath + (fileHeader[0]).Filename)

	//TODO: check extension
	file, err := fileHeader[0].Open()
	defer file.Close()
	if err != nil {
		log.Error(err)
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		log.Error(err)
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error(err)
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	client := mlsgrpc.NewIQAClient(conn)
	response, err := client.CalculateQuality(context.Background(), &mlsgrpc.IQARequest{
		Image: buf.Bytes(),
	})
	if err != nil {
		log.Error(err)
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// log.Printf("book list: %v", bookList)

	es := domain.ImageQualityAssesment{
		Assesment: response.Quality,
	}

	out, err := easyjson.Marshal(es)
	if err != nil {
		log.Error(err)
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write(out)
	if err != nil {
		log.Error(err)
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (handler *MLServicesHandler) DiarisationRequestToMS(diarisationId uint64, file multipart.File) {

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		log.Error(err)
	}

	conn, err := grpc.Dial("127.0.0.1:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error(err)
	}
	defer conn.Close()

	// client := mlsgrpc.AffectedAreaClient()
	client := mlsgrpc.NewDiarisationClient(conn)
	diarisationTextMLSResponse, err := client.TranscribeAudio(context.Background(), &mlsgrpc.DiarisationRequest{
		Audio: buf.Bytes(),
	})
	if err != nil {
		log.Error(err)
	}
	diarisationText := diarisationTextMLSResponse.Text
	err = handler.MLServicesUsecase.SetDiarisationText(diarisationId, diarisationText)
	if err != nil {
		log.Error(err)

	}
}

func (handler *MLServicesHandler) Diarisation(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	queryParameter := r.URL.Query().Get("vk_user_id")
	userId, err := strconv.ParseUint(queryParameter, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	params := mux.Vars(r)
	recordId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		log.Error(err)
		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
		return
	}

	audioNames, httpStatus, err := readMultipartDataAudio(r)
	if err != nil {
		log.Error(err)
		http.Error(w, err.Error(), httpStatus)
		w.WriteHeader(httpStatus)
		return
	}

	formdata := r.MultipartForm
	fileHeader := formdata.File["audio"]
	println(config.DevConfigStore.LoadedFilesPath + (fileHeader[0]).Filename)

	file, err := fileHeader[0].Open()
	if err != nil {
		log.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	creatingResponse, err := handler.MLServicesUsecase.CreateMedicRecordDiarisations(userId, recordId, domain.DiarisationBeforeCompletingInfo{
		Filename: audioNames[0],
	})
	if err != nil {
		log.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// es := domain.DiarisationResponse{
	// 	Diarisation: response.Text,
	// }
	go handler.DiarisationRequestToMS(creatingResponse.Id, file)

	audioNamesToSave := make([]string, 0)
	audioNamesToSave = append(audioNamesToSave, creatingResponse.DiarisationInfo.Filename)
	filesaver.SaveMultipartDataFiles(audioNamesToSave, r.MultipartForm.File["audio"])

	out, err := easyjson.Marshal(creatingResponse)
	if err != nil {
		log.Error(err)
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write(out)
	if err != nil {
		log.Error(err)
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// func (handler *MLServicesHandler) Diarisation (file multipart.File) (*mlsgrpc.DiarisationResponse, error) {
// 	// defer r.Body.Close()

// 	// audioNames, httpStatus, err := readMultipartDataAudio(r)
// 	// if err != nil {
// 	// 	log.Error(err)
// 	// 	http.Error(w, err.Error(), httpStatus)
// 	// 	w.WriteHeader(httpStatus)
// 	// 	return
// 	// }
// 	// filesaver.SaveMultipartDataFiles(audioNames, r.MultipartForm.File["audio"])

// 	// formdata := r.MultipartForm
// 	// fileHeader := formdata.File["audio"]
// 	// println(config.DevConfigStore.LoadedFilesPath + (fileHeader[0]).Filename)

// 	// file, err := fileHeader[0].Open()
// 	// if err != nil {
// 	// 	log.Error(err)
// 	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
// 	// 	w.WriteHeader(http.StatusInternalServerError)
// 	// 	return
// 	// }

// 	buf := bytes.NewBuffer(nil)
// 	if _, err := io.Copy(buf, file); err != nil {
// 		return nil, err
// 	}

// 	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer conn.Close()

// 	// client := mlsgrpc.AffectedAreaClient()
// 	client := mlsgrpc.NewDiarisationClient(conn)
// 	response, err := client.TranscribeAudio(context.Background(), &mlsgrpc.DiarisationRequest{
// 		Audio: buf.Bytes(),
// 	})
// 	if err != nil {
// 		return nil, err
// 	}
// 	return response, nil
// 	// log.Printf("book list: %v", bookList)

// 	// es := domain.DiarisationResponse{
// 	// 	Diarisation: response.Text,
// 	// }

// 	// out, err := easyjson.Marshal(es)
// 	// if err != nil {
// 	// 	log.Error(err)
// 	// 	http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
// 	// 	w.WriteHeader(http.StatusInternalServerError)
// 	// 	return
// 	// }

// 	// w.WriteHeader(http.StatusCreated)
// 	// w.Write(out)
// }

// func (handler *MLServicesHandler) AudioSummarization(w http.ResponseWriter, r *http.Request) {
// 	defer r.Body.Close()

// 	err := r.ParseMultipartForm(1 << 28) // maxMemory 256MB
// 	if err != nil {
// 		log.Error(err)
// 		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
// 		return
// 	}
// 	formdata := r.MultipartForm
//  	fileHeader := formdata.File["audio"]
// 	println(config.DevConfigStore.LoadedFilesPath + (fileHeader[0]).Filename)

// 	//TODO: check extension
// 	file, err := fileHeader[0].Open()
// 	defer file.Close()
// 	if err != nil {
// 		log.Error(err)
// 		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	buf := bytes.NewBuffer(nil)
// 	if _, err := io.Copy(buf, file); err != nil {
// 		log.Error(err)
// 		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
// 	if err != nil {
// 		log.Error(err)
// 		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	defer conn.Close()

// 	client := mlsgrpc.NewIQAClient(conn)
// 	response, err := client.CalculateQuality(context.Background(), &mlsgrpc.IQARequest{
// 		Image: buf.Bytes(),
// 	})
// 	if err != nil {
// 		log.Error(err)
// 		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	// log.Printf("book list: %v", bookList)

// 	es := domain.ImageQualityAssesment{
// 		Assesment: response.Quality,
// 	}

// 	out, err := easyjson.Marshal(es)
// 	if err != nil {
// 		log.Error(err)
// 		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusCreated)
// 	w.Write(out)

// }
