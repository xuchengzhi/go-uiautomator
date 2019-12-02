package UIAutomatorServer
// package main

import (
    "log"
    // "regexp"
    "github.com/xuchengzhi/UIautomator2/Base"
    "github.com/xuchengzhi/Library/Time"
    // "github.com/xuchengzhi/Library/Http"
    "github.com/xuchengzhi/Library/Encryption"
    "fmt"

)

var (
    url string
    port int
    server_url string
    addrs string
)

type APKManager struct {
    Path         string
    packageName  string
    mainActivity string
}

type U2 interface{
    Connect_wifi()
    Info()
}


type Server struct {
    Addr string
    server_url string
}


func  Server_url(s string) string {
    server_url := fmt.Sprintf("http://%v:%v",s,port)
    return server_url
}   


func init() {
    url = "127.0.0.1"
    port = 7912
    server_url = fmt.Sprintf("http://%v:%v",url,port)
}


type shellStr struct {
    Command string `json:"command"`
    Timeout int  `json:"timeout"`
}


func Adbshell(s Server,cmdargs string, stream bool, timeout int) bool {
    var params = make(map[string]string)
    log.Println(cmdargs)
    params["command"] = cmdargs
    params["timeout"] = "60"
    urls := Shellurl(s.Addr)
    res := Base.Post(urls,params,false)
    if res.Status == 200{
        return true
    }else{
        return false
    }
}

func (s Server) DevInfo() Base.ApiJson {
    var par Base.JsonStr
    par.ID = Jsonrpc_id("deviceInfo")
    par.Jsonrpc = "2.0"
    par.Method = "deviceInfo"
    var parms Base.Par
    parms.Url = Jsonrpcurl(s.Addr)
    parms.Params =par
    res := Base.PostJson(parms,false)
    return res
}

func Jsonrpc_id(methods string) string{
    tmp := fmt.Sprintf("%v at %v",methods,GetTime.TS())
    return XorEnc.Gmd5(tmp)
}


func Shellurl(s string) string{
    hurl := Server_url(s)
    jsonurl :=  fmt.Sprintf("%v/shell",hurl) 
    return jsonurl
}


func Jsonrpcurl(s string) string {
    hurl := Server_url(s)
    jsonurl :=  fmt.Sprintf("%v/jsonrpc/0",hurl) 
    return jsonurl
}


type AppInfo struct {
    Pkg_name string
    Activity string
    Extras interface{}
    Wait bool
    Stop bool
    Unlock bool
}


func App_start(s Server,appname string) bool {
    apps := fmt.Sprintf("monkey -p %v -c android.intent.category.LAUNCHER 1",appname)
    res := Adbshell(s,apps,false,1)
    return res
}







// func init(host, port=7912){
    
//     self._host = host
//     self._port = port
//     self._reqsess = TimeoutRequestsSession(
//     )  # use requests.Session to enable HTTP Keep-Alive
//     self._server_url = "http://{:{".format(host, port)
//     self._server_jsonrpc_url = self._server_url + "/jsonrpc/0"
//     self._default_session = Session(self, None)
//     self._cached_plugins = {
//     self.__devinfo = None
//     self._hooks = {
//     self.platform = None  # hot fix for weditor

//     self.ash = AdbShell(self.shell)  # the powerful adb shell
//     self.wait_timeout = 20.0  # wait element timeout
//     self.click_post_delay = None  # wait after each click
//     self._freeze()  # prevent creating new attrs
//     # self._atx_agent_check()
// }

// func _freeze(self){
//     self.__isfrozen = True
// }

// func plugins(){
//     return UIAutomatorServer.__plugins
// }

// func __setattr__(self, key, value){
//     // """
//      // Prevent creating new attributes outside __}init__ """

//     if self.__isfrozen and not hasattr(self, key){
//         raise TypeError("Key %s does not exist in class %r" % (key, self))
//     object.__setattr__(self, key, value)
// }


// func serial(self){
//     return self.shell(["getprop", "ro.serialno"])[0].strip()
// }

// func jsonrpc(self){
//     // """

//     Make jsonrpc call easier
//     For example:
//         self.jsonrpc.pressKey("home")
//     // """

//     return self.setup_jsonrpc()
// }

// func path2url(self, path){
//     return urlparse.urljoin(self._server_url, path)
// }

// func window_size(self){
//     // """
//      // return (width, height) """

//     info = self._reqsess.get(self.path2url("/info")).json()
//     w, h = info["display"]["width"], info["display"]["heig}ht"]
//     if (w > h) != (self.info["displayRotation"] % 2 == 1){
//         w, h = h, w
//     return w, h
// }

// func hooks_register(self, func){
//     // """

//     Args:
//         func: should accept 3 args. func_name:string, args:tuple, kwargs:dict
//     // """

//     self._hooks[func] = True
// }

// func hooks_apply(self, stage, func_name, args=(), kwargs={, ret=None){
//     // """

//     // Args:}
//     //     stage(str){ one of "before" or "after"
//     // """
//     for fn in self._hooks.keys(){
//         fn(stage, func_name, args, kwargs, ret)
// }

// func setup_jsonrpc(self, jsonrpc_url=None){
//     // """

//     // Wrap jsonrpc call into object
//     // Usage example:
//     //     self.setup_jsonrpc().pressKey("home")
//     // """

//     if not jsonrpc_url:
//         jsonrpc_url = self._server_jsonrpc_url
// }

// func alive(self){
//     try:
//         r = self._reqsess.get(self.path2url("/ping"), timeout=2)
//         if r.status_code != 200:
//             return False
//         r = self._reqsess.post(
//             self.path2url("/jsonrpc/0"),
//             data=json.dumps({
//                 "jsonrpc": "2.0",
//                 "id": 1,
//                 "method": "deviceInfo"
//             }),
//             timeout=2)
//         if r.status_code != 200:
//             return False}
//         if r.json().get("error"){
//             return False
//         return True
//     except requests.exceptions.ReadTimeout:
//         return False
//     except EnvironmentError:
//         return False
// }
//     func service(self, name){
//         // """
//          Manage service start or stop

//         Example:
//             d.service("uiautomator").start()
//             d.service("uiautomator").stop()
//         // """

//         u2obj = self
// }
//         class _Service(object){}
//             func __init__(self, name){
//                 self.name = n}ame
//                 # FIXME(ssx){ support other service: minicap, minitouch
//                 assert name == "uiautomator"
// }
//             func start(self){
//                 res = u2obj._reqsess.post(u2obj.path2url("/uiautomator"))
//                 res.raise_for_status()
// }
//             func stop(self){
//                 res = u2obj._reqsess.delete(u2obj.path2url("/uiautomator"))
//                 if res.status_code != 200:
//                     warnings.warn(res.text)

//         return _Service(name)
// }

    
// func app_install(self, url, installing_callback=None, server=None){
//     // """

//     {u"message": u"downloading", "progress": {u"totalSize": 407992690, u"copiedSize": 49152}}

//     Returns:
//         packageName

//     Raises:
//         RuntimeError
//     // """

//     r = self._reqsess.post(self.path2url("/install"), data={"url": url})
//     if r.status_code != 200:
//         raise RuntimeError("app install error:", r.text)
//     id = r.text.strip()
//     print(time.strftime("%H:%M:%S"), "id:", id)
//     return self._wait_install_finished(id, installing_callback)
// }
// func _wait_install_finished(self, id, installing_callback){
//     bar = None
//     downloaded = True

//     while True:
//         resp = self._reqsess.get(self.path2url("/install/" + id))
//         resp.raise_for_status()
//         jdata = resp.json()
//         message = jdata["message"]
//         pg = jdata.get("progress")
// }
//             func notty_print_progress(pg){
//                 written = pg["copiedSize"]
//                 total = pg["totalSize"]
//                 print(
//                     time.strftime("%H:%M:%S"), "downloading %.1f%% [%s/%s]" %
//                     (100.0 * written / total,
//                      humanize.naturalsize(written, gnu=True),
//                      humanize.naturalsize(total, gnu=True)))

//             if message == "downloading":
//                 downloaded = False
//                 if pg:  # if there is a progress}
//                     if hasattr(sys.stdout, "isa}tty"){
//                         if sys.stdout.isatty(){
//                             if not bar:
//                                 bar = _ProgressBar(
//                                     time.strftime("%H:%M:%S") + " downloading",
//                                     max=pg["totalSize"])
//                             written = pg["copiedSize"]
//                             bar.next(written - bar.index)
//                         else:
//                             notty_print_progress(pg)
//                     else:
//                         pass
//                 else:
//                     print(time.strftime("%H:%M:%S"), "download initialing")
//             else:
//                 if not downloaded:
//                     downloaded = True
//                     if bar:  # bar only set in atty
//                         bar.next(pg["copiedSize"] - bar.index) if pg else None
//                         bar.finish()
//                     else:
//                         print(time.strftime("%H:%M:%S"), "download 100%")
//                 print(time.strftime("%H:%M:%S"), message)
//             if message == "installing":}
//                 if callable(installing_callback){
//                     installing_callback(self)
//             if message == "success installed":
//                 return jdata.get("packageName")
// }
//             if jdata.get("error"){
//                 raise RuntimeError("error", jdata.get("error"))

//             try:
//                 time.sleep(1)
//             except KeyboardInterrupt:
//                 bar.finish() if bar else None
//                 print("keyboard interrupt catched, cancel install id", id)
//                 self._reqsess.delete(self.path2url("/install/" + id))
//                 raise
// }
//     func shell(self, cmdargs, stream=False, timeout=60){
//         // """

//         Run adb shell command with arguments and return its output. Require atx-agent >=0.3.3

//         Args:
//             cmdargs: str or list, example: "ls -l" or ["ls", "-l"]
//             timeout: seconds of command run, works on when stream is False
//             stream: bool used for long running process.

//         Returns:
//             (output, exit_code) when stream is False
//             requests.Response when stream is True, you have to close it after using

//         Raises:
//             RuntimeError

//         For atx-agent is not support return exit code now.
//         When command got something wrong, exit_code is always 1, otherwise exit_code is always 0
//         // """
//     }
//         if isinstance(cmdargs, (list, tuple)){
//             cmdargs = list2cmdline(cmdargs)
//         if stream:
//             return self._reqsess.get(
//                 self.path2url("/shell/stream"),
//                 params={"command": cmdargs},
//                 stream=True)
//         ret = self._reqsess.post(
//             self.path2url("/shell"),
//             data={
//                 "command": cmdargs,
//                 "timeout": str(timeout)
//             })
//         if ret.status_code != 200:
//             raise RuntimeError("device agent responds with an error code %d" %
//                                ret.status_code, ret.text)
//         resp = ret.json()
//         exit_code = 1 if resp.get("error") else 0
//         exit_code = resp.get("exitCode", exit_code)
//         shell_response = namedtuple("ShellResponse", ("output", "exit_code"))
//         return shell_response(resp.get("output"), exit_code)
// }
//     func adb_shell(self, *args){
//         // """

//         Example:
//             adb_shell("pwd")
//             adb_shell("ls", "-l")
//             adb_shell("ls -l")

//         Returns:
//             string for stdout merged with stderr, after the entire shell command is completed.
//         // """

//         # print(
//         #     "DeprecatedWarning: adb_shell is deprecated, use: output, exit_code = shell(["ls", "-l"]) instead"
//         # )
//         cmdline = args[0] if len(args) == 1 else list2cmdline(args)
//         return self.shell(cmdline)[0]



    
//     func current_app(self){
//         // """

//         Returns:
//             dict(package, activity, pid?)

//         Raises:
//             EnvironementError

//         For developer:
//             Function reset_uiautomator need this function, so can"t use jsonrpc here.
//         // """

//         # Related issue: https://github.com/openatx/uiautomator2/issues/200
//         # $ adb shell dumpsys window windows
//         # Example output:
//         #   mCurrentFocus=Window{41b37570 u0 com.incall.apps.launcher/com.incall.apps.launcher.Launcher}
//         #   mFocusedApp=AppWindowToken{422df168 token=Token{422def98 ActivityRecord{422dee38 u0 com.example/.UI.play.PlayActivity t14}}}
//         # Regexp
//         #   r"mFocusedApp=.*ActivityRecord{\w+ \w+ (?P<package>.*)/(?P<activity>.*) .*"
//         #   r"mCurrentFocus=Window{\w+ \w+ (?P<package>.*)/(?P<activity>.*)\}")
//         _focusedRE = re.compile(
//             r"mCurrentFocus=Window{.*\s+(?P<package>[^\s]+)/(?P<activity>[^\s]+)\}")
//         m = _focusedRE.search(self.shell(["dumpsys", "window", "windows"])[0])
//         if m:
//             return dict(
//                 package=m.group("package"), activity=m.group("activity"))

//         # try: adb shell dumpsys activity top
//         _activityRE = re.compile(
//             r"ACTIVITY (?P<package>[^\s]+)/(?P<activity>[^/\s]+) \w+ pid=(?P<pid>\d+)"
//         )
//         output, _ = self.shell(["dumpsys", "activity", "top"])
//         ms = _activityRE.finditer(output)
//         ret = None
//         for m in ms:
//             ret = dict(
//                 package=m.group("package"),
//                 activity=m.group("activity"),
//                 pid=int(m.group("pid")))
//         if ret:  # get last result
//             return ret
//         raise EnvironmentError("Couldn"t get focused app")
// }
//     func wait_activity(self, activity, timeout=10){
//         // """
//          wait activity
//         Args:}
//             activity (str){} name of activity
//             timeout (float){ max wait time

//         Returns:
//             bool of activity
//         // """

//         deadline = time.time() + timeout
//         while time.time() < deadline:
//             current_activity = self.current_app().get("activity")
//             if activity == current_activity:
//                 return True
//             time.sleep(.5)
//         return False
// }
//     func app_stop(self, pkg_name){
//         // """
//          // Stop one application: am force-stop"""

//         self.shell(["am", "force-stop", pkg_name])
// }
//     func app_stop_all(self, excludes=[]){
//         // """
//          Stop all third party applications
//         Args:}
//             excludes (list){ apps that do now want to kill

//         Returns:
//             a list of killed apps
//         // """

//         our_apps = ["com.github.uiautomator", "com.github.uiautomator.test"]
//         output, _ = self.shell(["pm", "list", "packages", "-3"])
//         pkgs = re.findall("package:([^\s]+)", output)
//         process_names = re.findall("([^\s]+)$", self.shell("ps")[0], re.M)
//         kill_pkgs = set(pkgs).intersection(process_names).difference(
//             our_apps + excludes)
//         kill_pkgs = list(kill_pkgs)
//         for pkg_name in kill_pkgs:
//             self.app_stop(pkg_name)
//         return kill_pkgs
// }
//     func app_clear(self, pkg_name){
//         // """
//          // Stop and clear app data: pm clear """

//         self.shell(["pm", "clear", pkg_name])
// }
//     func app_uninstall(self, pkg_name){
//         // """
//          // Uninstall an app """

//         self.shell(["pm", "uninstall", pkg_name])
// }
//     func app_uninstall_all(self, excludes=[], verbose=False){
//         // """
//          // Uninstall all apps """

//         our_apps = ["com.github.uiautomator", "com.github.uiautomator.test"]
//         output, _ = self.shell(["pm", "list", "packages", "-3"])
//         pkgs = re.findall("package:([^\s]+)", output)
//         pkgs = set(pkgs).difference(our_apps + excludes)
//         pkgs = list(pkgs)
//         for pkg_name in pkgs:
//             if verbose:
//                 print("uninstalling", pkg_name)
//             self.app_uninstall(pkg_name)
//         return pkgs
// }
//     func unlock(self){
//         // """
//          // unlock screen """

//         self.open_identify()
//         self._default_session.press("home")
// }
//     func open_identify(self, theme="black"){
//         // """

//         Args:}
//             theme (str){ black or red
//         // """

//         self.shell([
//             "am", "start", "-W", "-n",
//             "com.github.uiautomator/.IdentifyActivity", "-e", "theme", theme
//         ])
// }
//     func _pidof_app(self, pkg_name){
//         // """

//         Return pid of package name
//         // """

//         text = self._reqse}ss.get(self.path2url("/pidof/" + pkg_name)).text
//         if text.isdigit(){
//             return int(text)
// }
//     func push_url(self, url, dst, mode=0o644){
//         // """

//         Args:}
//             url (str){} http url address
//             dst (str){} destination
//             mode (str){ file mode

//         Raises:
//             FileNotFoundError(py3) OSError(py2)
//         // """

//         modestr = oct(mode).replace("o", ")
//         r = self._reqsess.post(
//             self.path2url("/download"),
//             data={
//                 "url": url,
//                 "filepath": dst,
//                 "mode": modestr
//             })
//         if r.status_code != 200:
//             raise IOError("push-url", "%s -> %s" % (url, dst), r.text)
//         key = r.text.strip()
//         while 1:
//             r = self._reqsess.get(self.path2url("/download/" + key))
//             jdata = r.json()
//             message = jdata.get("message")
//             if message == "downloaded":
//                 log_print("downloaded")
//                 break
//             elif message == "downloading":
//                 progress = jdata.get("progress")
//                 if progress:
//                     copied_size = progress.get("copiedSize")
//                     total_size = progress.get("totalSize")
//                     log_print("{ { / {".format(
//                         message, humanize.naturalsize(copied_size),
//                         humanize.naturalsize(total_size)))
//                 else:
//                     log_print("downloading")
//             else:
//                 log_print("unknown json:" + str(jdata))
//                 raise IOError(message)
//             time.sleep(1)
// }
//     func push(self, src, dst, mode=0o644){
//         Args:}
//             src (path }or fileobj){ source file
//             dst (str){ destination can be folder or file path

//         Returns:
//             dict object, for example:

//                 {"mode": "0660", "size": 63, "target": "/sdcard/ABOUT.rst"}

//             Since chmod may fail in android, the result "mode" may not same with input args(mode)

//         Raises:
//             IOError(if push got something wrong)
//         // """

//         modestr = oct(mode).replace("o", ")
//         pathname = self.path2url("/upload/" +} dst.lstrip("/"))
//         if isinstance(src, six.string_types){
//             src = open(src, "rb")
//         r = self._reqsess.post(
//             pathname, data={"mode": modestr}, files={"file": src})
//         if r.status_code == 200:
//             return r.json()
//         raise IOError("push", "%s -> %s" % (src, dst), r.text)
// }
//     func pull(self, src, dst){
//         // """

//         Pull file from device to local

//         Raises:
//             FileNotFoundError(py3) OSError(py2)

//         Require atx-agent >= 0.0.9
//         // """

//         pathname = self.path2url("/raw/" + src.lstrip("/"))
//         r = self._reqsess.get(pathname, stream=True)
//         if r.status_code != 200:
//             raise FileNotFoundError("pull", src, r.text)
//         with open(dst, "wb") as f:
//             shutil.copyfileobj(r.raw, f)

//     @property}
//     func screenshot_uri(self){
//         return "http://%s:%d/screenshot/0" % (self._host, self._port)

//     @property}
//     func device_info(self){
//         if self.__devinfo:
//             return self.__devinfo
//         self.__devinfo = self._reqsess.get(self.path2url("/info")).json()
//         return self.__devinfo
// }
// func app_info(self, pkg_name){
//     // Get app info

//     // Args:}
//     //     pkg_name (str){ package name

//     // Return example:
//     //     {
//     //         "mainActivity": "com.github.uiautomator.MainActivity",
//     //         "label": "ATX",
//     //         "versionName": "1.1.7",
//     //         "versionCode": 1001007,
//     //         "size":1760809
//     //     }

//     // Raises:
//     //     UiaError
  
//     url = self.path2url("/packages/{0}/info".format(pkg_name))
//     resp = self._reqsess.get(url)
//     resp.raise_for_status()
//     resp = resp.json()}
//     if not resp.get("success"){
//         raise UiaError(resp.get("description", "unknown"))
//     return resp.get("data")
// }
// func app_icon(pkg_name){
//     // """

//     Returns:
//         PIL.Image

//     Raises:
//         UiaError
//     // """

//     from PIL import Image
//     url = self.path2url("/packages/{0}/icon".format(pkg_name))
//     resp = self._reqsess.get(url)
//     resp.raise_for_status()
//     return Image.open(io.BytesIO(resp.content))
// }

// func wlan_ip(self){
//     return self._reqsess.get(self.path2url("/wlan/ip")).text.strip()
// }


// func disable_popups(self, enable=True){
//     // """

//     Automatic click all popups
//     TODO: need fix
//     // """

//     raise NotImplementedError()
//     # self.watcher

//     if enable:
//         self.jsonrpc.setAccessibilityPatterns({
//             "com.android.packageinstaller":
//             [u"确定", u"安装", u"下一步", u"好", u"允许", u"我知道"],
//             "com.miui.securitycenter": [u"继续安装"],  # xiaomi
//             "com.lbe.security.miui": [u"允许"],  # xiaomi
//             "android": [u"好", u"安装"],  # vivo
//             "com.huawei.systemmanager": [u"立即删除"],  # huawei
//             "com.android.systemui": [u"同意"],  # 锤子
//         })
//     else:
//         self.jsonrpc.setAccessibilityPatterns({)
// }

// func session( pkg_name string, attach bool, launch_timeout int){
//         // """

//         Create a new session

//         Args:}
//             pkg_name (str)}{ android package name
//             attach (bool){ attac}h to already running app
//             launch_timeout (int){ launch timeout

//         Raises:
//             requests.HTTPError, SessionBrokenError
//         // """

//         if pkg_name is None:
//             return self._default_session

//         if not attach:
//             request_data = {"flags": "-W -S"}
//             if launch_timeout:
//                 request_data["timeout"] = str(launch_timeout)
//             resp = self._reqsess.post(
//                 self.path2url("/session/" + pkg_name), data=request_data)
//             if resp.status_code == 410:  # Gone
//                 raise SessionBrokenError(pkg_name, resp.text)
//             resp.raise_for_status()
//             jsondata = resp.json()
//             if not jsondata["success"]:
//                 raise SessionBrokenError("app launch failed",
//                                          jsondata["error"], jsondata["output"])

//             time.sleep(2.5)  # wait launch finished, maybe no need
//         pid = self._pidof_app(pkg_name)
//         if not pid:
//             raise SessionBrokenError(pkg_name)
//         return Session(self, pkg_name, pid)
// }
//     