package v2

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"

	htmltomarkdown "github.com/JohannesKaufmann/html-to-markdown/v2"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"github.com/golang-jwt/jwt/v4"
	"github.com/zkep/my-geektime/internal/global"
	"github.com/zkep/my-geektime/internal/model"
	"github.com/zkep/my-geektime/internal/service"
	"github.com/zkep/my-geektime/internal/types/geek"
	"github.com/zkep/my-geektime/internal/types/task"
	"github.com/zkep/my-geektime/libs/zhttp"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Task struct{}

func NewTask() *Task {
	return &Task{}
}

func (t *Task) List(c *gin.Context) {
	var req task.TaskListRequest
	if err := c.ShouldBind(&req); err != nil {
		global.FAIL(c, "fail.msg", err.Error())
		return
	}
	if req.PerPage <= 0 || (req.PerPage > 200) {
		req.PerPage = 10
	}
	if req.Page <= 0 {
		req.Page = 1
	}
	ret := task.TaskListResponse{
		Rows: make([]task.Task, 0, 10),
	}
	var ls []*model.Task
	tx := global.DB.Model(&model.Task{})
	if req.Xstatus > 0 {
		tx = tx.Where("status = ?", req.Xstatus)
	}
	if req.ProductForm > 0 {
		tx = tx.Where("other_form = ?", req.ProductForm)
	}
	if req.ProductType > 0 {
		tx = tx.Where("other_type = ?", req.ProductType)
	}
	if req.Tag > 0 {
		tx = tx.Where("other_tag = ?", req.Tag)
	}
	if req.Direction > 0 {
		tx = tx.Where("other_group = ?", req.Direction)
	}
	if req.Keywords != "" {
		tx = tx.Where("task_name LIKE ?", "%"+req.Keywords+"%")
	}
	tx = tx.Where("task_pid = ?", req.TaskPid)
	tx = tx.Where("deleted_at = ?", 0)
	if req.TaskPid != "" {
		tx = tx.Order("id ASC")
	} else {
		tx = tx.Order("id DESC")
	}
	if err := tx.Count(&ret.Count).
		Offset((req.Page - 1) * req.PerPage).
		Limit(req.PerPage).
		Find(&ls).Error; err != nil {
		global.FAIL(c, "fail.msg", err.Error())
		return
	}

	for _, l := range ls {
		var statistics task.TaskStatistics
		if len(l.Statistics) > 0 {
			_ = json.Unmarshal(l.Statistics, &statistics)
		}
		row := task.Task{
			TaskId:     l.TaskId,
			TaskPid:    l.TaskPid,
			OtherId:    l.OtherId,
			TaskName:   l.TaskName,
			Status:     l.Status,
			Statistics: statistics,
			TaskType:   l.TaskType,
			Cover:      l.Cover,
			CreatedAt:  l.CreatedAt,
			UpdatedAt:  l.UpdatedAt,
		}
		switch l.TaskType {
		case service.TASK_TYPE_PRODUCT:
			var product geek.ProductBase
			if len(l.Raw) > 0 {
				_ = json.Unmarshal(l.Raw, &product)
			}
			row.Author = product.Author
			row.Share = product.Share
			row.Article = product.Article
			row.Subtitle = product.Subtitle
			row.IntroHTML = product.IntroHTML
			row.IsVideo = product.IsVideo
			row.IsAudio = product.IsAudio
			row.Sale = product.Price.Sale
			row.SaleType = product.Price.SaleType
			row.IsAudio = product.IsAudio
			var taskMessage task.TaskMessage
			if len(l.Message) > 0 {
				_ = json.Unmarshal(l.Message, &taskMessage)
				if len(taskMessage.Object) > 0 {
					row.Dir = global.Storage.GetUrl(taskMessage.Object)
					row.Dir = fmt.Sprintf("%s/", row.Dir)
				}
				if len(taskMessage.Doc) > 0 {
					row.Doc = global.Storage.GetUrl(taskMessage.Doc)
				}
			}
		case service.TASK_TYPE_ARTICLE:
			var articelInfo geek.ArticleData
			if len(l.Raw) > 0 {
				_ = json.Unmarshal(l.Raw, &articelInfo)
			}
			row.Author = articelInfo.Info.Author
			row.Subtitle = articelInfo.Info.Subtitle
			row.IntroHTML = articelInfo.Info.Summary
			row.IsVideo = articelInfo.Info.IsVideo
			var taskMessage task.TaskMessage
			if len(l.Message) > 0 {
				_ = json.Unmarshal(l.Message, &taskMessage)
				if len(taskMessage.Object) > 0 {
					row.Object = global.Storage.GetUrl(taskMessage.Object)
				}
			}
		}

		row.Cover = service.URLProxyReplace(row.Cover)
		row.Author.Avatar = service.URLProxyReplace(row.Author.Avatar)
		row.Share.Cover = service.URLProxyReplace(row.Share.Cover)
		if len(row.IntroHTML) > 0 {
			if introHTML, err1 := service.HtmlURLProxyReplace(row.IntroHTML); err1 == nil {
				row.IntroHTML = introHTML
			}
			if markdown, err := htmltomarkdown.ConvertString(row.IntroHTML); err == nil {
				row.IntroHTML = markdown
			}
		}
		ret.Rows = append(ret.Rows, row)
	}
	global.OK(c, ret)
}

func (t *Task) Info(c *gin.Context) {
	var req task.TaskInfoRequest
	if err := c.ShouldBind(&req); err != nil {
		global.FAIL(c, "fail.msg", err.Error())
		return
	}
	var l model.Task
	if err := global.DB.Model(&model.Task{}).
		Where(&model.Task{TaskId: req.Id}).First(&l).Error; err != nil {
		global.FAIL(c, "fail.msg", err.Error())
		return
	}
	var statistics task.TaskStatistics
	if len(l.Statistics) > 0 {
		_ = json.Unmarshal(l.Statistics, &statistics)
	}

	var articleData geek.ArticleData
	if len(l.Raw) > 0 {
		if err := json.Unmarshal(l.Raw, &articleData); err != nil {
			global.FAIL(c, "fail.msg", err.Error())
			return
		}
	}

	var taskMessage task.TaskMessage
	if len(l.Message) > 0 {
		_ = json.Unmarshal(l.Message, &taskMessage)
		if len(taskMessage.Object) > 0 {
			taskMessage.Object = global.Storage.GetUrl(taskMessage.Object)
		}
	}

	resp := task.TaskInfoResponse{
		Task: task.Task{
			TaskId:     l.TaskId,
			TaskPid:    l.TaskPid,
			OtherId:    l.OtherId,
			TaskName:   l.TaskName,
			Status:     l.Status,
			Cover:      l.Cover,
			Statistics: statistics,
			TaskType:   l.TaskType,
			CreatedAt:  l.CreatedAt,
			UpdatedAt:  l.UpdatedAt,
		},
		Article: articleData.Info,
		Message: taskMessage,
	}
	resp.Task.Cover = service.URLProxyReplace(resp.Task.Cover)
	resp.Article.Cover.Square = service.URLProxyReplace(resp.Article.Cover.Square)
	resp.Article.Author.Avatar = service.URLProxyReplace(resp.Article.Author.Avatar)
	if len(resp.Article.Cshort) > len(resp.Article.Content) {
		resp.Article.Content = resp.Article.Cshort
	}
	resp.Article.Cshort = ""
	if introHTML, err1 := service.HtmlURLProxyReplace(resp.Article.Content); err1 == nil {
		resp.Article.Content = introHTML
	}
	if markdown, err := htmltomarkdown.ConvertString(resp.Article.Content); err == nil {
		resp.Article.Content = markdown
	}
	if len(l.Ciphertext) > 0 || len(l.RewriteHls) > 0 {
		resp.PalyURL = fmt.Sprintf("%s/v2/task/play.m3u8?id=%s", global.CONF.Storage.Host, l.TaskId)
	}
	global.OK(c, resp)
}

func (t *Task) Retry(c *gin.Context) {
	var req task.RetryRequest
	if err := c.ShouldBind(&req); err != nil {
		global.FAIL(c, "fail.msg", err.Error())
		return
	}
	var info model.Task
	if err := global.DB.Model(&model.Task{}).Where("task_id", req.Pid).First(&info).Error; err != nil {
		global.FAIL(c, "fail.msg", err.Error())
		return
	}
	var statistics task.TaskStatistics
	if err := json.Unmarshal(info.Statistics, &statistics); err != nil {
		global.FAIL(c, "fail.msg", err.Error())
		return
	}
	if !req.Retry && len(req.Ids) == 0 && info.Bstatus > 0 {
		finishCount, ok := statistics.Items[service.TASK_STATUS_FINISHED]
		if !ok {
			global.FAIL(c, "tasks.illegal_operation", 0, statistics.Count)
			return
		}
		if statistics.Count != finishCount {
			global.FAIL(c, "tasks.illegal_operation", finishCount, statistics.Count)
			return
		}
		global.OkWithMsg(c, nil, "tasks.allow_operation", "OK", finishCount, statistics.Count)
		return
	}
	newStatistics := task.TaskStatistics{
		Count: statistics.Count,
		Items: map[int]int{
			service.TASK_STATUS_PENDING:  statistics.Count,
			service.TASK_STATUS_RUNNING:  0,
			service.TASK_STATUS_FINISHED: 0,
			service.TASK_STATUS_ERROR:    0,
		},
	}
	statisticsRaw, _ := json.Marshal(newStatistics)
	err := global.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&model.Task{}).
			Where("task_id", req.Pid).
			UpdateColumn("status", service.TASK_STATUS_PENDING).
			UpdateColumn("bstatus", service.TASK_STATUS_PENDING).
			UpdateColumn("statistics", statisticsRaw).Error; err != nil {
			return err
		}
		if len(req.Ids) == 0 {
			if err := tx.Model(&model.Task{}).
				Where("task_pid", req.Pid).
				UpdateColumn("status", service.TASK_STATUS_PENDING).
				UpdateColumn("bstatus", service.TASK_STATUS_PENDING).Error; err != nil {
				return err
			}
		} else {
			for _, idx := range strings.Split(req.Ids, ",") {
				if err := tx.Model(&model.Task{}).
					Where("task_id", idx).
					UpdateColumn("status", service.TASK_STATUS_PENDING).
					UpdateColumn("bstatus", service.TASK_STATUS_PENDING).Error; err != nil {
					return err
				}
			}
		}
		return nil
	})
	if err != nil {
		global.FAIL(c, "fail.msg", err.Error())
		return
	}
	global.OK(c, nil)
}

func (t *Task) Delete(c *gin.Context) {
	var req task.DeleteRequest
	if err := c.ShouldBind(&req); err != nil {
		global.FAIL(c, "fail.msg", err.Error())
		return
	}
	err := global.DB.Transaction(func(tx *gorm.DB) error {
		if len(req.Pid) > 0 {
			if err := tx.Model(&model.Task{}).
				Where("task_id", req.Pid).
				Updates(map[string]any{"deleted_at": time.Now().Unix()}).Error; err != nil {
				return err
			}
			if len(req.Ids) == 0 {
				if err := tx.Model(&model.Task{}).
					Where("task_pid", req.Pid).
					Updates(map[string]any{"deleted_at": time.Now().Unix()}).Error; err != nil {
					return err
				}
			}
		}
		for _, idx := range strings.Split(req.Ids, ",") {
			if err := tx.Model(&model.Task{}).
				Where("task_id", idx).
				Updates(map[string]any{"deleted_at": time.Now().Unix()}).Error; err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		global.FAIL(c, "fail.msg", err.Error())
		return
	}
	global.OK(c, nil)
}

func (t *Task) Download(c *gin.Context) {
	var req task.TaskDownloadRequest
	if err := c.ShouldBind(&req); err != nil {
		global.FAIL(c, "fail.msg", err.Error())
		return
	}
	var l model.Task
	if err := global.DB.Model(&model.Task{}).
		Where(&model.Task{TaskId: req.Id}).First(&l).Error; err != nil {
		global.FAIL(c, "fail.msg", err.Error())
		return
	}
	var articleData geek.ArticleData
	if err := json.Unmarshal(l.Raw, &articleData); err != nil {
		global.FAIL(c, "fail.msg", err.Error())
		return
	}
	var taskMessage task.TaskMessage
	if len(l.Message) > 0 {
		if err := json.Unmarshal(l.Message, &taskMessage); err != nil {
			global.FAIL(c, "fail.msg", err.Error())
			return
		}
	}
	baseName := service.VerifyFileName(articleData.Info.Title)
	switch req.Type {
	case "markdown":
		if len(articleData.Info.Cshort) > len(articleData.Info.Content) {
			articleData.Info.Content = articleData.Info.Cshort
		}
		markdown, err := htmltomarkdown.ConvertString(articleData.Info.Content)
		if err != nil {
			global.FAIL(c, "fail.msg", err.Error())
			return
		}
		fileName := baseName + ".md"
		c.Header("Content-Type", "application/octet-stream")
		c.Header("Content-Disposition", "attachment; filename="+url.QueryEscape(fileName))
		c.Header("Content-Transfer-Encoding", "binary")
		c.Data(200, "application/octet-stream", []byte(markdown))
	case "audio", "video":
		fileName := baseName + ".mp4"
		if req.Type == "audio" {
			fileName = baseName + ".mp3"
		}
		if len(req.Url) > 0 {
			err := zhttp.NewRequest().Client(global.HttpClient).
				Before(func(r *http.Request) {
					r.Header.Set("Accept", "application/json, text/plain, */*")
					r.Header.Set("Content-Type", "application/json")
					r.Header.Set("Sec-Ch-Ua", `"Google Chrome";v="119", "Chromium";v="119", "Not?A_Brand";v="24"`)
					r.Header.Set("User-Agent", zhttp.RandomUserAgent())
					r.Header.Set("Referer", req.Url)
					r.Header.Set("Origin", "https://time.geekbang.org")
				}).
				After(func(r *http.Response) error {
					if zhttp.IsHTTPSuccessStatus(r.StatusCode) {
						c.Header("Content-Type", "application/octet-stream")
						c.Header("Content-Disposition", "attachment; filename="+url.QueryEscape(fileName))
						c.Header("Content-Transfer-Encoding", "binary")
						c.Render(http.StatusOK, render.Reader{
							ContentLength: -1,
							ContentType:   "application/octet-stream",
							Reader:        r.Body,
						})
						return nil
					}
					if zhttp.IsHTTPStatusSleep(r.StatusCode) {
						time.Sleep(time.Second * 10)
					}
					if zhttp.IsHTTPStatusRetryable(r.StatusCode) {
						return errors.New("http status: " + r.Status)
					}
					return zhttp.BreakRetryError(errors.New("http status: " + r.Status))
				}).
				DoWithRetry(c, http.MethodGet, req.Url, nil)
			if err != nil {
				global.FAIL(c, "fail.msg", err.Error())
				return
			}
			return
		}
		c.Header("Content-Type", "application/octet-stream")
		c.Header("Content-Disposition", "attachment; filename="+url.QueryEscape(fileName))
		c.Header("Content-Transfer-Encoding", "binary")
		object := global.Storage.GetKey(taskMessage.Object, true)
		c.File(object)
	}
}

func (t *Task) Kms(c *gin.Context) {
	var req task.TaskKmsRequest
	if err := c.ShouldBind(&req); err != nil {
		global.FAIL(c, "fail.msg", err.Error())
		return
	}
	token, err := global.JWT.ParseToken(req.Ciphertext)
	if err != nil {
		global.FAIL(c, "fail.msg", err.Error())
		return
	}
	mapClaims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		global.FAIL(c, "fail.msg", errors.New("invalid token claims"))
		return
	}
	taskId, ok := mapClaims["task_id"]
	if !ok {
		global.FAIL(c, "fail.msg", errors.New("invalid vod"))
		return
	}
	tid, ok := taskId.(string)
	if !ok {
		global.FAIL(c, "fail.msg", errors.New("invalid vod"))
		return
	}
	var l model.Task
	if err = global.DB.Model(&model.Task{}).
		Where(&model.Task{TaskId: tid}).First(&l).Error; err != nil {
		global.FAIL(c, "fail.msg", err.Error())
		return
	}
	ciphertext, err := base64.StdEncoding.DecodeString(l.Ciphertext)
	if err != nil {
		global.FAIL(c, "fail.msg", err.Error())
		return
	}
	c.Header("Content-Type", "application/octet-stream")
	c.Data(200, "application/octet-stream", ciphertext)
}

func (t *Task) Play(c *gin.Context) {
	var req task.TaskPlayRequest
	if err := c.ShouldBind(&req); err != nil {
		global.FAIL(c, "fail.msg", err.Error())
		return
	}
	var l model.Task
	if err := global.DB.Model(&model.Task{}).
		Where(&model.Task{TaskId: req.Id}).First(&l).Error; err != nil {
		global.FAIL(c, "fail.msg", err.Error())
		return
	}
	l.RewriteHls = regexp.MustCompile("{host}").ReplaceAll(l.RewriteHls, []byte(global.CONF.Storage.Host))
	var (
		buff bytes.Buffer
	)
	bio := bufio.NewReader(bytes.NewReader(l.RewriteHls))
	for {
		line, _, err1 := bio.ReadLine()
		if err1 != nil {
			break
		}
		ln := string(line)
		if strings.HasPrefix(ln, "#EXT-X-KEY:") {
			sps := strings.Split(ln, `"`)
			token, _, er := global.JWT.TokenGenerator(func(claims jwt.MapClaims) {
				claims["task_id"] = l.TaskId
				now := time.Now().UTC()
				expire := now.Add(time.Minute)
				claims["exp"] = expire.Unix()
				claims["orig_iat"] = now.Unix()
			})
			if er != nil {
				global.FAIL(c, "fail.msg", er.Error())
				return
			}
			ln = fmt.Sprintf(`%s"%s/v2/task/kms?Ciphertext=%s"`, sps[0], global.CONF.Storage.Host, token)
		} else if strings.HasSuffix(ln, ".ts") {
			for _, proxyURL := range global.CONF.Site.Play.ProxyUrl {
				if strings.HasPrefix(ln, proxyURL) {
					ln = "/v2/task/play/part?p=" + ln
					break
				}
			}
		}
		buff.WriteString(ln + "\n")
	}

	l.RewriteHls = buff.Bytes()

	c.Data(200, "application/x-mpegurl", l.RewriteHls)
}

func (t *Task) PlayPart(c *gin.Context) {
	var req task.TaskPlayPartRequest
	if err := c.ShouldBind(&req); err != nil {
		global.FAIL(c, "fail.msg", err.Error())
		return
	}

	if err := zhttp.NewRequest().
		Before(func(r *http.Request) {
			r.Header.Set("origin", "https://www.geekbang.org")
			r.Header.Set("referer", "https://www.geekbang.org")
		}).
		After(func(r *http.Response) error {
			if r.StatusCode == 200 {
				c.Render(http.StatusOK, render.Reader{
					ContentLength: -1,
					ContentType:   "application/octet-stream",
					Reader:        r.Body,
				})
				return nil
			}
			return fmt.Errorf("not found part [%s]", req.P)
		}).Do(http.MethodGet, req.P, nil); err != nil {
		global.FAIL(c, "fail.msg", err.Error())
		return
	}
}

func (t *Task) Export(c *gin.Context) {
	var req task.TaskExportRequest
	if err := c.ShouldBind(&req); err != nil {
		global.FAIL(c, "fail.msg", err.Error())
		return
	}
	var l model.Task
	if err := global.DB.Model(&model.Task{}).
		Where(&model.Task{TaskId: req.Pid}).First(&l).Error; err != nil {
		global.FAIL(c, "fail.msg", err.Error())
		return
	}
	var statistics task.TaskStatistics
	if err := json.Unmarshal(l.Statistics, &statistics); err != nil {
		global.FAIL(c, "fail.msg", err.Error())
		return
	}
	finishCount, ok := statistics.Items[service.TASK_STATUS_FINISHED]
	if !ok {
		global.FAIL(c, "tasks.illegal_operation", 0, statistics.Count)
		return
	}
	if statistics.Count != finishCount {
		global.FAIL(c, "tasks.illegal_operation", finishCount, statistics.Count)
		return
	}
	var product geek.ProductBase
	if err := json.Unmarshal(l.Raw, &product); err != nil {
		global.FAIL(c, "fail.msg", err.Error())
		return
	}
	var taskMessage task.TaskMessage
	if len(l.Message) > 0 {
		if err := json.Unmarshal(l.Message, &taskMessage); err != nil {
			global.FAIL(c, "fail.msg", err.Error())
			return
		}
	}
	switch req.Type {
	case "markdown":
		dirName := service.VerifyFileName(product.Title)
		archiveName := fmt.Sprintf("%s.tar.gz", dirName)
		buf, err := service.MakeDocArchive(c, l.TaskId, product.Title, product.IntroHTML)
		if err != nil {
			global.FAIL(c, "fail.msg", err.Error())
			return
		}
		c.Header("Content-Type", "application/octet-stream")
		c.Header("Content-Disposition", "attachment; filename="+url.QueryEscape(archiveName))
		c.Header("Content-Transfer-Encoding", "binary")
		c.Data(200, "application/octet-stream", buf.Bytes())
	case "docsite":
		docURL, err := service.MakeDocsite(c, l.TaskId, product.Title, product.IntroHTML)
		if err != nil {
			global.FAIL(c, "fail.msg", err.Error())
			return
		}
		taskMessage.Doc = docURL
		l.Message, _ = json.Marshal(taskMessage)
		if err = global.DB.Model(&model.Task{}).
			Where(&model.Task{Id: l.Id}).
			UpdateColumn("message", l.Message).Error; err != nil {
			global.LOG.Error("worker Updates", zap.Error(err), zap.String("taskId", l.TaskId))
			global.FAIL(c, "fail.msg", err.Error())
			return
		}
		global.OK(c, nil)
	}
}
