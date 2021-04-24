package handler

import (
	"course/data"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

func UpLoad(ctx *gin.Context)  {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusOK, "FormFile failed")
		return
	}
	fileName := file.Filename
	upLoadFile, err := file.Open()
	if err != nil {
		ctx.JSON(http.StatusOK, "open failed")
		return
	}
	defer upLoadFile.Close()

	localFile, err := os.Create(fmt.Sprintf("%s/%s", data.DataPath, fileName))
	if err != nil {
		ctx.JSON(http.StatusOK, "create failed")
		return
	}
	_, err = io.Copy(localFile, upLoadFile)
	if err != nil {
		ctx.JSON(http.StatusOK, "copy failed")
		return
	}
	ctx.JSON(http.StatusOK, "upload success")
}

func DownLoadData(ctx *gin.Context)  {
	fileName := ctx.Query("file_name")
	ctx.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))//fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
	ctx.Writer.Header().Add("Content-Type", "application/octet-stream")
	ctx.File(fmt.Sprintf("%s/%s", data.DataPath, fileName))
}
