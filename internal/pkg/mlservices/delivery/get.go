package mlservicesdelivery

import (
	"bytes"
	"hesh/internal/pkg/domain"
	"io"

	// "mime/multipart"

	// "path/filepath"
	// "eventool/internal/pkg/sessions"
	// "hesh/internal/pkg/utils/cast"
	// "hesh/internal/pkg/utils/cast"
	"hesh/internal/pkg/utils/config"
	"hesh/internal/pkg/utils/log"

	// "hesh/internal/pkg/utils/sanitizer"
	// "strconv"

	// "strings"
	"context"
	mlsgrpc "hesh/internal/pkg/mlservices/delivery/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"

	"github.com/mailru/easyjson"
	// "encoding/json"
	// "io"
	// "os"
)

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

	w.WriteHeader(http.StatusCreated)
	w.Write(out)
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

	w.WriteHeader(http.StatusCreated)
	w.Write(out)
	
}

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
