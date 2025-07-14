package file

import (
	"Shop/internal/consts"
	"Shop/internal/dao"
	"Shop/internal/model"
	"Shop/internal/model/entity"
	"Shop/internal/service"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"time"
)

type sFile struct {
}

func init() {
	service.RegisterFile(New())
}

func New() *sFile {
	return &sFile{}
}

/*
1.定义图片上传位置
2.校验：上传位置是否正确，安全性校验：1分钟内只能上传10次
3.定义年月日Ymd
4.入库
5.返回数据
*/
func (s *sFile) Upload(ctx context.Context, in model.FileUploadInput) (out *model.FileUploadOutput, err error) {
	//1.定义图片上传位置
	uploadPath := g.Cfg().MustGet(ctx, "upload.path").String()
	if uploadPath == "" {
		return nil, gerror.New("读取配置文件失败 上传路径不存在")
	}
	if in.Name != "" {
		in.File.Filename = in.Name
	}
	//2.安全性校验：每个人1分钟内只能上传10次
	count, err := dao.FileInfo.Ctx(ctx).
		Where(dao.FileInfo.Columns().UserId, gconv.Int(ctx.Value(consts.CtxAdminId))).
		WhereGTE(dao.FileInfo.Columns().CreatedAt, gtime.Now().Add(time.Minute)).Count()
	if err != nil {
		return nil, err
	}
	if count > consts.FileMaxUploadCountMinute {
		return nil, gerror.New("上传频繁，1分钟内只能上传10次")
	}
	//3.定义年月日YMD
	dateDirName := gtime.Now().Format("Ymd")
	//gfile.Join 用“/"拼接
	fileName, err := in.File.Save(gfile.Join(uploadPath, dateDirName), in.RandomName)
	if err != nil {
		return nil, err
	}
	//4.入库
	data := entity.FileInfo{
		Name:   fileName,
		Src:    gfile.Join(uploadPath, dateDirName, fileName),
		Url:    "upload/" + dateDirName + "/" + fileName, //和上面gfile.Join效果一样
		UserId: gconv.Int(ctx.Value(consts.CtxAdminId)),
	}
	id, err := dao.FileInfo.Ctx(ctx).Data(data).OmitEmpty().InsertAndGetId()
	if err != nil {
		return nil, err
	}
	return &model.FileUploadOutput{
		Id:   uint(id),
		Name: data.Name,
		Src:  data.Src,
		Url:  data.Url,
	}, nil
}
