scode =open("payload_char.c")
buf2=scode.readline().replace("unsigned char ","",-1).replace("[]","",-1).replace(" ","",-1).replace("buf=","buf=b",-1)
ZsPy = open("a.py","w")
ZsPy.writelines(buf2)