import uiautomator2 as u2
import time
import hashlib

d = u2.connect_wifi("192.168.248.54")

print(d.info)
d.app_start("com.nearme.themespace")
print(d(text = "精选字体").info)