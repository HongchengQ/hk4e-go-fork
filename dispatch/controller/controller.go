package controller

import (
	"context"
	"net/http"
	"strconv"

	"hk4e/common/config"
	"hk4e/common/region"
	"hk4e/common/rpc"
	"hk4e/dispatch/dao"
	"hk4e/node/api"
	"hk4e/pkg/logger"
	"hk4e/pkg/random"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	dao          *dao.Dao
	discovery    *rpc.DiscoveryClient
	signRsaKey   []byte
	encRsaKeyMap map[string][]byte
	pwdRsaKey    []byte
	ec2b         *random.Ec2b
}

func NewController(dao *dao.Dao, discovery *rpc.DiscoveryClient) (r *Controller) {
	r = new(Controller)
	r.dao = dao
	r.discovery = discovery
	r.signRsaKey, r.encRsaKeyMap, r.pwdRsaKey = region.LoadRsaKey()
	rsp, err := r.discovery.GetRegionEc2B(context.TODO(), &api.NullMsg{})
	if err != nil {
		logger.Error("get region ec2b error: %v", err)
		return nil
	}
	ec2b, err := random.LoadEc2bKey(rsp.Data)
	if err != nil {
		logger.Error("parse region ec2b error: %v", err)
		return nil
	}
	r.ec2b = ec2b
	go r.registerRouter()
	return r
}

func (c *Controller) authorize() gin.HandlerFunc {
	return func(context *gin.Context) {
		if context.Query("key") == "flswld" {
			context.Next()
			return
		}
		context.Abort()
		context.JSON(http.StatusOK, gin.H{
			"code": "10001",
			"msg":  "没有访问权限",
		})
	}
}

func (c *Controller) registerRouter() {
	if config.CONF.Logger.Level == "DEBUG" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	engine := gin.Default()
	{
		// 404
		engine.NoRoute(func(context *gin.Context) {
			logger.Info("no route find, fallback to fuck mhy, url: %v", context.Request.RequestURI)
			context.Header("Content-type", "text/html; charset=UTF-8")
			context.Status(http.StatusNotFound)
			_, _ = context.Writer.WriteString("FUCK MHY")
		})
	}
	{
		// 调度
		// dispatchosglobal.yuanshen.com
		engine.GET("/query_security_file", c.querySecurityFile)
		engine.GET("/query_region_list", c.queryRegionList)
		// osusadispatch.yuanshen.com
		engine.GET("/query_cur_region", c.queryCurRegion)
	}
	{
		// 登录
		// hk4e-sdk-os.hoyoverse.com
		// 账号登录
		engine.POST("/hk4e_global/mdk/shield/api/login", c.apiLogin)
		// token登录
		engine.POST("/hk4e_global/mdk/shield/api/verify", c.apiVerify)
		// 获取combo token
		engine.POST("/hk4e_global/combo/granter/login/v2/login", c.v2Login)
	}
	{
		// 日志
		engine.POST("/sdk/dataUpload", c.sdkDataUpload)
		engine.GET("/perf/config/verify", c.perfConfigVerify)
		engine.POST("/perf/dataUpload", c.perfDataUpload)
		engine.POST("/log", c.log8888)
		engine.POST("/crash/dataUpload", c.crashDataUpload)
	}
	{
		// 返回固定数据
		// Windows
		engine.GET("/hk4e_global/mdk/agreement/api/getAgreementInfos", c.getAgreementInfos)
		engine.POST("/hk4e_global/combo/granter/api/compareProtocolVersion", c.postCompareProtocolVersion)
		engine.POST("/account/risky/api/check", c.check)
		engine.GET("/combo/box/api/config/sdk/combo", c.combo)
		engine.GET("/hk4e_global/combo/granter/api/getConfig", c.getConfig)
		engine.GET("/hk4e_global/mdk/shield/api/loadConfig", c.loadConfig)
		engine.POST("/data_abtest_api/config/experiment/list", c.list)
		engine.GET("/admin/mi18n/plat_oversea/m2020030410/m2020030410-version.json", c.version10Json)
		engine.GET("/admin/mi18n/plat_oversea/m2020030410/m2020030410-zh-cn.json", c.zhCN10Json)
		engine.GET("/geetestV2.html", c.geetestV2)
		// Android
		engine.POST("/common/h5log/log/batch", c.batch)
		engine.GET("/hk4e_global/combo/granter/api/getFont", c.getFont)
		engine.GET("/admin/mi18n/plat_oversea/m202003049/m202003049-version.json", c.version9Json)
		engine.GET("/admin/mi18n/plat_oversea/m202003049/m202003049-zh-cn.json", c.zhCN9Json)
		engine.GET("/hk4e_global/combo/granter/api/compareProtocolVersion", c.getCompareProtocolVersion)
		// Android geetest
		engine.GET("/gettype.php", c.gettype)
		engine.GET("/get.php", c.get)
		engine.POST("/ajax.php", c.ajax)
		engine.GET("/ajax.php", c.ajax)
		engine.GET("/static/appweb/app3-index.html", c.app3Index)
		engine.GET("/static/js/slide.7.8.6.js", c.slideJs)
		engine.GET("/favicon.ico", c.faviconIco)
		engine.GET("/static/js/gct.e7810b5b525994e2fb1f89135f8df14a.js", c.js)
		engine.GET("/static/ant/style_https.1.2.6.css", c.css)
		engine.GET("/pictures/gt/a330cf996/a330cf996.webp", c.webp)
		engine.GET("/pictures/gt/a330cf996/bg/86f9db021.webp", c.bgWebp)
		engine.GET("/pictures/gt/a330cf996/slice/86f9db021.png", c.slicePng)
		engine.GET("/static/ant/sprite2x.1.2.6.png", c.sprite2xPng)
	}
	engine.Use(c.authorize())
	engine.POST("/gate/token/verify", c.gateTokenVerify)
	engine.POST("/gate/token/reset", c.gateTokenReset)
	port := config.CONF.HttpPort
	addr := ":" + strconv.Itoa(int(port))
	err := engine.Run(addr)
	if err != nil {
		logger.Error("gin run error: %v", err)
	}
}
