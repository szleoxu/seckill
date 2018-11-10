package main

import (
	"fmt"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	"os"
	"time"
	"strings"
	"bufio"
)

const (
	seleniumPath = "chromedriver"
	port            = 9515
)

func login(){


	//链接本地的浏览器 chrome
	caps := selenium.Capabilities{
		"browserName": "chrome",
	}
	//禁止图片加载，加快渲染速度
	imagCaps := map[string]interface{}{
		"profile.managed_default_content_settings.images": 2,
	}
	chromeCaps := chrome.Capabilities{
		Prefs: imagCaps,
		Path:  "",
		Args: []string{
			//"--headless", // 设置Chrome无头模式，在linux下运行，需要设置这个参数，否则会报错
			//"--no-sandbox",
			"--user-agent=Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36", // 模拟user-agent，防反爬
		},
	}
	//以上是设置浏览器参数
	caps.AddChrome(chromeCaps)

	// 调起chrome浏览器
	w_b1, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		fmt.Println("connect to the webDriver faild", err.Error())
		return
	}
	//关闭一个webDriver会对应关闭一个chrome窗口
	//但是不会导致seleniumServer关闭
	defer w_b1.Quit()
	err = w_b1.Get("https://zhuanlan.zhihu.com/p/37752206")

	if err != nil {
		fmt.Println("get page faild", err.Error())
		return
	}



	/*# 打开淘宝登录页，并进行扫码登录
	driver.get("https://www.taobao.com")
	time.sleep(3)
	if driver.find_element_by_link_text("亲，请登录"):
	driver.find_element_by_link_text("亲，请登录").click()
	print("请在15秒内完成扫码")
	time.sleep(15)
	driver.get("https://cart.taobao.com/cart.htm")
	time.sleep(3)
	# 点击购物车里全选按钮
	if driver.find_element_by_id("J_SelectAll1"):
	driver.find_element_by_id("J_SelectAll1").click()
	now = datetime.datetime.now()
	print('login success:', now.strftime('%Y-%m-%d %H:%M:%S'))*/

}

func main() {

	chromeB()
}

func firefox(){
	// Start a Selenium WebDriver server instance (if one is not already
	// running).
	const (
		// These paths will be different on your system.
		seleniumPath    = "vendor/selenium-server-standalone-3.4.jar"
		geckoDriverPath = "vendor/geckodriver-v0.18.0-linux64"
		port            = 8080
	)
	opts := []selenium.ServiceOption{
		selenium.StartFrameBuffer(),           // Start an X frame buffer for the browser to run in.
		selenium.GeckoDriver(geckoDriverPath), // Specify the path to GeckoDriver in order to use Firefox.
		selenium.Output(os.Stderr),            // Output debug information to STDERR.
	}
	selenium.SetDebug(true)
	service, err := selenium.NewSeleniumService(seleniumPath, port, opts...)
	if err != nil {
		panic(err) // panic is used only as an example and is not otherwise recommended.
	}
	defer service.Stop()

	// Connect to the WebDriver instance running locally.
	caps := selenium.Capabilities{"browserName": "firefox"}
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		panic(err)
	}
	defer wd.Quit()

	// Navigate to the simple playground interface.
	if err := wd.Get("http://play.golang.org/?simple=1"); err != nil {
		panic(err)
	}

	// Get a reference to the text box containing code.
	elem, err := wd.FindElement(selenium.ByCSSSelector, "#code")
	if err != nil {
		panic(err)
	}
	// Remove the boilerplate code already in the text box.
	if err := elem.Clear(); err != nil {
		panic(err)
	}

	// Enter some new code in text box.
	err = elem.SendKeys(`
		package main
		import "fmt"
		func main() {
			fmt.Println("Hello WebDriver!\n")
		}
	`)
	if err != nil {
		panic(err)
	}

	// Click the run button.
	btn, err := wd.FindElement(selenium.ByCSSSelector, "#run")
	if err != nil {
		panic(err)
	}
	if err := btn.Click(); err != nil {
		panic(err)
	}

	// Wait for the program to finish running and get the output.
	outputDiv, err := wd.FindElement(selenium.ByCSSSelector, "#output")
	if err != nil {
		panic(err)
	}

	var output string
	for {
		output, err = outputDiv.Text()
		if err != nil {
			panic(err)
		}
		if output != "Waiting for remote server..." {
			break
		}
		time.Sleep(time.Millisecond * 100)
	}

	fmt.Printf("%s", strings.Replace(output, "\n\n", "\n", -1))

	// Example Output:
	// Hello WebDriver!
	//
	// Program exited.
}

func chromeB(){
	//opts := []selenium.ServiceOption{
	//    selenium.StartFrameBuffer(),           // Start an X frame buffer for the browser to run in.
	//    selenium.GeckoDriver(geckoDriverPath), // Specify the path to GeckoDriver in order to use Firefox.
	//}
	//selenium.SetDebug(true)

	//如果seleniumServer没有启动，就启动一个seleniumServer所需要的参数，可以为空，示例请参见https://github.com/tebeka/selenium/blob/master/example_test.go
	opts := []selenium.ServiceOption{}
	service, err := selenium.NewChromeDriverService(seleniumPath, port, opts...)
	if nil != err {
		fmt.Println("start a chromedriver service falid", err.Error())
		return
	}
	//注意这里，server关闭之后，chrome窗口也会关闭
	defer service.Stop()
	//链接本地的浏览器 chrome
	caps := selenium.Capabilities{
		"browserName": "chrome",
	}
	//禁止图片加载，加快渲染速度
	imagCaps := map[string]interface{}{
		"profile.managed_default_content_settings.images": 2,
	}
	chromeCaps := chrome.Capabilities{
		Prefs: imagCaps,
		Path:  "",
		Args: []string{
			"--headless", // 设置Chrome无头模式，在linux下运行，需要设置这个参数，否则会报错
			"--no-sandbox",
			"--user-agent=Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.97 Safari/537.36", // 模拟user-agent，防反爬
		},
	}
	//以上是设置浏览器参数
	caps.AddChrome(chromeCaps)

	// 调起chrome浏览器
	w_b1, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		fmt.Println("connect to the webDriver faild", err.Error())
		return
	}
	//关闭一个webDriver会对应关闭一个chrome窗口
	//但是不会导致seleniumServer关闭
	defer w_b1.Quit()
	err = w_b1.Get("https://www.taobao.com")
	if err != nil {
		fmt.Println("get page faild", err.Error())
		return
	}
	tit,_:=w_b1.Title()
	fmt.Println(tit)
	loginLink,_:=w_b1.FindElement(selenium.ByLinkText,"亲，请登录")
	loginLink.Click()
	print("5秒后开始登陆")
	time.Sleep(5*time.Second)
	userNameText,_:=w_b1.FindElement(selenium.ByID,"TPL_username_1")
	userNameText.SendKeys("13*********")
	pwdText,_:=w_b1.FindElement(selenium.ByID,"TPL_password_1")
	/*fmt.Println("input pwd:")
	inputPwd := bufio.NewScanner(os.Stdin)
	inputPwd.Scan()*/
	pwdText.SendKeys("pwd****")
	fmt.Println("input complete")
	loginBtn,_:=w_b1.FindElement(selenium.ByID,"J_SubmitStatic")

	draggerSpan,_:=w_b1.FindElement(selenium.ByID,"nc_1_n1z")

	time.Sleep(2*time.Second)

	for i := 0; i < 500; i++ {
		draggerSpan.MoveTo(i,0)

	}

	time.Sleep(2*time.Second)

	loginBtn.Click()
	print("15秒后开始验证手机登陆")
	time.Sleep(15*time.Second)
	/*err=w_b1.Get("https://login.taobao.com/member/login_unusual.htm")
	if err != nil {
		fmt.Println("get page faild", err.Error())
		os.Exit(1)
	}*/
	w_b1.Screenshot()
	fmt.Println(w_b1.Title())
	getCodeBtn,err:=w_b1.FindElement(selenium.ByID,"J_GetCode")
	if err!=nil{
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(getCodeBtn.Text())
	getCodeBtn.Click()
	checkCodeText,_:=w_b1.FindElement(selenium.ByID,"J_Phone_Checkcode")
	fmt.Println("input phone code:")
	inputCode := bufio.NewScanner(os.Stdin)
	inputCode.Scan()
	checkCodeText.SendKeys(inputCode.Text())
	fmt.Println("input complete")
	submitBtn,_:=w_b1.FindElement(selenium.ByID,"submitBtn")
	submitBtn.Click()

	os.Exit(1)

	/*// 重新调起chrome浏览器
	w_b2, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		fmt.Println("connect to the webDriver faild", err.Error())
		return
	}
	defer w_b2.Close()
	//打开一个网页
	err = w_b2.Get("https://www.toutiao.com/")
	if err != nil {
		fmt.Println("get page faild", err.Error())
		return
	}
	//打开一个网页
	err = w_b2.Get("https://www.baidu.com/")
	if err != nil {
		fmt.Println("get page faild", err.Error())
		return
	}
	//w_b就是当前页面的对象，通过该对象可以操作当前页面了
	//........
	time.Sleep(5* time.Minute)
	return*/
}