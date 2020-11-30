package main

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"math/rand"
	"os"
)

/*
点是可以绘制的最简单的图形对象。

在这例子中，我们在窗口上随机绘制了1000个红点
*/

func InitUi() *widgets.QMainWindow {
	// 创建窗口
	app := widgets.NewQMainWindow(nil, 0)

	// 设置窗口的标题
	app.SetWindowTitle("Qt 教程")

	// 设置窗口的位置和大小
	app.SetGeometry2(300, 300, 300, 220)

	// 设置窗口的图标，引用当前目录下的web.png图片
	app.SetWindowIcon(gui.NewQIcon5("images/app.ico"))

	// 布局窗口组件载体
	widget := widgets.NewQWidget(app, core.Qt__Widget)
	//widget.SetGeometry(core.NewQRect4(300, 300, 300, 220))
	//widget.SetGeometry2(0, 0, 300, 220)
	app.SetCentralWidget(widget)

	// 状态栏
	app.StatusBar()

	widget.ConnectPaintEvent(func(event *gui.QPaintEvent) {
		qp := gui.NewQPainter2(widget)
		//qp.Begin(widget)

		//qp.SetPen2(core.Qt__red)
		//每次我们改变窗口的大小,生成一个 paint event 事件。
		//我们得到的当前窗口的大小size。我们使用窗口的大小来分配点在窗口的客户区。demo.go
		size := widget.Size()

		for i := 0; i < 1000; i++ {
			rand.Int()
			x := rand.Intn(size.Width() - 1)
			y := rand.Intn(size.Height() - 1)
			qp.DrawPoint3(x, y)
		}

		qp.End()

	})

	return app
}

func main() {
	// 创建一个应用程序对象
	// sys.argv参数是一个列表，从命令行输入参数
	widgets.NewQApplication(len(os.Args), os.Args)

	// 初始化窗口
	app := InitUi()

	// 显示组件
	app.Show()

	// 确保应用程序干净的退出
	widgets.QApplication_Exec()
}
