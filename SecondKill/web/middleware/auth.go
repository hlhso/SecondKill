package middleware

import "github.com/kataras/iris/v12"

func AuthConProduct(ctx iris.Context) {

	uid := ctx.GetCookie("uid")
	if uid == "" {
		ctx.Application().Logger().Debug("必须先登录!")
		ctx.Redirect("/user/login")
		return
	}
	ctx.Application().Logger().Debug("sign:" + ctx.GetCookie("sign"))
	ctx.Application().Logger().Debug("已经登陆")
	ctx.Next()
}

/*func AuthConProduct(ctx iris.Context) {
	uid := ctx.GetCookie("sign")
	fmt.Println(uid)
	if uid == "" {
		ctx.Application().Logger().Debug("必须先登录!")
		ctx.Redirect("/user/login")
		return
	}
	code, err := encrypt.DePwdCode(uid)
	fmt.Println(code,err)
	if err != nil {
		ctx.Application().Logger().Error("登录校验失败：" + err.Error())
		return
	}
	ctx.Application().Logger().Debug("sign:" + string(code))
	ctx.Next()
}*/

