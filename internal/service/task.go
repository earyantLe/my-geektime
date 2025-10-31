package service

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/url"
	"path"
	"sort"
	"strings"
	"time"

	"github.com/zkep/my-geektime/internal/global"
	"github.com/zkep/my-geektime/internal/model"
	"github.com/zkep/my-geektime/internal/types/geek"
	"github.com/zkep/my-geektime/internal/types/task"
	"go.uber.org/zap"
)

const (
	TASK_STATUS_PENDING  = 0x01
	TASK_STATUS_RUNNING  = 0x02
	TASK_STATUS_FINISHED = 0x03
	TASK_STATUS_ERROR    = 0x04
)

const (
	TASK_TYPE_PRODUCT = "product"
	TASK_TYPE_ARTICLE = "article"
)

var ALLStatus = []int{
	TASK_STATUS_PENDING,
	TASK_STATUS_RUNNING,
	TASK_STATUS_FINISHED,
	TASK_STATUS_ERROR,
}

var Replaces = map[string]string{
	"/":  "-",
	"|":  "-",
	"｜":  "-",
	":":  "：",
	`"`:  "“",
	"?":  "？",
	"&":  "+",
	"\t": "",
	"\b": "",
	" ":  "",
}

func VerifyFileName(name string) string {
	for k, v := range Replaces {
		name = strings.ReplaceAll(name, k, v)
	}
	return strings.TrimSpace(name)
}

type PlayMeta struct {
	Spec         []byte
	LocalSpec    []byte
	KeyPath      string
	Ciphertext   string
	CipherMethod string
	Parts        []Part
}

type PlayMetaRequest struct {
	DowloadURL string
	Dir        string
	Filename   string
	TaskId     string
	Ciphertext []byte
	Spec       []byte
}

type Part struct {
	Src   string
	Dest  string
	IsKey bool
}

func Download(ctx context.Context, x *model.Task, data geek.ArticleData) error {
	t0 := time.Now()
	var (
		source      string
		downloadURL string
		playURL     string
		err         error
	)
	fileName := VerifyFileName(data.Info.Title)
	dir := path.Join(x.TaskPid, VerifyFileName(data.Product.Title))
	if data.Info.IsVideo {
		if len(data.Info.Video.HlsMedias) == 0 && len(data.Info.VideoPreview.Medias) > 0 {
			data.Info.Video.HlsMedias = data.Info.VideoPreview.Medias
		}
		if len(data.Info.Video.HlsMedias) == 0 && len(data.Info.VideoPreview.Medias) == 0 {
			return fmt.Errorf("article info not found or no VIP product %s", x.OtherId)
		}
		sort.Slice(data.Info.Video.HlsMedias, func(i, j int) bool {
			return data.Info.Video.HlsMedias[i].Size > data.Info.Video.HlsMedias[j].Size
		})
		downloadURL = data.Info.Video.HlsMedias[0].URL
		playURL = downloadURL
	} else if data.Info.Audio.DownloadURL != "" {
		downloadURL = data.Info.Audio.DownloadURL
		playURL = data.Info.Audio.URL
	}

	if len(downloadURL) > 0 && len(playURL) > 0 {
		rewritePlayReq := PlayMetaRequest{
			DowloadURL: playURL,
			Dir:        dir,
			Filename:   fileName,
			TaskId:     x.TaskId,
			Spec:       x.RewriteHls,
		}
		if len(x.Ciphertext) > 0 {
			cipher, err1 := base64.StdEncoding.DecodeString(x.Ciphertext)
			if err1 != nil {
				global.LOG.Error("download rewritePlay", zap.Error(err1), zap.String("taskId", x.TaskId))
				return err1
			}
			rewritePlayReq.Ciphertext = cipher
		}
		meta, err1 := RewritePlay(ctx, rewritePlayReq)
		if err1 != nil {
			global.LOG.Error("download rewritePlay", zap.Error(err1), zap.String("taskId", x.TaskId))
			return err1
		}
		x.RewriteHls = meta.Spec
		x.Ciphertext = meta.Ciphertext
		if x.Bstatus > 0 {
			if data.Info.IsVideo {
				source, err = Video(ctx, dir, fileName, meta)
				if err != nil {
					global.LOG.Error("download video", zap.Error(err), zap.String("taskId", x.TaskId))
					return err
				}
			} else {
				up, _ := url.Parse(downloadURL)
				if strings.Contains(up.Path, up.Host) {
					downloadURL = strings.TrimPrefix(up.Path, "/")
				}
				source, err = Audio(ctx, x, downloadURL, dir, fileName)
				if err != nil {
					global.LOG.Error("download audio", zap.Error(err), zap.String("taskId", x.TaskId))
					return err
				}
			}
		}
	}

	if x.Bstatus > 0 {
		message := task.TaskMessage{}
		if len(source) > 0 {
			message.Object = global.Storage.GetKey(source, false)
		} else {
			message.Text = "not found download url"
		}
		x.Message, _ = json.Marshal(message)
	}

	global.LOG.Info("download end", zap.String("taskId", x.TaskId),
		zap.String("url", downloadURL), zap.Duration("cost", time.Since(t0)),
	)
	return nil
}
